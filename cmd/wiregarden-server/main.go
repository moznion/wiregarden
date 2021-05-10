package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/moznion/wiregarden/grpc"
	"github.com/moznion/wiregarden/routes"
	"github.com/rs/zerolog/log"
)

func main() {
	defer leaveDyingMessageOnPanic()

	defaultPort := uint(0)
	defaultIPRoutingPolicyUsage := ""
	portUsage := "the port number to listen gRPC over TCP"
	ipRouteUsage := fmt.Sprintf(
		`the IP routing policy name (supported policies: "%s"). if this parameter is specified, this server manages ip route table automatically`,
		routes.IPRoutingPolicyIpcmd,
	)
	var port uint
	var ipRoutingPolicyName string
	flag.UintVar(&port, "port", defaultPort, portUsage)
	flag.UintVar(&port, "p", defaultPort, portUsage+" (shorthand)")
	flag.StringVar(&ipRoutingPolicyName, "ip-route", defaultIPRoutingPolicyUsage, ipRouteUsage)
	flag.Parse()

	s := grpc.Server{
		Port: uint16(port),
		IPRouter: func() routes.IPRouter {
			r := routes.IPRouterFrom(ipRoutingPolicyName)
			if r == nil {
				log.Info().Msg("ip routing policy is not specified")
			} else {
				log.Info().Str("policy", ipRoutingPolicyName).Msg("ip routing policy is specified; it starts auto ip route table management")
			}
			return r
		}(),
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
