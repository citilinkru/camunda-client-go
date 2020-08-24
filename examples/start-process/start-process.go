package main

import (
	"fmt"
	camundaClient "github.com/citilinkru/camunda-client-go"
	"time"
)

func main() {
	client := camundaClient.NewClient(camundaClient.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	// start 1000 process
	for n := 0; n < 1000; n++ {
		processKey := "hello-world-process"

		isWorld := "false"
		if n%2 == 0 {
			isWorld = "true"
		}
		variables := map[string]camundaClient.Variable{
			"isWorld": {Value: isWorld, Type: "boolean"},
		}
		result, err := client.ProcessDefinition.StartInstance(
			camundaClient.QueryProcessDefinitionBy{Key: &processKey},
			camundaClient.ReqStartInstance{Variables: &variables},
		)
		if err != nil {
			fmt.Printf("Error start process: %s\n", err)
			return
		}

		fmt.Printf("Result: %#+v\n", result)
	}
}
