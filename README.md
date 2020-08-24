# Camunda REST API client for golang
[![Build Status](https://travis-ci.org/citilinkru/camunda-client-go.svg?branch=master)](https://travis-ci.org/citilinkru/camunda-client-go)&nbsp;[![GoDoc](https://godoc.org/github.com/citilinkru/camunda-client-go?status.svg)](https://godoc.org/github.com/citilinkru/camunda-client-go)

Installation
-----------
latest stable version:

	go get gopkg.in/citilinkru/camunda-client-go.v1

latest unstable version (master branch):

	go get github.com/citilinkru/camunda-client-go
	
Usage
-----------

Create client:
```go
client := camunda_client_go.NewClient(camunda_client_go.ClientOptions{
	EndpointUrl: "http://localhost:8080/engine-rest",
    ApiUser: "demo",
    ApiPassword: "demo",
    Timeout: time.Second * 10,
})
```

Create deployment:
```go
file, err := os.Open("demo.bpmn")
if err != nil {
    fmt.Printf("Error read file: %s\n", err)
    return
}
result, err := client.Deployment.Create(camunda_client_go.ReqDeploymentCreate{
    DeploymentName: "DemoProcess",
    Resources: map[string]interface{}{
        "demo.bpmn": file,
    },
})
if err != nil {
    fmt.Printf("Error deploy process: %s\n", err)
    return
}

fmt.Printf("Result: %#+v\n", result)
```

Start instance:
```go
processKey := "demo-process"
result, err := client.ProcessDefinition.StartInstance(
	camunda_client_go.QueryProcessDefinitionBy{Key: &processKey},
	camunda_client_go.ReqStartInstance{},
)
if err != nil {
    fmt.Printf("Error start process: %s\n", err)
    return
}

fmt.Printf("Result: %#+v\n", result)
```

More examples
-----------
[Examples documentation](examples/README.md)

Usage for External task
-----------

Create external task processor:
```go
logger := func(err error) {
	fmt.Println(err.Error())
}
asyncResponseTimeout := 5000
proc := processor.NewProcessor(client, &processor.ProcessorOptions{
    WorkerId: "demo-worker",
    LockDuration: time.Second * 5,
    MaxTasks: 10,
    MaxParallelTaskPerHandler: 100,
    AsyncResponseTimeout: &asyncResponseTimeout,
}, logger)
```

Add and subscribe external task handler: 
```go
proc.AddHandler(
    &[]camunda_client_go.QueryFetchAndLockTopic{
        {TopicName: "HelloWorldSetter"},
    },
    func(ctx *processor.Context) error {
        fmt.Printf("Running task %s. WorkerId: %s. TopicName: %s\n", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

        err := ctx.Complete(processor.QueryComplete{
            Variables: &map[string]camunda_client_go.Variable {
                "result": {Value: "Hello world!", Type: "string"},
            },
        })
        if err != nil {
            fmt.Printf("Error set complete task %s: %s\n", ctx.Task.Id, err)
            
            return ctx.HandleFailure(processor.QueryHandleFailure{
                ErrorMessage: &errTxt,
                Retries: &retries,
                RetryTimeout: &retryTimeout,
            })
        }
        
        fmt.Printf("Task %s completed\n", ctx.Task.Id)
        return nil
    },
)
```

Features
-----------

* Support api version `7.11`
* Full support API `External Task`
* Full support API `Process Definition`
* Full support API `Deployment`
* Without external dependencies

Road map
-----------

* Full coverage by tests
* Full support references api

Testing
-----------
Unit-tests:
```bash
go test -v -race ./...
```

Run linter:
```bash
golangci-lint run
```

LICENSE
-----------
MIT

AUTHOR
-----------
Konstantin Osipov <k.osipov.msk@gmail.com>
