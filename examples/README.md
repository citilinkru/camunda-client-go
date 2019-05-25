# Examples
These examples will help you learn basic Camunda use cases.

## Before use:
Run Camunda:
```bash
docker run -d --name camunda -p 8080:8080 camunda/camunda-bpm-platform
```

## Examples use scenario

### Deploy helloWorld.pbmn
```bash
cd deployment
go build
./deployment
```

### Run external task processor
```bash
cd processor
go build
./processor
```

## Start 1000 process
```bash
cd start-process
go build
./start-process
 ```