package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/moznion/wiregarden/grpc/handlers"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/internal/service"
	"google.golang.org/grpc"
)

type Server struct {
	Port                   uint16
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
	log.Printf("start to listen gRPC over TCP; port = %d", port)

	return grpcServer.Serve(listener)
}

func (s *Server) registerHandlers(grpcServer *grpc.Server) error {
	deviceService, err := service.NewDevice()
	if err != nil {
		return err
	}
	peerService, err := service.NewPeer(deviceService)
	if err != nil {
		return err
	}

	messages.RegisterDevicesServer(grpcServer, handlers.NewDevices(deviceService))
	messages.RegisterPeersServer(grpcServer, handlers.NewPeers(peerService, s.PeersRegistrationHooks, s.PeersDeletionHooks))

	return nil
}
