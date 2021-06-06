package e2etest

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	wiregardenGrpc "github.com/moznion/wiregarden/grpc"
	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
)

var log zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log = zerolog.New(output).With().Timestamp().Logger()
}

func TestRun(t *testing.T) {
	if !shouldExecute() {
		log.Info().Msg("skips the E2E testing. if you'd like to enable the E2E testing, please set the environment value `E2E_TEST` with non-empty value")
		return
	}

	listener, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	s := wiregardenGrpc.Server{
		Port: uint16(port),
	}
	defer s.Stop()
	go func() {
		ctx := context.Background()
		_ = s.Run(ctx)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure(), grpc.WithBlock())
	assert.NoError(t, err)
	defer func() {
		_ = conn.Close()
	}()

	peersClient := messages.NewPeersClient(conn)
	devicesClient := messages.NewDevicesClient(conn)

	testDevicesUpdatePrivateKey(t, devicesClient)
	testPeersRegisterPeersAndRemovePeers(t, peersClient)
}

func testDevicesUpdatePrivateKey(t *testing.T, devicesClient messages.DevicesClient) {
	log.Info().Msg("Devices UpdatePrivateKey")

	numOfAttempts := 3
	for i := 0; i < numOfAttempts; i++ {
		wgPrivateKey, err := exec.Command("wg", "genkey").Output()
		assert.NoError(t, err)
		wgPublicKey, err := exec.Command("bash", "-c", fmt.Sprintf(`echo "%s" | wg pubkey`, wgPrivateKey)).Output()
		assert.NoError(t, err)
		log.Debug().Str("publicKey", string(wgPublicKey)).Send()

		ctx := context.Background()
		deviceName := "wg0"

		_, err = devicesClient.UpdatePrivateKey(ctx, &messages.UpdatePrivateKeyRequest{
			Name:       deviceName,
			PrivateKey: strings.TrimSpace(string(wgPrivateKey)),
		})
		assert.NoError(t, err)

		devicesResp, err := devicesClient.GetDevices(ctx, &messages.GetDevicesRequest{
			Name: deviceName,
		})
		assert.NoError(t, err)
		devices := devicesResp.GetDevices()
		assert.Len(t, devices, 1)
		assert.Equal(t, strings.TrimSpace(string(wgPublicKey)), devices[0].GetPublicKey())
	}
}

func testPeersRegisterPeersAndRemovePeers(t *testing.T, peersClient messages.PeersClient) {
	log.Info().Msg("Peers RegisterPeers")

	ctx := context.Background()
	deviceName := "wg0"

	psk, err := exec.Command("wg", "genkey").Output()
	assert.NoError(t, err)

	peer1 := &messages.Peer{
		PublicKey:       generateWgPublicKey(t),
		AllowedIps:      []string{"192.0.2.10/32"},
		EndpointUdpType: messages.UDPNetworkType_udp,
		Endpoint:        "198.51.100.10:51820",
	}
	peer2 := &messages.Peer{
		PublicKey:                          generateWgPublicKey(t),
		AllowedIps:                         []string{"192.0.2.20/32", "192.0.2.30/32"},
		EndpointUdpType:                    messages.UDPNetworkType_udp4,
		Endpoint:                           "198.51.100.20:51820",
		PresharedKey:                       strings.TrimSpace(string(psk)),
		PersistentKeepaliveIntervalSeconds: 30,
	}

	_, err = peersClient.RegisterPeers(ctx, &messages.RegisterPeersRequest{
		DeviceName: deviceName,
		Peers:      []*messages.Peer{peer1, peer2},
	})
	assert.NoError(t, err)

	peersResp, err := peersClient.GetPeers(ctx, &messages.GetPeersRequest{
		DeviceName:       "wg0",
		FilterPublicKeys: nil,
	})
	assert.NoError(t, err)
	peers := peersResp.GetPeers()
	assert.Len(t, peers, 2)

	for _, d := range []struct {
		actual   *messages.Peer
		expected *messages.Peer
	}{
		{
			actual:   peers[0],
			expected: peer1,
		},
		{
			actual:   peers[1],
			expected: peer2,
		},
	} {
		log.Debug().Str("publicKey", d.actual.GetPublicKey()).Send()
		assert.Equal(t, d.expected.GetPublicKey(), d.actual.GetPublicKey())
		log.Debug().Strs("allowedIps", d.actual.GetAllowedIps()).Send()
		assert.Equal(t, d.expected.GetAllowedIps(), d.actual.GetAllowedIps())
		log.Debug().Str("endpoint", d.actual.GetEndpoint()).Send()
		assert.Equal(t, d.expected.GetEndpoint(), d.actual.GetEndpoint())
		log.Debug().Uint64("protocolVersion", d.actual.GetProtocolVersion()).Send()
		assert.Equal(t, d.expected.GetProtocolVersion(), d.actual.GetProtocolVersion())
		log.Debug().Uint32("persistentKeepaliveInterval", d.actual.GetPersistentKeepaliveIntervalSeconds()).Send()
		assert.Equal(t, d.expected.GetPersistentKeepaliveIntervalSeconds(), d.actual.GetPersistentKeepaliveIntervalSeconds())
		log.Debug().Str("presharedKey", d.actual.GetPresharedKey()).Send()
	}
	assert.Equal(t, peers[0].GetPresharedKey(), "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")
	assert.Equal(t, peers[1].GetPresharedKey(), peer2.GetPresharedKey())

	_, err = peersClient.DeletePeers(ctx, &messages.DeletePeersRequest{
		DeviceName: deviceName,
		PublicKeys: []string{peer1.GetPublicKey(), peer2.GetPublicKey()},
	})
	assert.NoError(t, err)
	peersResp, err = peersClient.GetPeers(ctx, &messages.GetPeersRequest{
		DeviceName:       "wg0",
		FilterPublicKeys: nil,
	})
	assert.NoError(t, err)
	assert.Len(t, peersResp.GetPeers(), 0)
}

func shouldExecute() bool {
	return os.Getenv("E2E_TEST") != ""
}

func generateWgPublicKey(t *testing.T) string {
	privateKey, err := wgtypes.GeneratePrivateKey()
	assert.NoError(t, err)
	return privateKey.PublicKey().String()
}
