package service

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Peer struct {
	deviceService *Device
}

func NewPeer(device *Device) *Peer {
	return &Peer{
		deviceService: device,
	}
}

func (p *Peer) GetPeers(deviceName string, filterPublicKeys []string) ([]wgtypes.Peer, error) {
	device, err := p.deviceService.GetDevice(deviceName, []string{})
	if err != nil {
		return nil, err
	}

	if device == nil {
		return []wgtypes.Peer{}, nil
	}

	if len(filterPublicKeys) <= 0 {
		// no filter
		return device.Peers, nil
	}

	filterPublicKeysMap := make(map[string]bool, len(filterPublicKeys))
	for _, key := range filterPublicKeys {
		filterPublicKeysMap[key] = true
	}

	var peers []wgtypes.Peer
	for _, peer := range device.Peers {
		if filterPublicKeysMap[peer.PublicKey.String()] {
			peers = append(peers, peer)
		}
	}
	return peers, nil
}
