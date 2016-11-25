package crypto

import(
  "crypto/elliptic"
  "crypto/rand"
  "crypto/ecdsa"
  "crypto/x509"
  "crypto/sha256"
  "golang.org/x/crypto/ripemd160"

  "github.com/matiasinsaurralde/hellobitcoin/base58check"
)

func GenerateKeys() ([]byte, []byte, string) {
  var private_key_bytes, public_key_bytes []byte
  private_key, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	private_key_bytes, _ = x509.MarshalECPrivateKey(private_key)
	public_key_bytes, _ = x509.MarshalPKIXPublicKey(&private_key.PublicKey)

  shaHash := sha256.New()
  shaHash.Write(public_key_bytes)
  hash := shaHash.Sum(nil)

  ripemd160Hash := ripemd160.New()
  ripemd160Hash.Write(hash)
  hash = ripemd160Hash.Sum(nil)

  b58 := base58check.Encode("00", hash)

  return private_key_bytes, public_key_bytes, b58
}
