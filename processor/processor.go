package processor

import (
	"errors"
	"fmt"
	"github.com/citilinkru/camunda-client-go"
	"math/rand"
	"runtime/debug"
	"time"
)

// Processor external task processor
type Processor struct {
	client  *camunda_client_go.Client
	options *ProcessorOptions
	logger  func(err error)
}

// ProcessorOptions options for Processor
type ProcessorOptions struct {
	// workerId for all request (default: `worker-{random_int}`)
	WorkerId string
	// lock duration for all external task
	LockDuration time.Duration
	// maximum tasks to receive for 1 request to camunda
	MaxTasks int
	// maximum running parallel task per handler
	MaxParallelTaskPerHandler int
	// use priority
	UsePriority *bool
	// long polling timeout
	AsyncResponseTimeout *int
}

// NewProcessor a create new instance Processor
func NewProcessor(client *camunda_client_go.Client, options *ProcessorOptions, logger func(err error)) (*Processor, error) {
	if options.WorkerId == "" {
		options.WorkerId = fmt.Sprintf("worker-%d", rand.Int())
	}

	if options.MaxParallelTaskPerHandler < 1 {
		return nil, fmt.Errorf("MaxParallelTaskPerHandler must be greater than 0")
	}

	if options.LockDuration < 1 {
		return nil, fmt.Errorf("LockDuration must be greater than 0")
	}

	if options.AsyncResponseTimeout != nil && *options.AsyncResponseTimeout < 1 {
		return nil, fmt.Errorf("AsyncResponseTimeout must be greater than 0")
	}

	return &Processor{
		client:  client,
		options: options,
		logger:  logger,
	}, nil
}

// Handler a handler for external task
type Handler func(ctx *Context) error

// Context external task context
type Context struct {
	Task   *camunda_client_go.ResLockedExternalTask
	client *camunda_client_go.Client
}

// Complete a mark external task is complete
func (c *Context) Complete(query QueryComplete) error {
	return c.client.ExternalTask.Complete(c.Task.Id, camunda_client_go.QueryComplete{
		WorkerId:       &c.Task.WorkerId,
		Variables:      query.Variables,
		LocalVariables: query.LocalVariables,
	})
}

// HandleBPMNError handle external task BPMN error
func (c *Context) HandleBPMNError(query QueryHandleBPMNError) error {
	return c.client.ExternalTask.HandleBPMNError(c.Task.Id, camunda_client_go.QueryHandleBPMNError{
		WorkerId:     &c.Task.WorkerId,
		ErrorCode:    query.ErrorCode,
		ErrorMessage: query.ErrorMessage,
		Variables:    query.Variables,
	})
}

// HandleFailure handle external task failure
func (c *Context) HandleFailure(query QueryHandleFailure) error {
	return c.client.ExternalTask.HandleFailure(c.Task.Id, camunda_client_go.QueryHandleFailure{
		WorkerId:     &c.Task.WorkerId,
		ErrorMessage: query.ErrorMessage,
		ErrorDetails: query.ErrorDetails,
		Retries:      query.Retries,
		RetryTimeout: query.RetryTimeout,
	})
}

// AddHandler a add handler for external task
func (p *Processor) AddHandler(topics *[]camunda_client_go.QueryFetchAndLockTopic, handler Handler) {
	if topics != nil && p.options.LockDuration != 0 {
		for i := range *topics {
			v := &(*topics)[i]

			if v.LockDuration <= 0 {
				v.LockDuration = int(p.options.LockDuration / time.Millisecond)
			}
		}
	}
	go p.startPuller(camunda_client_go.QueryFetchAndLock{
		WorkerId:             p.options.WorkerId,
		MaxTasks:             p.options.MaxTasks,
		UsePriority:          p.options.UsePriority,
		AsyncResponseTimeout: p.options.AsyncResponseTimeout,
		Topics:               topics,
	}, handler)
}

func (p *Processor) startPuller(query camunda_client_go.QueryFetchAndLock, handler Handler) {
	var tasksChan = make(chan *camunda_client_go.ResLockedExternalTask)

	maxParallelTaskPerHandler := p.options.MaxParallelTaskPerHandler
	if maxParallelTaskPerHandler < 1 {
		maxParallelTaskPerHandler = 1
	}

	// create worker pool
	for i := 0; i < maxParallelTaskPerHandler; i++ {
		go p.runWorker(handler, tasksChan)
	}

	for {
		tasks, err := p.client.ExternalTask.FetchAndLock(query)
		if err != nil {
			p.logger(fmt.Errorf("failed pull: %s", err))
			continue
		}

		for _, task := range tasks {
			tasksChan <- task
		}
	}
}

func (p *Processor) runWorker(handler Handler, tasksChan chan *camunda_client_go.ResLockedExternalTask) {
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
