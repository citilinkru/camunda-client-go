package main

import (
	"fmt"
	camundaclientgo "github.com/citilinkru/camunda-client-go/v2"
	"github.com/citilinkru/camunda-client-go/v2/processor"
	"time"
)

func main() {
	client := camundaclientgo.NewClient(camundaclientgo.ClientOptions{
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
		&[]camundaclientgo.QueryFetchAndLockTopic{
			{TopicName: "PrintHello"},
		},
		func(ctx *processor.Context) error {
			fmt.Printf("Running task %s. WorkerId: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

			time.Sleep(time.Second * 1)
			fmt.Println("Hello")

			err := ctx.Complete(processor.QueryComplete{
				Variables: &map[string]camundaclientgo.Variable{
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
		&[]camundaclientgo.QueryFetchAndLockTopic{
			{TopicName: "PrintWorld"},
		},
		func(ctx *processor.Context) error {
			fmt.Printf("Running task %s. WorkerId: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

			time.Sleep(time.Second * 1)
			fmt.Println("World")

			err := ctx.Complete(processor.QueryComplete{
				Variables: &map[string]camundaclientgo.Variable{
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
