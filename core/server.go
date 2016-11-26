package loudp2p

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/matiasinsaurralde/loudp2p/pb"
	"google.golang.org/grpc"
)

// Server represents the standard data structure for servers.
type Server struct{}

// NewServer initializes a new server using the data from settings.
func NewServer(settings *Settings) (server Server, err error) {
	return server, err
}

func (s *Server) Hello(ctx context.Context, helloMessage *pb.HelloMessage) (message *pb.DummyMessage, err error) {
	log.Println("Receiving hello", helloMessage)
	message = &pb.DummyMessage{}
	return message, err
}

// Start starts the server.
func (s *Server) Start() {
	log.Println("Starting server...")

	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLoudServer(grpcServer, s)
	grpcServer.Serve(lis)

}
