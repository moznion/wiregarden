package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/moznion/wiregarden/grpc/handlers"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/grpc/metrics"
	"github.com/moznion/wiregarden/internal/service"
	"github.com/moznion/wiregarden/routes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Server struct {
	Port                      uint16
	IPRouter                  routes.IPRouter
	PeersRegistrationHooks    []handlers.PeersRegistrationHook
	PeersDeletionHooks        []handlers.PeersDeletionHook
	PrometheusMetricsRegister metrics.PrometheusMetricsRegisterable
	server                    *grpc.Server
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}

	s.server = grpc.NewServer()

	err = s.registerHandlers(s.server)
	if err != nil {
		return err
	}

	port := listener.Addr().(*net.TCPAddr).Port
	log.Info().Int("port", port).Msg("start to listen gRPC over TCP")

	return s.server.Serve(listener)
}

func (s *Server) Stop() {
	log.Info().Msg("received a gRPC server stopping instruction")
	if s.server == nil {
		log.Info().Msg("received a gRPC server stopping instruction, but the server has already been missing; nothing to do")
		return
	}
	s.server.Stop()
	log.Info().Msg("gRPC server stopped")
}

func (s *Server) registerHandlers(grpcServer *grpc.Server) error {
	deviceService, err := service.NewDevice()
	if err != nil {
		return err
	}
	peerService, err := service.NewPeer(deviceService, s.IPRouter)
	if err != nil {
		return err
	}

	if s.PrometheusMetricsRegister == nil {
		log.Info().Msg("no Prometheus metrics register is specified, it'll use the NOP register")
		s.PrometheusMetricsRegister = &metrics.NOPPrometheusMetricsRegister{}
	}

	messages.RegisterDevicesServer(grpcServer, handlers.NewDevices(deviceService, s.PrometheusMetricsRegister))
	messages.RegisterPeersServer(grpcServer, handlers.NewPeers(peerService, s.PeersRegistrationHooks, s.PeersDeletionHooks, s.PrometheusMetricsRegister))

	return nil
}
