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

type Devices struct {
	deviceService             *service.Device
	prometheusMetricsRegister metrics.PrometheusMetricsRegisterable
	messages.UnimplementedDevicesServer
}

func NewDevices(deviceService *service.Device, prometheusMetricsRegister metrics.PrometheusMetricsRegisterable) *Devices {
	return &Devices{
		deviceService:             deviceService,
		prometheusMetricsRegister: prometheusMetricsRegister,
	}
}

func (h *Devices) GetDevices(ctx context.Context, req *messages.GetDevicesRequest) (*messages.GetDevicesResponse, error) {
	const requestName = "get-devices"

	l := log.With().
		Str("requestID", uuid.NewString()).
		Str("request", requestName).
		Str("name", req.Name).
		Strs("filterPublicKeys", req.FilterPublicKeys).
		Logger()
	l.Info().Msg("received")

	h.prometheusMetricsRegister.IncrementRequestCount(requestName)

	resp, errStatus := h.getDevices(ctx, &l, req)
	if errStatus != nil {
		l.Info().Uint32("statusCode", uint32(errStatus.Code())).Msgf("return not successfully: %s", errStatus.Message())
		h.prometheusMetricsRegister.IncrementFailureResponseCount(requestName, uint32(errStatus.Code()))
		return nil, errStatus.Err()
	}

	h.prometheusMetricsRegister.IncrementSuccessResponseCount(requestName)

	l.Info().Msg("return successfully")
	return resp, nil
}

func (h *Devices) getDevices(ctx context.Context, l *zerolog.Logger, req *messages.GetDevicesRequest) (*messages.GetDevicesResponse, *status.Status) {
	gotDevices, err := func() ([]*wgtypes.Device, error) {
		if req.Name == "" {
			return h.deviceService.GetDevices(ctx, req.FilterPublicKeys)
		}

		gotDevice, err := h.deviceService.GetDevice(ctx, req.Name, req.FilterPublicKeys)
		if err != nil {
			return nil, err
		}
		if gotDevice == nil {
			return []*wgtypes.Device{}, nil
		}
		return []*wgtypes.Device{gotDevice}, nil
	}()
	if err != nil {
		errMsg := "failed to collect the devices"
		l.Error().Err(err).Msg(errMsg)
		return nil, status.Newf(codes.Internal, errMsg)
	}

	devices := make([]*messages.Device, len(gotDevices))
	for i, gotDevice := range gotDevices {
		devices[i] = messages.ConvertFromWgctrlDevice(gotDevice)
	}

	return &messages.GetDevicesResponse{
		Devices: devices,
	}, nil
}

func (h *Devices) UpdatePrivateKey(ctx context.Context, req *messages.UpdatePrivateKeyRequest) (*messages.UpdatePrivateKeyResponse, error) {
	const requestName = "update-private-key"

	l := log.With().
		Str("requestID", uuid.NewString()).
		Str("request", requestName).
		Logger()
	l.Info().Msg("received")

	h.prometheusMetricsRegister.IncrementRequestCount(requestName)

	resp, errStatus := h.updatePrivateKey(ctx, &l, req)

	if errStatus != nil {
		l.Info().Uint32("statusCode", uint32(errStatus.Code())).Msgf("update not successfully: %s", errStatus.Message())
		h.prometheusMetricsRegister.IncrementFailureResponseCount(requestName, uint32(errStatus.Code()))
		return nil, errStatus.Err()
	}

	h.prometheusMetricsRegister.IncrementSuccessResponseCount(requestName)

	l.Info().Msg("updated successfully")
	return resp, nil
}

func (h *Devices) updatePrivateKey(ctx context.Context, l *zerolog.Logger, req *messages.UpdatePrivateKeyRequest) (*messages.UpdatePrivateKeyResponse, *status.Status) {
	if req.Name == "" {
		return nil, status.Newf(codes.InvalidArgument, "`name` is mandatory parameter to update the private key")
	}

	err := h.deviceService.UpdatePrivateKey(ctx, req.Name, req.PrivateKey)
	if err != nil {
		switch err {
		case service.ErrInvalidPrivateKey:
			return nil, status.Newf(codes.InvalidArgument, "invalid private key")
		case service.ErrDeviceNotFound:
			return nil, status.Newf(codes.NotFound, "device not found")
		default:
			errMsg := "failed to update the private key"
			l.Err(err).Msg(errMsg)
			return nil, status.Newf(codes.Internal, errMsg)
		}
	}

	return &messages.UpdatePrivateKeyResponse{}, nil
}
