# Weather Oracle

## Install


1. Install Go 1.16 for [OS](https://golang.org)


2. Install the `task` task runner. Read the [installation instructions](https://github.com/go-task/task/blob/master/docs/installation.md) or [go to the website](https://taskfile.dev/)


## Run

`API_ID=<weather-api-id> task run`

## Run tests

`task test`

## Use

`curl http://localhost:8081/v1/weather/get`

## Regenerate the code for the API spec

In case you want to update the API spec in `.proto` files 

1. Install [protoc](https://grpc.io/docs/protoc-installation/)

2. `task gen`