package loudp2p

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/matiasinsaurralde/loudp2p/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// Server represents the standard data structure for servers.
type Server struct {
	Settings *Settings
	Events   *EventHandler
}

// NewServer initializes a new server using the data from settings.
func NewServer(settings *Settings, events *EventHandler) (server Server, err error) {
	server.Settings = settings
	server.Events = events
	return server, err
}

// Event handler:
func (s *Server) HandleEvents() {
	log.Println("Server is listening for events.")
	eventsChan := s.Events.AddListener()
	for {
		event := <-*eventsChan
		log.Println("Server gets event:", event)
	}
}

// Service implementation:
func (s *Server) Hello(ctx context.Context, helloMessage *pb.HelloMessage) (message *pb.DummyMessage, err error) {
	thisPeer, _ := peer.FromContext(ctx)
	log.Println("peer", thisPeer)
	log.Println("helloMessage", helloMessage)

	eventData := make(map[int]interface{})
	eventData[HelloPeerAddr] = thisPeer.Addr.String()
	eventData[HelloRPCPort] = helloMessage.RpcPort

	event := Event{
		Type:     SayHello,
		Metadata: eventData,
	}

	s.Events.Emit(event)

	message = &pb.DummyMessage{}
	return message, err
}

func (s *Server) Dispatch(stream pb.Loud_DispatchServer) (err error) {
	log.Println("Read dispatch...")
	go func(recvStream pb.Loud_DispatchServer) {
		for {
			m, err := recvStream.Recv()
			if err != nil {
				log.Println("Stream error?")
				break
			}
			log.Println("Receiving:", m)
		}
	}(stream)

	for {
		testMessage := pb.HelloMessage{"hello123", 1000}
		anyMessage, err := ptypes.MarshalAny(&testMessage)
		if err != nil {
			log.Println("Stream error?")
			break
		}
		log.Println("Sending:", anyMessage)
		stream.Send(anyMessage)
		time.Sleep(1 * time.Second)
	}
	return err
}

// Start starts the server.
func (s *Server) Start() {
	go s.HandleEvents()

	if s.Settings.RPCPort == 0 {
		s.Settings.RPCPort = DefaultRPCPort
	}

	log.Printf("Starting server on %d\n", s.Settings.RPCPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Settings.RPCPort))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLoudServer(grpcServer, s)
	grpcServer.Serve(lis)
}
