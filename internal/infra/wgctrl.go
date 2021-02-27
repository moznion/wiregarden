package infra

import (
	"context"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type WGCtrl struct {
	client *wgctrl.Client
}

func NewWGCtrl() (*WGCtrl, error) {
	client, err := wgctrl.New()
	if err != nil {
		return nil, err
	}
	return &WGCtrl{
		client: client,
	}, nil
}

func (w *WGCtrl) GetSingleDevice(ctx context.Context, name string, filterPublicKeys []string) ([]*wgtypes.Device, error) {
	gotDevice, err := w.client.Device(name)
	if err != nil {
		return nil, err
	}

	if len(filterPublicKeys) <= 0 {
		return []*wgtypes.Device{gotDevice}, nil
	}

	publicKeyFilterMap := w.publicKeysToFilterMap(filterPublicKeys)

	if publicKeyFilterMap[gotDevice.PublicKey.String()] {
		return []*wgtypes.Device{gotDevice}, nil
	}

	return []*wgtypes.Device{}, nil
}

func (w *WGCtrl) GetDevices(ctx context.Context, filterPublicKeys []string) ([]*wgtypes.Device, error) {
	gotDevices, err := w.client.Devices()
	if err != nil {
		return nil, err
	}

	if len(filterPublicKeys) <= 0 {
		return gotDevices, nil
	}

	publicKeyFilterMap := w.publicKeysToFilterMap(filterPublicKeys)

	var devices []*wgtypes.Device
	for _, device := range gotDevices {
		if publicKeyFilterMap[device.PublicKey.String()] {
			devices = append(devices, device)
		}
	}
	return devices, nil
}

func (w *WGCtrl) RegisterPeers(ctx context.Context, deviceName string, peers []wgtypes.PeerConfig) error {
	err := w.client.ConfigureDevice(deviceName, wgtypes.Config{
		ReplacePeers: false,
		Peers:        peers,
	})
	if err != nil {
		return err
	}
	return nil
}

// FIXME unify with RegisterPeers?
func (w *WGCtrl) DeletePeers(ctx context.Context, deviceName string, publicKeys []string) error {
	peersToRemove := make([]wgtypes.PeerConfig, len(publicKeys))

	for i, publicKey := range publicKeys {
		parsedKey, err := wgtypes.ParseKey(publicKey)
		if err != nil {
			return err
		}

		peersToRemove[i] = wgtypes.PeerConfig{
			PublicKey: parsedKey,
			Remove:    true,
		}
	}

	err := w.client.ConfigureDevice(deviceName, wgtypes.Config{
		Peers: peersToRemove,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *WGCtrl) publicKeysToFilterMap(publicKeys []string) map[string]bool {
	m := make(map[string]bool)
	for _, publicKey := range publicKeys {
		m[publicKey] = true
	}
	return m
}
