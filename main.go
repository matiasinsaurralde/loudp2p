package main

import(
  "log"

  loudp2pclient "github.com/matiasinsaurralde/loudp2p/client"
  loudp2pserver "github.com/matiasinsaurralde/loudp2p/server"
)

var client *loudp2pclient.Client

func main() {
  log.Println("Init...")
  client := loudp2pclient.NewClient()
  server := loudp2pserver.NewServer()
  go client.StartDiscovery()
  go server.Start()
  for {}
}
