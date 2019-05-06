package processor

import (
	"fmt"
	"github.com/citilinkru/camunda-client-go"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"runtime/debug"
	"sync"
	"time"
)

// Processor external task processor
type Processor struct {
	client  *camunda_client_go.Client
	options *ProcessorOptions
	logger  *log.Logger
}

// ProcessorOptions options for Processor
type ProcessorOptions struct {
	// workerId for all request (default: `worker-{random_int}`)
	WorkerId string
	// lock duration for all external task
	LockDuration time.Duration
	// maximum tasks to receive for 1 request
	MaxTasks int
	// use priority
	UsePriority *bool
	// long polling timeout
	AsyncResponseTimeout *int
}

// NewProcessor a create new instance Processor
func NewProcessor(client *camunda_client_go.Client, options *ProcessorOptions, logger *log.Logger) *Processor {
	if options.WorkerId == "" {
		options.WorkerId = fmt.Sprintf("worker-%d", rand.Int())
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
	for {
		var wg sync.WaitGroup
		tasks, err := p.client.ExternalTask.FetchAndLock(query)
		if err != nil {
			p.logger.Errorf("failed pull: %s", err)
			continue
		}

		for _, task := range tasks {
			wg.Add(1)
			go p.handle(&Context{
				Task:   task,
				client: p.client,
			},
				handler,
				&wg,
			)
		}

		wg.Wait()
	}
}

func (p *Processor) handle(ctx *Context, handler Handler, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			errMessage := fmt.Sprintf("Fatal error in task: %s", r)
			errDetails := fmt.Sprintf("Fatal error in task: %s\nStack trace: %s", r, string(debug.Stack()))
			err := ctx.HandleFailure(QueryHandleFailure{
				ErrorMessage: &errMessage,
				ErrorDetails: &errDetails,
			})

			if err != nil {
				p.logger.Errorf("Error send handle failure: %s", err)
			}

			p.logger.Error(errDetails)
		}

		wg.Done()
	}()

	err := handler(ctx)
	if err != nil {
		errMessage := fmt.Sprintf("Task error: %s", err)
		err = ctx.HandleFailure(QueryHandleFailure{
			ErrorMessage: &errMessage,
		})

		if err != nil {
			p.logger.Errorf("Error send handle failure: %s", err)
		}

		p.logger.Warn(errMessage)
	}
}
