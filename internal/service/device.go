package service

import (
	"github.com/moznion/wiregarden/internal/infra"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Device struct {
	wgctrl *infra.WGCtrl
}

func NewDevice() (*Device, error) {
	wgctrl, err := infra.NewWGCtrl()
	if err != nil {
		return nil, err
	}

	return &Device{
		wgctrl: wgctrl,
	}, nil
}

func (d *Device) GetDevices(name string, filterPublicKeys []string) ([]*wgtypes.Device, error) {
	gotDevices, err := d.wgctrl.GetDevices(name, filterPublicKeys)
	if err != nil {
		return nil, err
	}
	return gotDevices, err
}
