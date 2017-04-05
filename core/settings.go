package loudp2p

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"crypto/ecdsa"

	crypto "github.com/matiasinsaurralde/loudp2p/crypto"
)

const (
	DefaultSettingsFilename = "settings.json"
	DefaultRPCPort          = 5555
)

// Settings holds the key pair & peer ID.
type Settings struct {
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	PublicKey  *ecdsa.PublicKey  `json:"-"`

	PrivKeyBytes []byte
	PubKeyBytes  []byte
	PeerID       string

	IgnoreInitialPeers bool `json:"-"`
	RPCPort            int64
}

// LoadSettings will load the settings from disk.
func LoadSettings() (settings *Settings) {
	var data []byte
	var err error
	data, err = ioutil.ReadFile(DefaultSettingsFilename)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		log.Println("Couldn't parse settings!")
		return nil
	}
	err = settings.LoadKeys()
	if err != nil {
		log.Println("Couldn't parse keys!")
		return nil
	}
	return settings
}

// Persist will persist the settings to disk.
func (s *Settings) Persist() (err error) {
	log.Println("Writing settings to disk.")
	var data []byte
	data, err = json.Marshal(s)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(DefaultSettingsFilename, data, 0700)
}

// Validate will validate the settings fields.
func (s *Settings) Validate() (err error) {
	if s.PublicKey == nil {
		err = errors.New("No public key is present")
	} else if s.PrivateKey == nil {
		err = errors.New("No private key is present")
	} else if s.PeerID == "" {
		err = errors.New("No peer ID is present")
	}
	return err
}

// LoadKeys will call crypto.ParseKeys.
func (s *Settings) LoadKeys() (err error) {
	s.PrivateKey, s.PublicKey, err = crypto.ParseKeys(s.PrivKeyBytes, s.PubKeyBytes)
	return err
}
