package loudp2p

import (
	loud "github.com/matiasinsaurralde/loudp2p/pb"
	"google.golang.org/grpc"
)

// Peer represents the standard data structure for peers.
type Peer struct {
	Address string
	Conn    *grpc.ClientConn
	Client  loud.LoudClient
}
