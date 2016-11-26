package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"log"

	"golang.org/x/crypto/ripemd160"

	"github.com/matiasinsaurralde/hellobitcoin/base58check"
)

// GenerateKeys will generate the peer initial keys and ID.
func GenerateKeys() (privateKeyBytes []byte, publicKeyBytes []byte, peerID string) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	privateKeyBytes, _ = x509.MarshalECPrivateKey(privateKey)
	publicKeyBytes, _ = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)

	shaHash := sha256.New()
	shaHash.Write(publicKeyBytes)
	hash := shaHash.Sum(nil)

	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(hash)
	hash = ripemd160Hash.Sum(nil)

	peerID = base58check.Encode("00", hash)

	return privateKeyBytes, publicKeyBytes, peerID
}

// ParseKeys will parse existing key buffers and return the appropiate data structures.
func ParseKeys(privateKeyBytes []byte, publicKeyBytes []byte) (privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, err error) {
	privateKey, err = x509.ParseECPrivateKey(privateKeyBytes)
	if err != nil {
		log.Println("Couldn't parse private key!")
		return nil, nil, err
	}
	var pub interface{}
	pub, err = x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		log.Println("Couldn't parse public key!")
		return nil, nil, err
	}
	publicKey = pub.(*ecdsa.PublicKey)
	return privateKey, publicKey, err
}
