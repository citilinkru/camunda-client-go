package processor

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"

	camundaClient "github.com/fundingasiagroup/camunda-client-go"
)

// Processor external task processor
type Processor struct {
	client  *camundaClient.Client
	options *Options
	logger  func(err error)
}

// Options options for Processor
type Options struct {
	// workerId for all request (default: `worker-{random_int}`)
	WorkerID string
	// lock duration for all external task
	LockDuration time.Duration
	// maximum tasks to receive for 1 request to camunda
	MaxTasks int
	// maximum running parallel task per handler
	MaxParallelTaskPerHandler int
	// use priority
	UsePriority *bool
	// long polling timeout
	//
	// Deprecated: Use LongPollingTimeout instead
	AsyncResponseTimeout *int
	// long polling timeout
	LongPollingTimeout time.Duration
}

// NewProcessor a create new instance Processor
func NewProcessor(client *camundaClient.Client, options *Options, logger func(err error)) *Processor {
	if options.WorkerID == "" {
		options.WorkerID = fmt.Sprintf("worker-%d", rand.Int())
	}

	return &Processor{
		client:  client,
		options: options,
		logger:  logger,
	}
}

// Handler a handler for external task
type Handler func(ctx *Context) error

// Context external task context
type Context struct {
	Task   *camundaClient.ResLockedExternalTask
	client *camundaClient.Client
}

// Complete a mark external task is complete
func (c *Context) Complete(query QueryComplete) error {
	return c.client.ExternalTask.Complete(c.Task.Id, camundaClient.QueryComplete{
		WorkerId:       &c.Task.WorkerId,
		Variables:      query.Variables,
		LocalVariables: query.LocalVariables,
	})
}

// HandleBPMNError handle external task BPMN error
func (c *Context) HandleBPMNError(query QueryHandleBPMNError) error {
	return c.client.ExternalTask.HandleBPMNError(c.Task.Id, camundaClient.QueryHandleBPMNError{
		WorkerId:     &c.Task.WorkerId,
		ErrorCode:    query.ErrorCode,
		ErrorMessage: query.ErrorMessage,
		Variables:    query.Variables,
	})
}

// HandleFailure handle external task failure
func (c *Context) HandleFailure(query QueryHandleFailure) error {
	return c.client.ExternalTask.HandleFailure(c.Task.Id, camundaClient.QueryHandleFailure{
		WorkerId:     &c.Task.WorkerId,
		ErrorMessage: query.ErrorMessage,
		ErrorDetails: query.ErrorDetails,
		Retries:      query.Retries,
		RetryTimeout: query.RetryTimeout,
	})
}

// AddHandler a add handler for external task
func (p *Processor) AddHandler(topics *[]camundaClient.QueryFetchAndLockTopic, handler Handler) {
	if topics != nil && p.options.LockDuration != 0 {
		for i := range *topics {
			v := &(*topics)[i]

			if v.LockDuration <= 0 {
				v.LockDuration = int(p.options.LockDuration / time.Millisecond)
			}
		}
	}

	var asyncResponseTimeout *int
	if p.options.AsyncResponseTimeout != nil {
		asyncResponseTimeout = p.options.AsyncResponseTimeout
	} else if p.options.LongPollingTimeout.Nanoseconds() > 0 {
		msValue := int(p.options.LongPollingTimeout.Nanoseconds() / int64(time.Millisecond))
		asyncResponseTimeout = &msValue
	}

	go p.startPuller(camundaClient.QueryFetchAndLock{
		WorkerId:             p.options.WorkerID,
		MaxTasks:             p.options.MaxTasks,
		UsePriority:          p.options.UsePriority,
		AsyncResponseTimeout: asyncResponseTimeout,
		Topics:               topics,
	}, handler)
}

func (p *Processor) startPuller(query camundaClient.QueryFetchAndLock, handler Handler) {
	var tasksChan = make(chan *camundaClient.ResLockedExternalTask)

	maxParallelTaskPerHandler := p.options.MaxParallelTaskPerHandler
	if maxParallelTaskPerHandler < 1 {
		maxParallelTaskPerHandler = 1
	}

	// create worker pool
	for i := 0; i < maxParallelTaskPerHandler; i++ {
		go p.runWorker(handler, tasksChan)
	}

	retries := 0
	for {
		tasks, err := p.client.ExternalTask.FetchAndLock(query)
		if err != nil {
			if retries < 60 {
				retries++
			}
			p.logger(fmt.Errorf("failed pull: %s, sleeping: %d seconds", err, retries))
			time.Sleep(time.Duration(retries) * time.Second)
			continue
		}
		retries = 0

		for _, task := range tasks {
			tasksChan <- task
		}
	}
}

func (p *Processor) runWorker(handler Handler, tasksChan chan *camundaClient.ResLockedExternalTask) {
	for task := range tasksChan {
		p.handle(&Context{
			Task:   task,
			client: p.client,
		}, handler)
	}
}

func (p *Processor) handle(ctx *Context, handler Handler) {
	defer func() {
		if r := recover(); r != nil {
			errMessage := fmt.Sprintf("fatal error in task: %s", r)
			errDetails := fmt.Sprintf("fatal error in task: %s\nStack trace: %s", r, string(debug.Stack()))
			err := ctx.HandleFailure(QueryHandleFailure{
				ErrorMessage: &errMessage,
				ErrorDetails: &errDetails,
			})

			if err != nil {
				p.logger(fmt.Errorf("error send handle failure: %s", err))
			}

			p.logger(errors.New(errDetails))
		}
	}()

	err := handler(ctx)
	if err != nil {
		errMessage := fmt.Sprintf("task error: %s", err)
		err = ctx.HandleFailure(QueryHandleFailure{
			ErrorMessage: &errMessage,
		})

		if err != nil {
			p.logger(fmt.Errorf("error send handle failure: %s", err))
		}

		p.logger(errors.New(errMessage))
	}
}
