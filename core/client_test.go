package loudp2p

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"testing"
)

const (
	testPeerID = "1Jxm2g1B4AFdBNjcfXFsVjtiU2d748cS8a"
)

var (
	testPrivKey     *ecdsa.PrivateKey
	privateKeyBytes []byte
	publicKeyBytes  []byte

	testEventHandler EventHandler
)

func init() {
	testPrivKey, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	privateKeyBytes, _ = x509.MarshalECPrivateKey(testPrivKey)
	publicKeyBytes, _ = x509.MarshalPKIXPublicKey(&testPrivKey.PublicKey)

	testEventHandler = NewEventHandler()
}

func TestClientInitialization(t *testing.T) {
	var err error
	_, err = NewClient(nil, &testEventHandler)

	if err == nil {
		t.Fatal("A client with no settings must fail to initialize.")
	}

	var settings Settings
	settings = Settings{}
	_, err = NewClient(&settings, &testEventHandler)
	if err == nil {
		t.Fatal("A client with an empty settings data structure must fail to initialize.")
	}

	settings = Settings{PrivKeyBytes: privateKeyBytes}
	_, err = NewClient(&settings, &testEventHandler)
	if err == nil {
		t.Fatal("A client with no public key must fail to initialize.")
	}

	settings = Settings{PubKeyBytes: publicKeyBytes}
	_, err = NewClient(&settings, &testEventHandler)
	if err == nil {
		t.Fatal("A client with no private key must fail to initialize.")
	}

	settings = Settings{PrivKeyBytes: privateKeyBytes, PubKeyBytes: publicKeyBytes}
	_, err = NewClient(&settings, &testEventHandler)
	if err == nil {
		t.Fatal("A client with no peerID must fail to initialize.")
	}

	settings = Settings{PrivKeyBytes: privateKeyBytes, PubKeyBytes: publicKeyBytes, PeerID: testPeerID}
	settings.LoadKeys()
	_, err = NewClient(&settings, &testEventHandler)
	if err != nil {
		t.Fatal(err)
		t.Fatal("A client with a key pair and a peer ID should initialize correctly.")
	}

}
