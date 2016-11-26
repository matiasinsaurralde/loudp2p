package loudp2p

import (
	"testing"
)

const (
	testPrivKey = "abc"
	testPubKey  = "abc"
	testPeerID  = "1Jxm2g1B4AFdBNjcfXFsVjtiU2d748cS8a"
)

func TestClientInitialization(t *testing.T) {
	var err error
	_, err = NewClient(nil)

	if err == nil {
		t.Fatal("A client with no settings must fail to initialize.")
	}

	var settings Settings
	settings = Settings{}
	_, err = NewClient(&settings)
	if err == nil {
		t.Fatal("A client with an empty settings data structure must fail to initialize.")
	}

	settings = Settings{PrivKey: []byte(testPrivKey)}
	_, err = NewClient(&settings)
	if err == nil {
		t.Fatal("A client with no public key must fail to initialize.")
	}

	settings = Settings{PubKey: []byte(testPubKey)}
	_, err = NewClient(&settings)
	if err == nil {
		t.Fatal("A client with no private key must fail to initialize.")
	}

	settings = Settings{PrivKey: []byte(testPrivKey), PubKey: []byte(testPubKey)}
	_, err = NewClient(&settings)
	if err == nil {
		t.Fatal("A client with no peerID must fail to initialize.")
	}

	settings = Settings{PrivKey: []byte(testPrivKey), PubKey: []byte(testPubKey), PeerID: testPeerID}
	_, err = NewClient(&settings)
	if err != nil {
		t.Fatal("A client with a key pair and a peer ID should initialize correctly.")
	}

}
