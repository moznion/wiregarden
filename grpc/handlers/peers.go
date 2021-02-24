package handlers

import (
	"context"

	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Peers struct {
	deviceService *service.Device
	messages.UnimplementedPeersServer
}

func NewPeers(deviceService *service.Device) *Peers {
	return &Peers{
		deviceService: deviceService,
	}
}

func (h *Peers) GetPeers(ctx context.Context, req *messages.GetPeersRequest) (*messages.GetPeersResponse, error) {
	if req.DeviceName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}

	device, err := h.deviceService.GetDevice(req.DeviceName, req.FilterPublicKeys)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to collect the devices")
	}

	if len(req.FilterPublicKeys) <= 0 {
		peers := make([]*messages.Peer, len(device.Peers))
		for i, peer := range device.Peers {
			peers[i] = messages.ConvertFromWgctrlPeer(&peer)
		}
		return &messages.GetPeersResponse{Peers: peers}, nil
	}

	// FIXME duplicated code
	filterPublicKeysMap := make(map[string]bool, len(req.FilterPublicKeys))
	for _, key := range req.FilterPublicKeys {
		filterPublicKeysMap[key] = true
	}

	var peers []*messages.Peer
	for _, peer := range device.Peers {
		if filterPublicKeysMap[peer.PublicKey.String()] {
			peers = append(peers, messages.ConvertFromWgctrlPeer(&peer))
		}
	}

	return &messages.GetPeersResponse{Peers: peers}, nil
}

func (h *Peers) UpdatePeers(context.Context, *messages.UpdatePeersRequest) (*messages.UpdatePeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeers not implemented")
}

func (h *Peers) DeletePeers(context.Context, *messages.DeletePeersRequest) (*messages.DeletePeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeers not implemented")
}
