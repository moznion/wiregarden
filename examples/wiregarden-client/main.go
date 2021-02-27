package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/moznion/wiregarden/grpc/messages"
	"google.golang.org/grpc"
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
	flag.StringVar(&host, "H", defaultHost, hostUsage)
	flag.Parse()

	if port == defaultPort {
		log.Fatal("mandatory parameter `port` is missing")
	}
	if host == defaultHost {
		log.Fatal("mandatory parameter `host` is missing")
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	peersClient := messages.NewPeersClient(conn)
	devicesClient := messages.NewDevicesClient(conn)

	getDevices(devicesClient)
	getPeers(peersClient)
}

func getDevices(devicesClient messages.DevicesClient) {
	resp, err := devicesClient.GetDevices(context.Background(), &messages.GetDevicesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range resp.Devices {
		fmt.Printf("%#v\n", device)
	}
}

func getPeers(peersClient messages.PeersClient) {
	resp, err := peersClient.GetPeers(context.Background(), &messages.GetPeersRequest{
		DeviceName: "wg0",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, peer := range resp.Peers {
		fmt.Printf("%#v\n", peer)
	}
}

func registerPeers(peersClient messages.PeersClient) {
	_, err := peersClient.RegisterPeers(context.Background(), &messages.RegisterPeersRequest{
		DeviceName: "wg0",
		Peers: []*messages.Peer{
			{
				PublicKey:       "8xGbima8VGlpt4D2+4YCIITeTSnfFtF5/zJYshu3oAY=",
				AllowedIps:      []string{"172.16.2.100/32"},
				EndpointUdpType: messages.UDPNetworkType_udp4,
				Endpoint:        "18.183.228.1:55556",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("registered")
}

func deletePeers(peersClient messages.PeersClient) {
	_, err := peersClient.DeletePeers(context.Background(), &messages.DeletePeersRequest{
		DeviceName: "wg0",
		PublicKeys: []string{"8xGbima8VGlpt4D2+4YCIITeTSnfFtF5/zJYshu3oAY="},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("removed")
}
