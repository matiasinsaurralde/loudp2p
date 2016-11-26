package main

import (
	"log"

	"crypto/ecdsa"

	loudp2p "github.com/matiasinsaurralde/loudp2p/core"
	crypto "github.com/matiasinsaurralde/loudp2p/crypto"
)

func main() {
	log.Println("loudp2p starts")

	var err error

	var settings *loudp2p.Settings
	settings = loudp2p.LoadSettings()

	if settings == nil {
		log.Println("No keys present, generating.")

		var privateKey *ecdsa.PrivateKey
		var publicKey *ecdsa.PublicKey
		var privateKeyBytes, publicKeyBytes []byte
		var peerID string

		privateKey, privateKeyBytes, publicKey, publicKeyBytes, peerID, err = crypto.GenerateKeys()

		if err != nil {
			log.Println("Couldn't generate keys!")
			panic(err)
		}

		settings = &loudp2p.Settings{
			PrivateKey:   privateKey,
			PublicKey:    publicKey,
			PrivKeyBytes: privateKeyBytes,
			PubKeyBytes:  publicKeyBytes,
			PeerID:       peerID,
		}
		err = settings.Persist()
		if err != nil {
			log.Println("Couldn't persist settings!")
			panic(err)
		}
	} else {
		log.Println("Using existing keys.")
	}

	log.Println("Peer ID is", settings.PeerID)

	var client loudp2p.Client
	var server loudp2p.Server

	client, err = loudp2p.NewClient(settings)
	if err != nil {
		log.Println("Couldn't initialize client!")
		panic(err)
	}

	server, err = loudp2p.NewServer(settings)
	if err != nil {
		log.Println("Couldn't initialize server!")
		panic(err)
	}

	// go client.StartDiscovery()
	// go server.Start()
	for {
	}
	log.Println(1, client, server)
}
