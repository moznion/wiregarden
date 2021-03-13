package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/moznion/wiregarden/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Devices struct {
	deviceService *service.Device
	messages.UnimplementedDevicesServer
}

func NewDevices(deviceService *service.Device) *Devices {
	return &Devices{
		deviceService: deviceService,
	}
}

func (h *Devices) GetDevices(ctx context.Context, req *messages.GetDevicesRequest) (*messages.GetDevicesResponse, error) {
	l := log.With().
		Str("requestID", uuid.NewString()).
		Str("request", "get-devices").
		Str("name", req.Name).
		Strs("filterPublicKeys", req.FilterPublicKeys).
		Logger()
	l.Info().Msg("received")

	resp, errStatus := h.getDevices(ctx, &l, req)
	if errStatus != nil {
		l.Info().Uint32("statusCode", uint32(errStatus.Code())).Msgf("return not successfully: %s", errStatus.Message())
		return nil, errStatus.Err()
	}

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
