package messages

import (
	"net"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func ConvertFromWgctrlPeer(peer *wgtypes.Peer) *Peer {
	allowedIPs := make([]string, len(peer.AllowedIPs))
	for i, ip := range peer.AllowedIPs {
		allowedIPs[i] = ip.String()
	}

	return &Peer{
		PublicKey:  peer.PublicKey.String(),
		AllowedIps: allowedIPs,
		Endpoint:   peer.Endpoint.String(),
	}
}

func (x *Peer) ToWgctrlPeer() (*wgtypes.Peer, error) {
	var err error

	var publicKey wgtypes.Key
	if len(x.PublicKey) > 0 {
		publicKey, err = wgtypes.ParseKey(x.PublicKey)
		if err != nil {
			return nil, err
		}
	}

	var presharedKey wgtypes.Key
	if len(x.PresharedKey) > 0 {
		presharedKey, err = wgtypes.ParseKey(x.PresharedKey)
		if err != nil {
			return nil, err
		}
	}

	var endpoint *net.UDPAddr
	if len(x.Endpoint) > 0 {
		endpoint, err = net.ResolveUDPAddr(x.EndpointUdpType.String(), x.Endpoint)
		if err != nil {
			return nil, err
		}
	}

	allowIPs := make([]net.IPNet, len(x.AllowedIps))
	for i, ipStr := range x.AllowedIps {
		_, ipNet, err := net.ParseCIDR(ipStr)
		if err != nil {
			return nil, err
		}
		allowIPs[i] = *ipNet
	}

	return &wgtypes.Peer{
		PublicKey:                   publicKey,
		PresharedKey:                presharedKey,
		Endpoint:                    endpoint,
		AllowedIPs:                  allowIPs,
		PersistentKeepaliveInterval: time.Duration(x.PersistentKeepaliveIntervalSeconds) * time.Second,
		ProtocolVersion:             int(x.ProtocolVersion),
	}, nil
}
