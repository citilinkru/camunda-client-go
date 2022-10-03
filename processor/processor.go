package processor

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime/debug"
	"sync"
	"time"

	camundaclientgo "github.com/citilinkru/camunda-client-go/v3"
)

// Processor external task processor
type Processor struct {
	client  *camundaclientgo.Client
	options *Options
	logger  func(err error)

	// shutdown support
	workerGroup *sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
}

// Options options for Processor
type Options struct {
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
	//
	// Deprecated: Use LongPollingTimeout instead
	AsyncResponseTimeout *int
	// long polling timeout
	LongPollingTimeout time.Duration
}

// NewProcessor a create new instance Processor
func NewProcessor(client *camundaclientgo.Client, options *Options, logger func(err error)) *Processor {
	rand.Seed(time.Now().UnixNano())
	if options.WorkerId == "" {
		// #nosec G404 This is valid for worker selection
		options.WorkerId = fmt.Sprintf("worker-%d", rand.Int())
	}

	ctx, cancel := context.WithCancel(context.Background())
	workerGroup := new(sync.WaitGroup)

	return &Processor{
		client:      client,
		options:     options,
		logger:      logger,
		workerGroup: workerGroup,
		ctx:         ctx,
		cancel:      cancel}
}

// Handler a handler for external task
type Handler func(ctx *Context) error

// Context external task context
type Context struct {
	Task   *camundaclientgo.ResLockedExternalTask
	client *camundaclientgo.Client
}

// Complete a mark external task is complete
func (c *Context) Complete(query QueryComplete) error {
	return c.client.ExternalTask.Complete(c.Task.Id, camundaclientgo.QueryComplete{
		WorkerId:       &c.Task.WorkerId,
		Variables:      query.Variables,
		LocalVariables: query.LocalVariables,
	})
}

// Extend the lock for a new duration
func (c *Context) ExtendLock(newDurationMS int) error {
	return c.client.ExternalTask.ExtendLock(c.Task.Id, camundaclientgo.QueryExtendLock{
		NewDuration: &newDurationMS,
		WorkerId:    &c.Task.WorkerId,
	})
}

// HandleBPMNError handle external task BPMN error
func (c *Context) HandleBPMNError(query QueryHandleBPMNError) error {
	return c.client.ExternalTask.HandleBPMNError(c.Task.Id, camundaclientgo.QueryHandleBPMNError{
		WorkerId:     &c.Task.WorkerId,
		ErrorCode:    query.ErrorCode,
		ErrorMessage: query.ErrorMessage,
		Variables:    query.Variables,
	})
}

// HandleFailure handle external task failure
func (c *Context) HandleFailure(query QueryHandleFailure) error {
	return c.client.ExternalTask.HandleFailure(c.Task.Id, camundaclientgo.QueryHandleFailure{
		WorkerId:     &c.Task.WorkerId,
		ErrorMessage: query.ErrorMessage,
		ErrorDetails: query.ErrorDetails,
		Retries:      query.Retries,
		RetryTimeout: query.RetryTimeout,
	})
}

// Shutdown stop this processor and wait for running handlers to complete in-flight processing.
// The Processor cannot be reused after shutdown.
func (p *Processor) Shutdown() {
	p.cancel()
	p.workerGroup.Wait()
}

// AddHandler register an external task handler and start pulling for work. Calling this after a Shutdown has no effect.
func (p *Processor) AddHandler(topics []*camundaclientgo.QueryFetchAndLockTopic, handler Handler) {
	if topics != nil && p.options.LockDuration != 0 {
		for _, v := range topics {
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

	p.startPuller(camundaclientgo.QueryFetchAndLock{
		WorkerId:             p.options.WorkerId,
		MaxTasks:             p.options.MaxTasks,
		UsePriority:          p.options.UsePriority,
		AsyncResponseTimeout: asyncResponseTimeout,
		Topics:               topics,
	}, handler)
}

func (p *Processor) startPuller(query camundaclientgo.QueryFetchAndLock, handler Handler) {
	var tasksChan = make(chan *camundaclientgo.ResLockedExternalTask)

	maxParallelTaskPerHandler := p.options.MaxParallelTaskPerHandler
	if maxParallelTaskPerHandler < 1 {
		maxParallelTaskPerHandler = 1
	}

	// create worker pool
	for i := 0; i < maxParallelTaskPerHandler; i++ {
		p.workerGroup.Add(1)
		go p.runWorker(handler, tasksChan)
	}

	go func() {
		retries := 0
		for {
			select {
			case <-p.ctx.Done():
				close(tasksChan)
				return
			default:
				tasks, err := p.client.ExternalTask.FetchAndLock(query)
				if err != nil {
					if retries < 60 {
						retries += 1
					}
					p.logger(fmt.Errorf("failed pull: %w, sleeping: %d seconds", err, retries))
					time.Sleep(time.Duration(retries) * time.Second)
					continue
				}
				retries = 0

				for _, task := range tasks {
					tasksChan <- task
				}
			}
		}
	}()
}

func (p *Processor) runWorker(handler Handler, tasksChan chan *camundaclientgo.ResLockedExternalTask) {
	defer p.workerGroup.Done()
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
				p.logger(fmt.Errorf("error send handle failure: %w", err))
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
			p.logger(fmt.Errorf("error send handle failure: %w", err))
		}

		p.logger(errors.New(errMessage))
	}
}
