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
	"github.com/moznion/wiregarden/grpc/metrics"
	"github.com/moznion/wiregarden/internal"
	"github.com/moznion/wiregarden/routes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	defer leaveDyingMessageOnPanic()

	defaultPort := uint(0)
	defaultIPRoutingPolicyUsage := ""
	defaultUnixSocketPath := ""
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
	var unixSocketPath string
	flag.BoolVar(&shouldShowVersionInfo, "version", false, versionUsage)
	flag.BoolVar(&shouldShowVersionInfo, "v", false, versionUsage+" (shorthand)")
	flag.UintVar(&port, "port", defaultPort, portUsage)
	flag.UintVar(&port, "p", defaultPort, portUsage+" (shorthand)")
	flag.StringVar(&ipRoutingPolicyName, "ip-route", defaultIPRoutingPolicyUsage, ipRouteUsage)
	flag.UintVar(&prometheusExporterPort, "prom-port", defaultPort, "the port number to export the prometheus metrics")
	flag.StringVar(&unixSocketPath, "unix-socket", defaultUnixSocketPath, "path to a UNIX domain socket to listen")
	flag.Parse()

	if shouldShowVersionInfo {
		v, _ := json.Marshal(map[string]string{
			"version":   internal.Version,
			"revision":  internal.Revision,
			"goVersion": runtime.Version(),
		})
		fmt.Printf("%s\n", v)
		os.Exit(0)
	}

	if port != defaultPort && unixSocketPath != defaultUnixSocketPath {
		internal.Logger.Fatal().Msg("port and unix-socket options are exclusive; it must be used either one")
		os.Exit(1)
	}

	var grpcPrometheusMetricsRegister metrics.PrometheusMetricsRegisterable = &metrics.NOPPrometheusMetricsRegister{}
	if prometheusExporterPort > 0 {
		go func() {
			http.Handle("/metrics", promhttp.HandlerFor(
				prometheus.DefaultGatherer,
				promhttp.HandlerOpts{
					EnableOpenMetrics: true,
				},
			))
			internal.Logger.Info().Uint("port", prometheusExporterPort).Msg("start the HTTP prometheus metrics exporter; you can retrieve metrics by 'GET /metrics'")
			internal.Logger.Err(http.ListenAndServe(fmt.Sprintf(":%d", prometheusExporterPort), nil)).Send()
		}()
		grpcPrometheusMetricsRegister = metrics.NewPrometheusMetricsRegister()
	}

	s := grpc.Server{
		Port: uint16(port),
		IPRouter: func() routes.IPRouter {
			r := routes.IPRouterFrom(ipRoutingPolicyName)
			if r == nil {
				internal.Logger.Info().Msg("ip routing policy is not specified")
			} else {
				internal.Logger.Info().Str("policy", ipRoutingPolicyName).Msg("ip routing policy is specified; it starts auto ip route table management")
			}
			return r
		}(),
		PrometheusMetricsRegister: grpcPrometheusMetricsRegister,
	}

	ctx := context.Background()
	err := s.Run(ctx)
	internal.Logger.Fatal().Err(err).Msg("")
}

func leaveDyingMessageOnPanic() {
	if err := recover(); err != nil {
		internal.Logger.Error().Interface("recoveredErr", err).Msg("panic occurred; show stacktrace")
		for depth := 0; ; depth++ {
			_, filename, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			internal.Logger.Error().Str("filename", filename).Int("lineNumber", line).Msg("")
		}
		os.Exit(1)
	}
}
