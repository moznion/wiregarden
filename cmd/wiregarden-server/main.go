package main

import (
	"context"
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/moznion/wiregarden/grpc"
)

func main() {
	defer leaveDyingMessageOnPanic()

	defaultPort := uint(0)
	portUsage := "the port number to listen gRPC over TCP"
	var port uint
	flag.UintVar(&port, "port", defaultPort, portUsage)
	flag.UintVar(&port, "p", defaultPort, portUsage+" (shorthand)")
	flag.Parse()

	s := grpc.Server{
		Port: uint16(port),
	}

	ctx := context.Background()
	log.Fatal(s.Run(ctx))
}

func leaveDyingMessageOnPanic() {
	if err := recover(); err != nil {
		log.Printf("[CRIT] panic: %+v\n", err)
		for depth := 0; ; depth++ {
			_, filename, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			log.Printf("[CRIT]   %v:%d", filename, line)
		}
		os.Exit(1)
	}
}
