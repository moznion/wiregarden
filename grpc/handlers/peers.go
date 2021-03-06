package handlers

import (
	"context"
	"log"

	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/internal/service"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Peers struct {
	peerService            *service.Peer
	peersRegistrationHooks []PeersRegistrationHook
	peersDeletionHooks     []PeersDeletionHook
	messages.UnimplementedPeersServer
}

// PeersRegistrationHook is an interface that defines the hook function to do when the peers' registration has done successfully.
type PeersRegistrationHook interface {
	Do(req *messages.RegisterPeersRequest) error
}

// PeersDeletionHook is an interface that defines the hook function to do when the peers' deletion has done successfully.
type PeersDeletionHook interface {
	Do(req *messages.DeletePeersRequest) error
}

func NewPeers(
	peerService *service.Peer,
	peersRegistrationHooks []PeersRegistrationHook,
	peersDeletionHooks []PeersDeletionHook,
) *Peers {
	return &Peers{
		peerService:            peerService,
		peersRegistrationHooks: peersRegistrationHooks,
		peersDeletionHooks:     peersDeletionHooks,
	}
}

func (h *Peers) GetPeers(ctx context.Context, req *messages.GetPeersRequest) (*messages.GetPeersResponse, error) {
	deviceName := req.DeviceName
	if deviceName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}

	gotPeers, err := h.peerService.GetPeers(ctx, deviceName, req.FilterPublicKeys)
	if err != nil {
		log.Printf("[error] %s", err)
		return nil, status.Error(codes.Internal, "failed to collect the peers")
	}

	peers := make([]*messages.Peer, len(gotPeers))
	for i := range gotPeers {
		peers[i] = messages.ConvertFromWgctrlPeer(&gotPeers[i])
	}

	return &messages.GetPeersResponse{Peers: peers}, nil
}

func (h *Peers) RegisterPeers(ctx context.Context, req *messages.RegisterPeersRequest) (*messages.RegisterPeersResponse, error) {
	deviceName := req.DeviceName
	if deviceName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}

	peers := make([]wgtypes.Peer, len(req.Peers))
	for i, reqPeer := range req.Peers {
		peer, err := reqPeer.ToWgctrlPeer()
		if err != nil {
			return nil, err
		}
		peers[i] = *peer
	}

	err := h.peerService.RegisterPeers(ctx, deviceName, peers)
	if err != nil {
		log.Printf("[error] %s", err)
		return nil, status.Error(codes.Internal, "failed to register peers")
	}

	for _, hook := range h.peersRegistrationHooks {
		err := hook.Do(req)
		if err != nil {
			log.Printf("[error] %s", err)
			return nil, status.Error(codes.Unknown, "failed to do a hook on peers registered")
		}
	}

	return &messages.RegisterPeersResponse{}, nil
}

func (h *Peers) DeletePeers(ctx context.Context, req *messages.DeletePeersRequest) (*messages.DeletePeersResponse, error) {
	deviceName := req.DeviceName
	if deviceName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}
	publicKeys := req.PublicKeys
	if len(publicKeys) <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "public_keys parameter is mandatory, but missing or empty")
	}

	err := h.peerService.DeletePeers(ctx, deviceName, publicKeys)
	if err != nil {
		log.Printf("[error] %s", err)
		return nil, status.Error(codes.Internal, "failed to delete peers")
	}

	for _, hook := range h.peersDeletionHooks {
		err := hook.Do(req)
		if err != nil {
			log.Printf("[error] %s", err)
			return nil, status.Error(codes.Unknown, "failed to do a hook on peers deleted")
		}
	}

	return &messages.DeletePeersResponse{}, nil
}
