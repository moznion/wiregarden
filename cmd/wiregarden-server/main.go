package main

import (
	"context"
	"flag"
	"os"
	"runtime"

	"github.com/moznion/wiregarden/grpc"
	"github.com/rs/zerolog/log"
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
	err := s.Run(ctx)
	log.Fatal().Err(err).Msg("")
}

func leaveDyingMessageOnPanic() {
	if err := recover(); err != nil {
		log.Error().Interface("recoveredErr", err).Msg("panic occurred; show stacktrace")
		for depth := 0; ; depth++ {
			_, filename, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			log.Error().Str("filename", filename).Int("lineNumber", line).Msg("")
		}
		os.Exit(1)
	}
}
