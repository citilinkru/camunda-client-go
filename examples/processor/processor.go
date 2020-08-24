package main

import (
	"fmt"
	camundaClient "github.com/fundingasiagroup/camunda-client-go"
	"github.com/fundingasiagroup/camunda-client-go/processor"
	"time"
)

func main() {
	client := camundaClient.NewClient(camundaClient.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	logger := func(err error) {
		fmt.Println(err.Error())
	}
	proc := processor.NewProcessor(client, &processor.ProcessorOptions{
		WorkerId:                  "hello-world-worker",
		LockDuration:              time.Second * 5,
		MaxTasks:                  10,
		MaxParallelTaskPerHandler: 100,
		LongPollingTimeout:        5 * time.Second,
	}, logger)

	proc.AddHandler(
		&[]camundaClient.QueryFetchAndLockTopic{
			{TopicName: "PrintHello"},
		},
		func(ctx *processor.Context) error {
			fmt.Printf("Running task %s. WorkerID: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

			time.Sleep(time.Second * 1)
			fmt.Println("Hello")

			err := ctx.Complete(processor.QueryComplete{
				Variables: &map[string]camundaClient.Variable{
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
		&[]camundaClient.QueryFetchAndLockTopic{
			{TopicName: "PrintWorld"},
		},
		func(ctx *processor.Context) error {
			fmt.Printf("Running task %s. WorkerID: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

			time.Sleep(time.Second * 1)
			fmt.Println("World")

			err := ctx.Complete(processor.QueryComplete{
				Variables: &map[string]camundaClient.Variable{
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
