package client

import(
  "log"

  "github.com/matiasinsaurralde/loudp2p/peer"
)

var(
  InitialPeerList = []string{
    "127.0.0.1:5555",
    "127.0.0.1:5556",
  }
)

type Client struct {
  Peers []loudp2p.Peer
}

func NewClient() Client {
  log.Println("Starting client...")
  client := Client{
    Peers: make([]loudp2p.Peer, 0),
  }

  for _, peerAddr := range InitialPeerList {
    peer := loudp2p.Peer{peerAddr}
    client.Peers = append(client.Peers, peer)
  }

  return client
}

func(c *Client) StartDiscovery() {
  log.Println("Discovery starts.")

  for _, peer := range c.Peers {
    c.AnnounceTo(peer)
  }
}

func(c *Client) AnnounceTo(peer loudp2p.Peer) {
  log.Println("Announcing to:", peer)
}
