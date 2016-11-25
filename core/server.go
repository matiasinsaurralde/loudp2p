package loudp2p

import(
  "log"
)

type Server struct {}

func NewServer(settings *Settings) Server {
  return Server{}
}

func(s *Server) Start() {
  log.Println("Starting server...")
}
