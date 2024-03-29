# wiregarden [![Go Reference](https://pkg.go.dev/badge/github.com/moznion/wiregarden.svg)](https://pkg.go.dev/github.com/moznion/wiregarden) [![Check](https://github.com/moznion/wiregarden/actions/workflows/check.yml/badge.svg)](https://github.com/moznion/wiregarden/actions/workflows/check.yml)

A broker daemon to provision the [WireGuard](https://www.wireguard.com/) peers over [gRPC](https://grpc.io/).

## Usage

### gRPC Server

```
$ wiregarden-server --port $PORT
```

Note: if you faced like `operation not permitted` error, please run the server by the legit user.

### Client

See the example: [examples/wiregarden-client](https://github.com/moznion/wiregarden/tree/main/examples/wiregarden-client)

The following code is a simple example to retrieve peers of `wg0` device.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/moznion/wiregarden/grpc/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:54321", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	peersClient := messages.NewPeersClient(conn)

	resp, err := peersClient.GetPeers(context.Background(), &messages.GetPeersRequest{
		DeviceName: "wg0",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.Peers)
}
```

## Features

Currently, it supports the following features:

- GetDevices
- GetPeers
- RegisterPeers
- DeletePeers

## Hooks - **EXPERIMENTAL FEATURE**

It provides "hook" mechanism by the following interfaces:

- [handlers.PeersRegistrationHook](https://pkg.go.dev/github.com/moznion/wiregarden/grpc/handlers#PeersRegistrationHook) for `RegisterPeers`
- [handlers.PeersDeletionHook](https://pkg.go.dev/github.com/moznion/wiregarden/grpc/handlers#PeersDeletionHook) for `DeletePeers`

If you'd like to do the hook(s) on any operations, please pass the implementation(s) of the interface to [handlers.Peers](https://pkg.go.dev/github.com/moznion/wiregarden/grpc/handlers#Peers) struct.

Note: currently it doesn't provide a way to register the hooks by the default `wiregarden-server` command. If you'd like to run the server with the hooks, please make your own server launcher based on the [cmd/wiregarden-server/main.go](./cmd/wiregarden-server/main.go).

And, `RegisterPeersRequest.HooksPayload []byte` and `DeletePeersRequest.HooksPayload []byte` are the extension properties for each hook.

## Logging

Internally, this application / library uses [rs/zerolog](https://github.com/rs/zerolog) as a logger. You can configure the logger according to the manner of the zerolog. Please refer to the document of that.

## gRPC Library

It provides the wiregarden gRPC library for Java. Please refer to [this page](./ext/lib/java).

## Development Guide

### Pre-requirements

- Docker

### How to build a server binary

```
$ make build GOOS=linux GOARCH=amd64
```

Please change the `$GOOS` and `GOARCH` to your desired ones.

### How to generate protobuf files

#### Preparation

```sh
$ make container4protogen
```

#### Generate code

```sh
$ make proto
```

### E2E Testing

If you'd like to run the E2E tests, please set the environment value `E2E_TEST` with the non-empty value.

#### Build a container image for E2E testing

```
$ make e2e-docker-container
```

#### Push a container image to GitHub Docker Registry

```
$ make e2e-docker-push DOCKER_USER=${GITHUB_USERNAME} DOCKER_PSWD_FILE=/path/to/your/github/token/file
```

## Author

moznion (<moznion@mail.moznion.net>)

