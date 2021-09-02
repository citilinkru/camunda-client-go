package main

import (
	"fmt"
	"time"

	camundaclientgo "github.com/citilinkru/camunda-client-go/v2"
)

func main() {
	client := camundaclientgo.NewClient(camundaclientgo.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	result, err := client.ProcessInstance.GetList(nil)
	if err != nil {
		fmt.Printf("Error deploy process: %s\n", err)
		return
	}
	fmt.Printf("Result: %#+v\n", result)
}
