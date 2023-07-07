# Camunda REST API client for golang
[![Go Report Card](https://goreportcard.com/badge/github.com/citilinkru/camunda-client-go)](https://goreportcard.com/report/github.com/citilinkru/camunda-client-go)
[![codecov](https://codecov.io/gh/citilinkru/camunda-client-go/branch/master/graph/badge.svg?token=53NH949TQY)](https://codecov.io/gh/citilinkru/camunda-client-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/citilinkru/camunda-client-go.svg)](https://pkg.go.dev/github.com/citilinkru/camunda-client-go)
[![Release](https://img.shields.io/github/release/citilinkru/camunda-client-go.svg?style=flat-square)](https://github.com/citilinkru/camunda-client-go/releases/latest)

Installation
-----------
	go get github.com/citilinkru/camunda-client-go/v3
	
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
proc := processor.NewProcessor(client, &processor.Options{
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
    []*camunda_client_go.QueryFetchAndLockTopic{
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
* Full support API `Process Instance`
* Full support API `Deployment`
* Partial support API `History`
* Partial support API `Tenant`
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
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.45.2 golangci-lint run -v
```

Integration tests:
```bash
docker run --rm --name camunda -p 8080:8080 camunda/camunda-bpm-platform
```

```bash
go test -tags=integration -failfast ./...
```

Examples:
---------
Go to [examples directory](examples/README.md) and follow the instructions to run the examples.

CONTRIBUTE
-----------
 * write code
 * run `go fmt ./...`
 * run all linters and tests (see above)
 * run all examples (see above)
 * create a PR describing the changes

LICENSE
-----------
MIT

AUTHOR
-----------
Konstantin Osipov <k.osipov.msk@gmail.com>