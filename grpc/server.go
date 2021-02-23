package grpc

import (
	"fmt"
	"net"

	"github.com/moznion/wiregarden/grpc/handlers"
	"github.com/moznion/wiregarden/grpc/messages"
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

func (s *Server) registerHandlers(grpcServer *grpc.Server) {
	messages.RegisterPeersServer(grpcServer, &handlers.Peers{})
}
