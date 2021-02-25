package service

import (
	"context"

	"github.com/moznion/wiregarden/internal/infra"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Peer struct {
	wgctrl        *infra.WGCtrl
	deviceService *Device
}

func NewPeer(device *Device) (*Peer, error) {
	wgctrl, err := infra.NewWGCtrl()
	if err != nil {
		return nil, err
	}

	return &Peer{
		wgctrl:        wgctrl,
		deviceService: device,
	}, nil
}

func (p *Peer) GetPeers(ctx context.Context, deviceName string, filterPublicKeys []string) ([]wgtypes.Peer, error) {
	device, err := p.deviceService.GetDevice(ctx, deviceName, []string{})
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

func (p *Peer) RegisterPeers(ctx context.Context, deviceName string, peers []wgtypes.Peer) error { // FIXME don't pass wgtypes.peer directly
	peerConfigurations := make([]wgtypes.PeerConfig, len(peers))
	for i, peer := range peers {
		peerConfigurations[i] = wgtypes.PeerConfig{
			PublicKey:         peer.PublicKey,
			Remove:            false,
			UpdateOnly:        false,
			Endpoint:          peer.Endpoint,
			ReplaceAllowedIPs: true,
			AllowedIPs:        peer.AllowedIPs,
		}
	}

	err := p.wgctrl.RegisterPeers(ctx, deviceName, peerConfigurations)
	if err != nil {
		return nil
	}
	return err
}
