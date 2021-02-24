package messages

import (
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
