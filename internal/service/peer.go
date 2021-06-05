package service

import (
	"context"

	"github.com/moznion/wiregarden/internal/infra"
	"github.com/moznion/wiregarden/routes"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Peer struct {
	wgctrl        *infra.WGCtrl
	deviceService *Device
	ipRouter      routes.IPRouter
}

func NewPeer(device *Device, ipRouter routes.IPRouter) (*Peer, error) {
	wgctrl, err := infra.NewWGCtrl()
	if err != nil {
		return nil, err
	}

	return &Peer{
		wgctrl:        wgctrl,
		deviceService: device,
		ipRouter:      ipRouter,
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
		psk := peer.PresharedKey
		keepaliveInterval := peer.PersistentKeepaliveInterval
		peerConfigurations[i] = wgtypes.PeerConfig{
			PublicKey:                   peer.PublicKey,
			Endpoint:                    peer.Endpoint,
			AllowedIPs:                  peer.AllowedIPs,
			PresharedKey:                &psk,
			PersistentKeepaliveInterval: &keepaliveInterval,
			Remove:                      false,
			UpdateOnly:                  false,
			ReplaceAllowedIPs:           true,
		}
	}

	err := p.wgctrl.RegisterPeers(ctx, deviceName, peerConfigurations)
	if err != nil {
		return err
	}

	if p.ipRouter != nil {
		for _, peer := range peers {
			for _, ip := range peer.AllowedIPs {
				err := p.ipRouter.AddRoute(ip.String(), deviceName)
				if err != nil {
					return err
				}
			}
		}
	}

	return err
}

func (p *Peer) DeletePeers(ctx context.Context, deviceName string, publicKeys []string) error {
	if p.ipRouter != nil {
		peers, err := p.GetPeers(ctx, deviceName, publicKeys)
		if err != nil {
			return err
		}
		for _, peer := range peers {
			for _, ip := range peer.AllowedIPs {
				err := p.ipRouter.DelRoute(ip.String(), deviceName)
				if err != nil {
					return err
				}
			}
		}
	}

	err := p.wgctrl.DeletePeers(ctx, deviceName, publicKeys)
	if err != nil {
		return err
	}
	return nil
}
