package main

import (
	"fmt"
	camunda_client_go "github.com/citilinkru/camunda-client-go"
	"github.com/citilinkru/camunda-client-go/processor"
	"os"
	"time"
)

func main() {
	client := camunda_client_go.NewClient(camunda_client_go.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	logger := func(err error) {
		fmt.Println(err.Error())
	}
	asyncResponseTimeout := 5000
	proc, err := processor.NewProcessor(client, &processor.ProcessorOptions{
		WorkerId:                  "hello-world-worker",
		LockDuration:              time.Second * 5,
		MaxTasks:                  10,
		MaxParallelTaskPerHandler: 100,
		AsyncResponseTimeout:      &asyncResponseTimeout,
	}, logger)
	if err != nil {
		fmt.Printf("Can`t create processor: %s\n", err)
		os.Exit(1)
	}

	proc.AddHandler(
		&[]camunda_client_go.QueryFetchAndLockTopic{
			{TopicName: "PrintHello"},
		},
		func(ctx *processor.Context) error {
			fmt.Printf("Running task %s. WorkerId: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

			time.Sleep(time.Second * 1)
			fmt.Println("Hello")

			err := ctx.Complete(processor.QueryComplete{
				Variables: &map[string]camunda_client_go.Variable{
					"status": {Value: "true", Type: "boolean"},
				},
			})
			if err != nil {
				fmt.Printf("Error set complete task %s: %s\n", ctx.Task.Id, err)
			}

			fmt.Printf("Task %s completed\n", ctx.Task.Id)
			return nil
		},
	)

	proc.AddHandler(
		&[]camunda_client_go.QueryFetchAndLockTopic{
			{TopicName: "PrintWorld"},
		},
		func(ctx *processor.Context) error {
			fmt.Printf("Running task %s. WorkerId: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

			time.Sleep(time.Second * 1)
			fmt.Println("World")

			err := ctx.Complete(processor.QueryComplete{
				Variables: &map[string]camunda_client_go.Variable{
					"status": {Value: "true", Type: "boolean"},
				},
			})
			if err != nil {
				fmt.Printf("Error set complete task %s: %s\n", ctx.Task.Id, err)
			}

			fmt.Printf("Task %s completed\n", ctx.Task.Id)
			return nil
		},
	)

	fmt.Println("Processor is started")

	// wait...
	for {
		time.Sleep(time.Second * 180)
	}
}
