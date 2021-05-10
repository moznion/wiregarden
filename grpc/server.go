package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/moznion/wiregarden/grpc/handlers"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/internal/service"
	"github.com/moznion/wiregarden/routes"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Server struct {
	Port                   uint16
	IPRouter               routes.IPRouter
	PeersRegistrationHooks []handlers.PeersRegistrationHook
	PeersDeletionHooks     []handlers.PeersDeletionHook
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	err = s.registerHandlers(grpcServer)
	if err != nil {
		return err
	}

	port := listener.Addr().(*net.TCPAddr).Port
	log.Info().Int("port", port).Msg("start to listen gRPC over TCP")

	return grpcServer.Serve(listener)
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

	messages.RegisterDevicesServer(grpcServer, handlers.NewDevices(deviceService))
	messages.RegisterPeersServer(grpcServer, handlers.NewPeers(peerService, s.PeersRegistrationHooks, s.PeersDeletionHooks))

	return nil
}
