package loudp2p

import (
	"log"
)

var (
	initialPeerList = []string{
		"127.0.0.1:5555",
		"127.0.0.1:5556",
	}
)

// Client represents a standard data structure for clients.
type Client struct {
	Peers []Peer
}

// NewClient initializes a new client, using the data from settings.
func NewClient(settings *Settings) Client {
	// log.Println("Starting client.")
	client := Client{
		Peers: make([]Peer, 0),
	}

	for _, peerAddr := range initialPeerList {
		peer := Peer{peerAddr}
		client.Peers = append(client.Peers, peer)
	}

	return client
}

// StartDiscovery starts the peer discovery process.
func (c *Client) StartDiscovery() {
	log.Println("Discovery starts.")

	for _, peer := range c.Peers {
		c.AnnounceTo(peer)
	}
}

// AnnounceTo is used to announce a client to another peer.
func (c *Client) AnnounceTo(peer Peer) {
	log.Println("Announcing to:", peer)
}
