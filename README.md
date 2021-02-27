# wiregarden

A broker daemon to provision the WireGuard peers over [gRPC](https://grpc.io/).

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

## Features

Currently, it supports the following features:

- GetDevices
- GetPeers
- RegisterPeers
- DeletePeers

## Development Guide

### Pre-requirements

- Docker

### How to generate a server binary

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

