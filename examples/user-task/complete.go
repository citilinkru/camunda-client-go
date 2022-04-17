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

	query := camundaclientgo.UserTaskGetListQuery{
		CreatedAfter: time.Now().Add(-500 * time.Hour),
	}

	cnt, err := client.UserTask.GetListCount(&query)
	if err != nil {
		fmt.Printf("Can't get task list count: %s", err)
		os.Exit(1)
	}

	fmt.Printf("total task count: %d\n", cnt)

	tasks, err := client.UserTask.GetList(&query)
	if err != nil {
		fmt.Printf("Can't get task list: %s", err)
		os.Exit(1)
	}

	for i, taks := range tasks {
		fmt.Printf("%02d. UserTask: %s - %s - ", i+1, taks.Id, taks.Name)

		err = taks.Complete(camundaclientgo.QueryUserTaskComplete{})
		if err != nil {
			fmt.Printf("ERROR: %s", err)
		} else {
			fmt.Printf("Complete")
		}

		fmt.Printf("\n")
	}
}
