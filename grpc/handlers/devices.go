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
	gotDevices, err := func() ([]*wgtypes.Device, error) {
		if req.Name == "" {
			return h.deviceService.GetDevices(req.FilterPublicKeys)
		}

		gotDevice, err := h.deviceService.GetDevice(req.Name, req.FilterPublicKeys)
		if err != nil {
			return nil, err
		}
		if gotDevice == nil {
			return []*wgtypes.Device{}, nil
		}
		return []*wgtypes.Device{gotDevice}, nil
	}()
	if err != nil {
		log.Printf("[error] %s", err)
		return nil, status.Errorf(codes.Internal, "failed to collect the devices")
	}

	devices := make([]*messages.Device, len(gotDevices))
	for i, gotDevice := range gotDevices {
		devices[i] = messages.ConvertFromWgctrlDevice(gotDevice)
	}

	return &messages.GetDevicesResponse{
		Devices: devices,
	}, nil
}
