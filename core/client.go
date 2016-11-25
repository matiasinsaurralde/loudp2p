package loudp2p

import(
  "log"
)

var(
  InitialPeerList = []string{
    "127.0.0.1:5555",
    "127.0.0.1:5556",
  }
)

type Client struct {
  Peers []Peer
}

func NewClient() Client {
  log.Println("Starting client...")
  client := Client{
    Peers: make([]Peer, 0),
  }

  for _, peerAddr := range InitialPeerList {
    peer := Peer{peerAddr}
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

func(c *Client) AnnounceTo(peer Peer) {
  log.Println("Announcing to:", peer)
}
