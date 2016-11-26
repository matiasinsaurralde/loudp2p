package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"log"

	"golang.org/x/crypto/ripemd160"

	base58 "github.com/matiasinsaurralde/loudp2p/crypto/base58check"
)

// GenerateKeys will generate the peer initial keys and ID.
func GenerateKeys() (privateKey *ecdsa.PrivateKey, privateKeyBytes []byte, publicKey *ecdsa.PublicKey, publicKeyBytes []byte, peerID string, err error) {
	privateKey, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	privateKeyBytes, err = x509.MarshalECPrivateKey(privateKey)
	publicKeyBytes, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)

	shaHash := sha256.New()
	shaHash.Write(publicKeyBytes)
	hash := shaHash.Sum(nil)

	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(hash)
	hash = ripemd160Hash.Sum(nil)

	peerID = base58.Encode("00", hash)

	return privateKey, privateKeyBytes, &privateKey.PublicKey, publicKeyBytes, peerID, err
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
