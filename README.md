# Weather Oracle

The intention of this application is educational.

The code will demonstrate how gRPC can be used in Go.

## Install


1. Install Go 1.16 for [OS](https://golang.org)


2. Install the `task` task runner. Read the [installation instructions](https://github.com/go-task/task/blob/master/docs/installation.md) or [go to the website](https://taskfile.dev/)


## Run

`API_ID=<weather-api-id> task run`

## Run tests

`task test`

## Use

`curl http://localhost:8081/v1/weather/get?city=oslo`

## Regenerate the code from the API spec

In case you want to update the API spec in `.proto` files 

1. Install [protoc](https://grpc.io/docs/protoc-installation/)

2. `task gen`
