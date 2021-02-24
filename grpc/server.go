package grpc

import (
	"fmt"
	"net"

	"github.com/moznion/wiregarden/grpc/handlers"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/internal/service"
	"google.golang.org/grpc"
)

type Server struct {
	Port uint16
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	messages.RegisterPeersServer(grpcServer, &handlers.Peers{})

	return grpcServer.Serve(lis)
}

func (s *Server) registerHandlers(grpcServer *grpc.Server) error {
	deviceService, err := service.NewDevice()
	if err != nil {
		return err
	}

	messages.RegisterPeersServer(grpcServer, handlers.NewPeers(deviceService))
	messages.RegisterDevicesServer(grpcServer, handlers.NewDevices(deviceService))

	return nil
}
