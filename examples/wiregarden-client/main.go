package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/moznion/wiregarden/grpc/messages"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	defaultPort := uint(0)
	defaultHost := ""
	portUsage := "the port number to dial a gRPC over TCP server"
	hostUsage := "the host to dial a gRPC over TCP server"
	var port uint
	var host string
	flag.UintVar(&port, "port", defaultPort, portUsage)
	flag.UintVar(&port, "p", defaultPort, portUsage+" (shorthand)")
	flag.StringVar(&host, "host", defaultHost, hostUsage)
	flag.StringVar(&host, "H", defaultHost, hostUsage+" (shorthand)")
	flag.Parse()

	if port == defaultPort {
		log.Fatal().Msg("mandatory parameter `port` is missing")
	}
	if host == defaultHost {
		log.Fatal().Msg("mandatory parameter `host` is missing")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal().Err(err).Msg("did not connect: %v")
	}
	defer conn.Close()
	peersClient := messages.NewPeersClient(conn)
	devicesClient := messages.NewDevicesClient(conn)

	ctx := context.Background()

	getDevices(ctx, devicesClient)
	getPeers(ctx, peersClient)
	updatePrivateKey(ctx, devicesClient)
}

func getDevices(ctx context.Context, devicesClient messages.DevicesClient) {
	resp, err := devicesClient.GetDevices(ctx, &messages.GetDevicesRequest{})
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	for _, device := range resp.Devices {
		fmt.Printf("%#v\n", device)
	}
}

func getPeers(ctx context.Context, peersClient messages.PeersClient) {
	resp, err := peersClient.GetPeers(ctx, &messages.GetPeersRequest{
		DeviceName: "wg0",
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	for _, peer := range resp.Peers {
		fmt.Printf("%#v\n", peer)
	}
}

func registerPeers(ctx context.Context, peersClient messages.PeersClient) {
	_, err := peersClient.RegisterPeers(ctx, &messages.RegisterPeersRequest{
		DeviceName: "wg0",
		Peers: []*messages.Peer{
			{
				PublicKey:       "<snip>",
				AllowedIps:      []string{"192.0.2.100/32"},
				EndpointUdpType: messages.UDPNetworkType_udp4,
				Endpoint:        "192.0.100.100:54321",
			},
		},
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Info().Msg("registered")
}

func deletePeers(ctx context.Context, peersClient messages.PeersClient) {
	_, err := peersClient.DeletePeers(ctx, &messages.DeletePeersRequest{
		DeviceName: "wg0",
		PublicKeys: []string{"<snip>"},
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Info().Msg("removed")
}

func updatePrivateKey(ctx context.Context, devicesClient messages.DevicesClient) {
	_, err := devicesClient.UpdatePrivateKey(ctx, &messages.UpdatePrivateKeyRequest{
		Name:       "wg0",
		PrivateKey: "<snip>",
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Info().Msg("private key updated")
}
