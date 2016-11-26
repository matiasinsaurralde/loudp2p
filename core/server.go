package loudp2p

import (
	"log"
)

// Server represents the standard data structure for servers.
type Server struct{}

// NewServer initializes a new server using the data from settings.
func NewServer(settings *Settings) Server {
	return Server{}
}

// Start starts the server.
func (s *Server) Start() {
	log.Println("Starting server...")
}
