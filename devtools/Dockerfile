FROM golang:1.16.2-buster

RUN apt -y update && apt -y upgrade && apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

