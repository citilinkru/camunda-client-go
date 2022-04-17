package main

import (
	"fmt"
	camundaclientgo "github.com/citilinkru/camunda-client-go/v3"
	"os"
	"time"
)

func main() {
	client := camundaclientgo.NewClient(camundaclientgo.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	file, err := os.Open("HelloWorld.bpmn")
	if err != nil {
		fmt.Printf("Error read file: %s\n", err)
		return
	}
	result, err := client.Deployment.Create(camundaclientgo.ReqDeploymentCreate{
		DeploymentName: "HelloWorldProcessDemo",
		Resources: map[string]interface{}{
			"HelloWorld.bpmn": file,
		},
	})
	if err != nil {
		fmt.Printf("Error deploy process: %s\n", err)
		return
	}

	fmt.Printf("Result: %#+v\n", result)
}
