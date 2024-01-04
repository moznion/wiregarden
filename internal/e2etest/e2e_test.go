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
	"github.com/moznion/wiregarden/routes"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log = zerolog.New(output).With().Timestamp().Logger()
}

func TestDevices_UpdatePrivateKey(t *testing.T) {
	assertSkip(t)

	deviceName := "wg0"
	flushWireGuardConfig(t, deviceName)

	conn := setupServerAndConn(t, nil)
	devicesClient := messages.NewDevicesClient(conn)

	numOfAttempts := 3
	for i := 0; i < numOfAttempts; i++ {
		wgPrivateKey, err := wgtypes.GeneratePrivateKey()
		assert.NoError(t, err)
		wgPublicKey := wgPrivateKey.PublicKey()
		log.Debug().Str("publicKey", wgPublicKey.String()).Send()

		ctx := context.Background()

		_, err = devicesClient.UpdatePrivateKey(ctx, &messages.UpdatePrivateKeyRequest{
			Name:       deviceName,
			PrivateKey: strings.TrimSpace(wgPrivateKey.String()),
		})
		assert.NoError(t, err)

		devicesResp, err := devicesClient.GetDevices(ctx, &messages.GetDevicesRequest{
			Name: deviceName,
		})
		assert.NoError(t, err)
		devices := devicesResp.GetDevices()
		assert.Len(t, devices, 1)
		assert.Equal(t, strings.TrimSpace(wgPublicKey.String()), devices[0].GetPublicKey())
	}

	flushWireGuardConfig(t, deviceName)
}

func TestPeers_RegisterPeersAndRemovePeers(t *testing.T) {
	assertSkip(t)

	deviceName := "wg0"
	flushWireGuardConfig(t, deviceName)

	conn := setupServerAndConn(t, nil)
	peersClient := messages.NewPeersClient(conn)

	ctx := context.Background()

	psk, err := wgtypes.GenerateKey()
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
		PresharedKey:                       strings.TrimSpace(psk.String()),
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

	flushWireGuardConfig(t, deviceName)
}

func TestPeers_RegisterPeersAndRemovePeersWithIPCmdRouter(t *testing.T) {
	assertSkip(t)

	deviceName := "wg0"
	flushIpRoute(t, deviceName)
	flushWireGuardConfig(t, deviceName)

	conn := setupServerAndConn(t, routes.IPRouterFrom(routes.IPRoutingPolicyIpcmd))
	peersClient := messages.NewPeersClient(conn)

	ctx := context.Background()

	ipaddr := "192.0.2.10"
	peer1 := &messages.Peer{
		PublicKey:       generateWgPublicKey(t),
		AllowedIps:      []string{ipaddr + "/32"},
		EndpointUdpType: messages.UDPNetworkType_udp,
		Endpoint:        "198.51.100.10:51820",
	}

	_, err := peersClient.RegisterPeers(ctx, &messages.RegisterPeersRequest{
		DeviceName: deviceName,
		Peers:      []*messages.Peer{peer1},
	})
	assert.NoError(t, err)

	output, err := exec.Command("ip", "route", "list", "dev", deviceName).Output()
	assert.NoError(t, err)
	assert.Equal(t, ipaddr, strings.Split(string(output), " ")[0], "route added for wg0")

	_, err = peersClient.DeletePeers(ctx, &messages.DeletePeersRequest{
		DeviceName: deviceName,
		PublicKeys: []string{peer1.GetPublicKey()},
	})
	assert.NoError(t, err)

	output, err = exec.Command("ip", "route", "list", "dev", deviceName).Output()
	assert.NoError(t, err)
	assert.Empty(t, output, "route removed for wg0")

	flushIpRoute(t, deviceName)
	flushWireGuardConfig(t, deviceName)
}

func shouldExecute() bool {
	return os.Getenv("E2E_TEST") != ""
}

func assertSkip(t *testing.T) {
	if !shouldExecute() {
		log.Info().Msg("skips the E2E testing. if you'd like to enable the E2E testing, please set the environment value `E2E_TEST` with non-empty value")
		t.Skip()
	}
}

func setupServerAndConn(t *testing.T, ipRouter routes.IPRouter) *grpc.ClientConn {
	listener, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	_ = listener.Close()

	s := wiregardenGrpc.Server{
		Port:     uint16(port),
		IPRouter: ipRouter,
	}
	t.Cleanup(func() {
		s.Stop()
	})

	go func() {
		ctx := context.Background()
		_ = s.Run(ctx)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	assert.NoError(t, err)
	t.Cleanup(func() {
		_ = conn.Close()
	})

	return conn
}

func flushIpRoute(t *testing.T, deviceName string) {
	err := exec.Command("ip", "route", "flush", "dev", deviceName).Run()
	assert.NoError(t, err)
}

func flushWireGuardConfig(t *testing.T, deviceName string) {
	_ = exec.Command("wg-quick", "down", deviceName).Run()
	err := exec.Command("wg-quick", "up", deviceName).Run()
	assert.NoError(t, err)
}

func generateWgPublicKey(t *testing.T) string {
	privateKey, err := wgtypes.GeneratePrivateKey()
	assert.NoError(t, err)
	return privateKey.PublicKey().String()
}
