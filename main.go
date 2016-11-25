package main

import(
  "log"

  loudp2p "github.com/matiasinsaurralde/loudp2p/core"
)

var client *loudp2p.Client

func main() {
  log.Println("Init...")
  client := loudp2p.NewClient()
  server := loudp2p.NewServer()
  go client.StartDiscovery()
  go server.Start()
  for {}
}
