package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/grpc/metrics"
	"github.com/moznion/wiregarden/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Peers struct {
	peerService               *service.Peer
	peersRegistrationHooks    []PeersRegistrationHook
	peersDeletionHooks        []PeersDeletionHook
	prometheusMetricsRegister metrics.PrometheusMetricsRegisterable
	messages.UnimplementedPeersServer
}

// PeersRegistrationHook is an interface that defines the hook function to do when the peers' registration has done successfully.
type PeersRegistrationHook interface {
	Do(ctx context.Context, req *messages.RegisterPeersRequest) error
}

// PeersDeletionHook is an interface that defines the hook function to do when the peers' deletion has done successfully.
type PeersDeletionHook interface {
	Do(ctx context.Context, req *messages.DeletePeersRequest) error
}

func NewPeers(
	peerService *service.Peer,
	peersRegistrationHooks []PeersRegistrationHook,
	peersDeletionHooks []PeersDeletionHook,
	prometheusMetricsRegister metrics.PrometheusMetricsRegisterable,
) *Peers {
	return &Peers{
		peerService:               peerService,
		peersRegistrationHooks:    peersRegistrationHooks,
		peersDeletionHooks:        peersDeletionHooks,
		prometheusMetricsRegister: prometheusMetricsRegister,
	}
}

func (h *Peers) GetPeers(ctx context.Context, req *messages.GetPeersRequest) (*messages.GetPeersResponse, error) {
	const requestName = "get-peers"

	l := log.With().
		Str("requestID", uuid.NewString()).
		Str("request", requestName).
		Str("deviceName", req.DeviceName).
		Strs("filterPublicKeys", req.FilterPublicKeys).
		Logger()
	l.Info().Msg("received request")

	h.prometheusMetricsRegister.IncrementRequestCount(requestName)

	resp, errStatus := h.getPeers(ctx, &l, req)
	if errStatus != nil {
		l.Info().Uint32("statusCode", uint32(errStatus.Code())).Msgf("return not successfully: %s", errStatus.Message())
		h.prometheusMetricsRegister.IncrementFailureResponseCount(requestName, uint32(errStatus.Code()))
		return nil, errStatus.Err()
	}

	h.prometheusMetricsRegister.IncrementSuccessResponseCount(requestName)

	l.Info().Msg("return successfully")
	return resp, nil
}

func (h *Peers) getPeers(ctx context.Context, l *zerolog.Logger, req *messages.GetPeersRequest) (*messages.GetPeersResponse, *status.Status) {
	deviceName := req.DeviceName
	if deviceName == "" {
		return nil, status.Newf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}

	gotPeers, err := h.peerService.GetPeers(ctx, deviceName, req.FilterPublicKeys)
	if err != nil {
		l.Error().Err(err).Msg("failed to collect the peers")
		return nil, status.Newf(codes.Internal, "failed to collect the peers")
	}

	peers := make([]*messages.Peer, len(gotPeers))
	for i := range gotPeers {
		peers[i] = messages.ConvertFromWgctrlPeer(&gotPeers[i])
	}

	return &messages.GetPeersResponse{Peers: peers}, nil
}

func (h *Peers) RegisterPeers(ctx context.Context, req *messages.RegisterPeersRequest) (*messages.RegisterPeersResponse, error) {
	const requestName = "register-peers"

	l := log.With().
		Str("requestID", uuid.NewString()).
		Str("request", requestName).
		Str("deviceName", req.DeviceName).
		Bytes("hooksPayload", req.HooksPayload).
		Interface("peers", req.Peers).
		Logger()
	l.Info().Msg("received request")

	h.prometheusMetricsRegister.IncrementRequestCount(requestName)

	resp, errStatus := h.registerPeers(ctx, &l, req)
	if errStatus != nil {
		l.Info().Uint32("statusCode", uint32(errStatus.Code())).Msgf("return not successfully: %s", errStatus.Message())
		h.prometheusMetricsRegister.IncrementFailureResponseCount(requestName, uint32(errStatus.Code()))
		return nil, errStatus.Err()
	}

	h.prometheusMetricsRegister.IncrementSuccessResponseCount(requestName)

	l.Info().Msg("return successfully")
	return resp, nil
}

