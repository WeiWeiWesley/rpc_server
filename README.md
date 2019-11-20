RPC Server
===

## About

This package provide a customize json-rpc mico-service framework

## Install

```console
go get -u github.com/WeiWeiWesley/rpc_server
```

## Usage

### Run

* ENV
* SERVICE (Service name)

For start service requires ENV & SERVICE variables

```console
SERVICE=Echo ENV=local go run main.go
```

### Add Service

You can add your service in service floder

```
rpc_server
    |-config
        |-env
            |-local
                |-default.toml
                |-echo.toml
    |-core
    |-scheme
    |-service
        |-echo.go
    |-main.go
    |-README.md
```

Here is a example that you have to do.

1. Import "rpc_server/core/rpc"
2. Use rpc.Register() when init()
3. Implement your service

```go
package service

// 1. Import "rpc_server/core/rpc"
import (
	"rpc_server/core/rpc"
)

// 2. Use rpc.Register() when init()
func init() {
	rpc.Register("Echo", "A Echo Service", &Echo{})
}

// 3. Implement your service
// Echo echo service
type Echo struct{}

//Ping Ping
func (d *Echo) Ping(args *string, result *string) (rpcErr error) {
	*result = "Pong"
	return
}

```

### Run All Service At Once

We used to run one service at one pod or one mechine in production envrinment. But if you want to start all at once, here is it.

```console
SERVICE=All ENV=local go run main.go
```