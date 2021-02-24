package handlers

import (
	"context"

	"github.com/moznion/wiregarden/grpc/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Peers struct {
	messages.UnimplementedPeersServer
}

func (h *Peers) GetPeers(context.Context, *messages.GetPeersRequest) (*messages.GetPeersResponse, error) {
	peers := make([]*messages.Peer, 0)
	return &messages.GetPeersResponse{Peers: peers}, nil
}

func (h *Peers) UpdatePeers(context.Context, *messages.UpdatePeersRequest) (*messages.UpdatePeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeers not implemented")
}

func (h *Peers) DeletePeers(context.Context, *messages.DeletePeersRequest) (*messages.DeletePeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeers not implemented")
}