func (h *Peers) registerPeers(ctx context.Context, l *zerolog.Logger, req *messages.RegisterPeersRequest) (*messages.RegisterPeersResponse, *status.Status) {
	deviceName := req.DeviceName
	if deviceName == "" {
		return nil, status.Newf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}

	peers := make([]wgtypes.Peer, len(req.Peers))
	for i, reqPeer := range req.Peers {
		peer, err := reqPeer.ToWgctrlPeer()
		if err != nil {
			// TODO more client friendly error message
			return nil, status.Newf(codes.InvalidArgument, "peers parameter contains invalid parameter")
		}
		peers[i] = *peer
	}

	err := h.peerService.RegisterPeers(ctx, deviceName, peers)
	if err != nil {
		errMsg := "failed to register peers"
		l.Error().Err(err).Msg(errMsg)
		return nil, status.Newf(codes.Internal, errMsg)
	}

	for _, hook := range h.peersRegistrationHooks {
		err := hook.Do(ctx, req)
		if err != nil {
			errMsg := "failed to do a hook on peers registered, but peers registration has been succeeded"
			l.Error().Err(err).Msg(errMsg)
			return nil, status.Newf(codes.Unknown, errMsg)
		}
	}

	return &messages.RegisterPeersResponse{}, nil
}

func (h *Peers) DeletePeers(ctx context.Context, req *messages.DeletePeersRequest) (*messages.DeletePeersResponse, error) {
	const requestName = "delete-peers"

	l := log.With().
		Str("requestID", uuid.NewString()).
		Str("request", requestName).
		Str("deviceName", req.DeviceName).
		Strs("publicKeys", req.PublicKeys).
		Bytes("hooksPayload", req.HooksPayload).
		Logger()
	l.Info().Msg("received request")

	h.prometheusMetricsRegister.IncrementRequestCount(requestName)

	resp, errStatus := h.deletePeers(ctx, &l, req)
	if errStatus != nil {
		l.Info().Uint32("statusCode", uint32(errStatus.Code())).Msgf("return not successfully: %s", errStatus.Message())
		h.prometheusMetricsRegister.IncrementFailureResponseCount(requestName, uint32(errStatus.Code()))
		return nil, errStatus.Err()
	}

	h.prometheusMetricsRegister.IncrementSuccessResponseCount(requestName)

	l.Info().Msg("return successfully")
	return resp, nil
}

func (h *Peers) deletePeers(ctx context.Context, l *zerolog.Logger, req *messages.DeletePeersRequest) (*messages.DeletePeersResponse, *status.Status) {
	deviceName := req.DeviceName
	if deviceName == "" {
		return nil, status.Newf(codes.InvalidArgument, "device_name is a mandatory parameter, but missing")
	}
	publicKeys := req.PublicKeys
	if len(publicKeys) <= 0 {
		return nil, status.Newf(codes.InvalidArgument, "public_keys parameter is mandatory, but missing or empty")
	}

	err := h.peerService.DeletePeers(ctx, deviceName, publicKeys)
	if err != nil {
		errMsg := "failed to delete peers"
		l.Error().Err(err).Msg(errMsg)
		return nil, status.Newf(codes.Internal, errMsg)
	}

	for _, hook := range h.peersDeletionHooks {
		err := hook.Do(ctx, req)
		if err != nil {
			errMsg := "failed to do a hook on peers deleted, but peers deletion has been succeeded"
			l.Error().Err(err).Msg(errMsg)
			return nil, status.Newf(codes.Unknown, errMsg)
		}
	}

	return &messages.DeletePeersResponse{}, nil
}
