// +build integration

package processor

import (
	"fmt"
	camundaclientgo "github.com/citilinkru/camunda-client-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var client *camundaclientgo.Client
var logger func(err error)

func init() {
	client = camundaclientgo.NewClient(camundaclientgo.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	logger = func(err error) {
		fmt.Println(err.Error())
	}

	file, err := os.Open("../examples/deployment/HelloWorld.bpmn")
	if err != nil {
		fmt.Printf("Error read file: %s\n", err)
		os.Exit(1)
	}
	_, err = client.Deployment.Create(camundaclientgo.ReqDeploymentCreate{
		DeploymentName: "HelloWorldProcessDemo",
		Resources: map[string]interface{}{
			"HelloWorld.bpmn": file,
		},
	})
	if err != nil {
		fmt.Printf("Error deploy process: %s\n", err)
		os.Exit(1)
	}
}

func TestComplete(t *testing.T) {
	proc := NewProcessor(client, &Options{
		WorkerId:                  "hello-world-worker",
		LockDuration:              time.Second * 5,
		MaxTasks:                  10,
		MaxParallelTaskPerHandler: 100,
		LongPollingTimeout:        5 * time.Second,
	}, logger)

	processKey := "hello-world-process"

	variables := map[string]camundaclientgo.Variable{
		"isWorld": {Value: false, Type: "boolean"},
	}

	_, err := client.ProcessDefinition.StartInstance(
		camundaclientgo.QueryProcessDefinitionBy{Key: &processKey},
		camundaclientgo.ReqStartInstance{Variables: &variables},
	)
	assert.NoError(t, err)

	done := make(chan bool)
	proc.AddHandler(
		[]*camundaclientgo.QueryFetchAndLockTopic{
			{TopicName: "PrintHello"},
		},
		func(ctx *Context) error {
			err := ctx.Complete(QueryComplete{})
			assert.NoError(t, err)
			done <- true
			return err
		},
	)

	select {
	case <-done:
		return
	case <-time.After(time.Second * 10):
		t.Error("Handler timeout")
	}
}
