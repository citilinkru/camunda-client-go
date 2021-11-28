package main

import (
	"fmt"
	camundaclientgo "github.com/citilinkru/camunda-client-go"
	"time"
)

func main() {
	client := camundaclientgo.NewClient(camundaclientgo.ClientOptions{
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
		variables := map[string]camundaclientgo.Variable{
			"isWorld": {Value: isWorld, Type: "boolean"},
		}
		result, err := client.ProcessDefinition.StartInstance(
			camundaclientgo.QueryProcessDefinitionBy{Key: &processKey},
			camundaclientgo.ReqStartInstance{Variables: &variables},
		)
		if err != nil {
			fmt.Printf("Error start process: %s\n", err)
			return
		}

		fmt.Printf("Result: %#+v\n", result)
	}
}
