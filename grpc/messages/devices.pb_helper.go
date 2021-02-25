package messages

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func ConvertFromWgctrlDevice(device *wgtypes.Device) *Device {
	peers := make([]*Peer, len(device.Peers))
	for i := range device.Peers {
		peers[i] = ConvertFromWgctrlPeer(&device.Peers[i])
	}

	return &Device{
		Name:           device.Name,
		DeviceType:     uint32(device.Type),
		DeviceTypeName: device.Type.String(),
		PublicKey:      device.PublicKey.String(),
		ListenPort:     uint32(device.ListenPort),
		FirewallMark:   int64(device.FirewallMark),
		Peers:          peers,
	}
}
