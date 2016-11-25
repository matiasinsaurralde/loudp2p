package main

import(
  "log"

  loudp2p "github.com/matiasinsaurralde/loudp2p/core"
  crypto  "github.com/matiasinsaurralde/loudp2p/crypto"
)

func main() {
  log.Println("loudp2p starts")

  var err error

  var settings *loudp2p.Settings
  settings = loudp2p.LoadSettings()

  if settings == nil {
    log.Println("No keys present, generating.")
    privKey, pubKey, peerId := crypto.GenerateKeys()
    settings = &loudp2p.Settings{
      PrivKey: privKey,
      PubKey: pubKey,
      PeerId: peerId,
    }
    err = settings.Persist()
    if err != nil {
      log.Println("Couldn't persist settings!")
      panic(err)
    }
  } else {
    log.Println("Using existing keys.")
  }

  log.Println("Peer ID is", settings.PeerId)

  client := loudp2p.NewClient(settings)
  server := loudp2p.NewServer(settings)

  log.Println(1,client,server)

  // go client.StartDiscovery()
  // go server.Start()
  for {}
}
