package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"

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
