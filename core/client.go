package loudp2p

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"strings"

	"github.com/golang/protobuf/ptypes"
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
	Events   *EventHandler
}

// NewClient initializes a new client, using the data from settings.
func NewClient(settings *Settings, events *EventHandler) (client Client, err error) {
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
		Events:   events,
	}

	if settings.IgnoreInitialPeers {
		log.Println("Ignoring initial peers.")
		return client, err
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
	go c.HandleEvents()

	for _, peer := range c.Peers {
		go c.AnnounceTo(peer)
	}
}

// Event handler:
func (c *Client) HandleEvents() {
	log.Println("Client is listening for events.")
	eventsChan := c.Events.AddListener()
	for {
		event := <-*eventsChan
		log.Println("Client gets event:", event)

		switch event.Type {
		case SayHello:
			log.Println("Client received a request to say hello to:", event.Metadata)

			eventData := event.Metadata.(map[int]interface{})

			var peerIP string
			peerIP = strings.Split(eventData[HelloPeerAddr].(string), ":")[0]

			var peerRPCPort int64
			peerRPCPort = eventData[HelloRPCPort].(int64)

			peer := Peer{
				Address: fmt.Sprintf("%s:%d", peerIP, peerRPCPort),
			}
			log.Println("peer", peer)
			// go c.AnnounceTo(peer)
		}
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
		Origin:  c.Settings.PeerID,
		RpcPort: c.Settings.RPCPort,
	}

	peer.Client.Hello(context.Background(), &helloMessage)

	dispatchClient, err := peer.Client.Dispatch(context.Background())

	if err != nil {
		log.Println("Couldn't announce!")
		return
	}

	go func(client pb.Loud_DispatchClient) {
		for {
			object, err := client.Recv()
			log.Println("Receiving:", object, err)
		}
	}(dispatchClient)

	go func(client pb.Loud_DispatchClient) {
		for {
			helloMessage := pb.HelloMessage{
				Origin:  c.Settings.PeerID,
				RpcPort: c.Settings.RPCPort,
			}
			out, err := ptypes.MarshalAny(&helloMessage)
			log.Println("Sending:", out, err)
			client.Send(out)
			time.Sleep(1 * time.Second)
			// client.Send(helloMessage)
		}
	}(dispatchClient)
}

func HandleExchange() {
}
