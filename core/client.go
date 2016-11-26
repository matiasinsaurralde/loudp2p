package loudp2p

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"

	pb "github.com/matiasinsaurralde/loudp2p/pb"
)

var (
	initialPeerList = []string{
		"127.0.0.1:5555",
	}
	dialOpts []grpc.DialOption
)

// Client represents a standard data structure for clients.
type Client struct {
	Peers    []Peer
	Settings *Settings
}

// NewClient initializes a new client, using the data from settings.
func NewClient(settings *Settings) (client Client, err error) {
	if settings == nil {
		err = errors.New("Invalid settings")
		return client, err
	}

	err = settings.Validate()
	if err != nil {
		return client, err
	}

	// Set global gRPC connection options:
	dialOpts = append(dialOpts, grpc.WithInsecure())

	client = Client{
		Peers:    make([]Peer, 0),
		Settings: settings,
	}

	for _, peerAddr := range initialPeerList {
		peer := Peer{
			Address: peerAddr,
		}
		client.Peers = append(client.Peers, peer)
	}

	return client, err
}

// Start starts the peer discovery process.
func (c *Client) Start() {
	log.Println("Discovery starts.")

	for _, peer := range c.Peers {
		go c.AnnounceTo(peer)
	}
}

// AnnounceTo is used to announce a client to another peer.
func (c *Client) AnnounceTo(peer Peer) {
	log.Println("Announcing to:", peer)
	var err error
	peer.Conn, err = grpc.Dial(peer.Address, dialOpts...)
	if err != nil {
		log.Println(peer, "couldn't establish connection!")
		return
	}

	peer.Client = pb.NewLoudClient(peer.Conn)
	helloMessage := pb.HelloMessage{
		Origin: c.Settings.PeerID,
	}

	peer.Client.Hello(context.Background(), &helloMessage)
}
