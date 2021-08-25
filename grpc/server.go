package grpc

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/moznion/wiregarden/grpc/handlers"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/grpc/metrics"
	"github.com/moznion/wiregarden/internal"
	"github.com/moznion/wiregarden/internal/service"
	"github.com/moznion/wiregarden/routes"
	"google.golang.org/grpc"
)

type Server struct {
	Port                      uint16
	UnixSocketPath            string
	IPRouter                  routes.IPRouter
	PeersRegistrationHooks    []handlers.PeersRegistrationHook
	PeersDeletionHooks        []handlers.PeersDeletionHook
	PrometheusMetricsRegister metrics.PrometheusMetricsRegisterable
	server                    *grpc.Server
}

func (s *Server) Run(ctx context.Context) error {
	listenerProducer := s.listenWithAddress
	if s.UnixSocketPath != "" {
		listenerProducer = s.listenWithUnixSocket
	}

	listener, cleanUpHandler, err := listenerProducer()
	if err != nil {
		return err
	}

	defer cleanUpHandler()

	s.server = grpc.NewServer()

	err = s.registerHandlers(s.server)
	if err != nil {
		return err
	}

	port := listener.Addr().(*net.TCPAddr).Port
	internal.Logger.Info().Int("port", port).Msg("start to listen gRPC over TCP")

	return s.server.Serve(listener)
}

func (s *Server) Stop() {
	internal.Logger.Info().Msg("received a gRPC server stopping instruction")
	if s.server == nil {
		internal.Logger.Info().Msg("received a gRPC server stopping instruction, but the server has already been missing; nothing to do")
		return
	}
	s.server.Stop()
	internal.Logger.Info().Msg("gRPC server stopped")
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
		internal.Logger.Info().Msg("no Prometheus metrics register is specified, it'll use the NOP register")
		s.PrometheusMetricsRegister = &metrics.NOPPrometheusMetricsRegister{}
	}

	messages.RegisterDevicesServer(grpcServer, handlers.NewDevices(deviceService, s.PrometheusMetricsRegister))
	messages.RegisterPeersServer(grpcServer, handlers.NewPeers(peerService, s.PeersRegistrationHooks, s.PeersDeletionHooks, s.PrometheusMetricsRegister))

	return nil
}

func (s *Server) listenWithUnixSocket() (net.Listener, func(), error) {
	_, err := os.Stat(s.UnixSocketPath)
	if err == nil {
		_ = os.Remove(s.UnixSocketPath)
	}

	listener, err := net.Listen("unix", s.UnixSocketPath)
	if err != nil {
		return nil, nil, err
	}

	err = os.Chmod(s.UnixSocketPath, 0600)
	if err != nil {
		return nil, nil, err
	}

	return listener, func() {
		err := os.Remove(s.UnixSocketPath)
		if err != nil {
			internal.Logger.Warn().Err(err).Str("unixSocketPath", s.UnixSocketPath).Msg("failed to clean up a UNIX socket on finish")
		}
	}, nil
}

func (s *Server) listenWithAddress() (net.Listener, func(), error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return nil, nil, err
	}
	return listener, func() {
		// NOP
	}, nil
}
