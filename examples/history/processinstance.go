package main

import (
	"flag"
	"fmt"
	camundaclientgo "github.com/citilinkru/camunda-client-go/v2"
	"os"
	"time"
)

var historyId string

func init() {
	flag.StringVar(&historyId, "id", "", "The id of the historic process instance to be retrieved")
	flag.Parse()
}

func main() {
	if historyId == "" {
		fmt.Println("Please use flag `--id` (The `id` of the historic process instance to be retrieved)")
		os.Exit(1)
	}

	client := camundaclientgo.NewClient(camundaclientgo.ClientOptions{
		EndpointUrl: "http://localhost:8080/engine-rest",
		ApiUser:     "demo",
		ApiPassword: "demo",
		Timeout:     time.Second * 10,
	})

	query := camundaclientgo.QueryHistoryProcessInstanceBy{
		Id: historyId,
	}

	processInstance, err := client.History.GetProcessInstance(&query)
	if err != nil {
		fmt.Printf("Can't get history Process Instance: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response Process Instance from history:\n %#+v\n", processInstance)
}
