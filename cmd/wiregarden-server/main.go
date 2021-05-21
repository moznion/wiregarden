package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/moznion/wiregarden/grpc"
	"github.com/moznion/wiregarden/routes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var revision string

func main() {
	defer leaveDyingMessageOnPanic()

	defaultPort := uint(0)
	defaultIPRoutingPolicyUsage := ""
	versionUsage := "show the version information"
	portUsage := "the port number to listen gRPC over TCP"
	ipRouteUsage := fmt.Sprintf(
		`the IP routing policy name (supported policies: "%s"). if this parameter is specified, this server manages ip route table automatically`,
		routes.IPRoutingPolicyIpcmd,
	)
	var shouldShowVersionInfo bool
	var port uint
	var prometheusExporterPort uint
	var ipRoutingPolicyName string
	flag.BoolVar(&shouldShowVersionInfo, "version", false, versionUsage)
	flag.BoolVar(&shouldShowVersionInfo, "v", false, versionUsage+" (shorthand)")
	flag.UintVar(&port, "port", defaultPort, portUsage)
	flag.UintVar(&port, "p", defaultPort, portUsage+" (shorthand)")
	flag.StringVar(&ipRoutingPolicyName, "ip-route", defaultIPRoutingPolicyUsage, ipRouteUsage)
	flag.UintVar(&prometheusExporterPort, "prom-port", defaultPort, "the port number to export the prometheus metrics")
	flag.Parse()

	if shouldShowVersionInfo {
		v, _ := json.Marshal(map[string]string{
			"revision":  revision,
			"goVersion": runtime.Version(),
		})
		fmt.Printf("%s\n", v)
		os.Exit(0)
	}

	if prometheusExporterPort > 0 {
		go func() {
			http.Handle("/metrics", promhttp.HandlerFor(
				prometheus.DefaultGatherer,
				promhttp.HandlerOpts{
					EnableOpenMetrics: true,
				},
			))
			log.Info().Uint("port", prometheusExporterPort).Msg("start the HTTP prometheus metrics exporter; you can retrieve metrics by 'GET /metrics'")
			log.Err(http.ListenAndServe(fmt.Sprintf(":%d", prometheusExporterPort), nil)).Send()
		}()
	}

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
