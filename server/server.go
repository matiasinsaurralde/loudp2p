package server

import(
  "log"
)

type Server struct {}

func NewServer() Server {
  return Server{}
}

func(s *Server) Start() {
  log.Println("Starting server...")
}
