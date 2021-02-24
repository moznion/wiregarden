package infra

import (
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

func (w *WGCtrl) GetDevices(name string, publicKeys []string) ([]*wgtypes.Device, error) {
	if name != "" {
		return w.getSingleDevice(name, publicKeys)
	}
	return w.getDevices(publicKeys)
}

func (w *WGCtrl) getSingleDevice(name string, filterPublicKeys []string) ([]*wgtypes.Device, error) {
	gotDevice, err := w.client.Device(name)
	if err != nil {
		return nil, err
	}

	if len(filterPublicKeys) <= 0 {
		return []*wgtypes.Device{gotDevice}, nil
	}

	publicKeyFilterMap := w.publicKeysToFilterMap(filterPublicKeys)

	if publicKeyFilterMap[gotDevice.PublicKey] {
		return []*wgtypes.Device{gotDevice}, nil
	}

	return []*wgtypes.Device{}, nil
}

func (w *WGCtrl) getDevices(filterPublicKeys []string) ([]*wgtypes.Device, error) {
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
		if publicKeyFilterMap[device.PublicKey] {
			devices = append(devices, device)
		}
	}
	return devices, nil
}

func (w *WGCtrl) publicKeysToFilterMap(publicKeys []string) map[[wgtypes.KeyLen]byte]bool {
	m := make(map[[wgtypes.KeyLen]byte]bool)
	for _, publicKey := range publicKeys {
		var publicKeyBytes [wgtypes.KeyLen]byte
		copy(publicKeyBytes[:], publicKey)
		m[publicKeyBytes] = true
	}
	return m
}
