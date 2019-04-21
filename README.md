# Camunda REST API client for golang
[![Build Status](https://travis-ci.org/citilinkru/camunda-client-go.svg?branch=master)](https://travis-ci.org/citilinkru/camunda-client-go)&nbsp;[![GoDoc](https://godoc.org/github.com/citilinkru/camunda-client-go?status.svg)](https://godoc.org/github.com/citilinkru/camunda-client-go)

Installation
-----------
	go get github.com/citilinkru/camunda-client-go

Usage
-----------

Create client:
```go
timeout := time.Second * 10
client := camunda_client_go.NewClient(camunda_client_go.ClientOptions{
    ApiUser: "demo",
    ApiPassword: "demo",
    Timeout: &timeout,
})
```

Create deployment:
```go
file, err := os.Open("demo.bpmn")
if err != nil {
    logger.Errorf("Error read file: %s", err)
    return
}
result, err = client.Deployment.Create(camunda_client_go.ReqDeploymentCreate{
    DeploymentName: "DemoProcess",
    Resources: map[string]interface{}{
        "demo.bpmn": file,
    },
})
if err != nil {
    logger.Errorf("Error deploy process: %s", err)
    return
}

logger.infof("Result: %#+v", result)
```

Start instance:
```go
processKey := "DemoProcess"
result, err = client.ProcessDefinition.StartInstance(
	camunda_client_go.QueryProcessDefinitionBy{Key: &processKey},
	camunda_client_go.ReqStartInstance{},
)
if err != nil {
    logger.Errorf("Error start process: %s", err)
    return
}

logger.infof("Result: %#+v", result)
```


Usage for External task
-----------

Create external task processor:
```go
logger := logrus.New()
asyncResponseTimeout := 5000
proc := processor.NewProcessor(client, &processor.ProcessorOptions{
    WorkerId: "demo-worker",
    LockDuration: time.Second * 5,
    MaxTasks: 1,
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
        logger.Infof("Running task %s. WorkerId: %s. TopicName: %s", ctx.Task.Id, ctx.Task.WorkerId, ctx.Task.TopicName)

        err := ctx.Complete(processor.QueryComplete{
            Variables: &map[string]camunda_client_go.Variable {
                "result": {Value: "Hello world!", Type: "string"},
            },
        })
        if err != nil {
            logger.Errorf("Error set complete task %s: %s", ctx.Task.Id, err)
            
            return ctx.HandleFailure(processor.QueryHandleFailure{
                ErrorMessage: &errTxt,
                Retries: &retries,
                RetryTimeout: &retryTimeout,
            })
        }
        
        logger.Infof("Task %s completed", ctx.Task.Id)
        return nil
    },
)
```

WARNING
-----------
**This project is still under development. Use code with caution!**

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