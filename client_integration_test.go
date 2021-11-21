// +build integration

package camunda_client_go

import (
	"fmt"
	"os"
	"time"
)

var client *Client
var logger func(err error)

func init() {
	client = NewClient(ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	logger = func(err error) {
		fmt.Println(err.Error())
	}

	file, err := os.Open("examples/deployment/HelloWorld.bpmn")
	if err != nil {
		fmt.Printf("Error read file: %s\n", err)
		os.Exit(1)
	}
	_, err = client.Deployment.Create(ReqDeploymentCreate{
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
