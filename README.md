# wiregarden

A broker daemon to provision the [WireGuard](https://www.wireguard.com/) peers over [gRPC](https://grpc.io/).

THIS PROJECT IS CURRENTLY UNDER DEVELOPMENT, JUST A PoC PHASE.
THERE IS THE POSSIBILITY TO BE CHANGED EVERYTHING WITHOUT NOTICES.

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
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:54321", grpc.WithInsecure(), grpc.WithBlock())
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

## Author

moznion (<moznion@mail.moznion.net>)

