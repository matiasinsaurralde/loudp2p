package loudp2p

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

const (
	settingsFilename = "settings.json"
)

// Settings holds the key pair & peer ID.
type Settings struct {
	PubKey  []byte
	PrivKey []byte
	PeerID  string
}

// Persist will persist the settings to disk.
func (s *Settings) Persist() (err error) {
	log.Println("Writing settings to disk.")
	var data []byte
	data, err = json.Marshal(s)
	err = ioutil.WriteFile(settingsFilename, data, 0700)
	return err
}

// Validate will validate the settings fields.
func (s *Settings) Validate() (err error) {
	if s.PubKey == nil {
		err = errors.New("No public key is present")
	} else if s.PrivKey == nil {
		err = errors.New("No public key is present")
	} else if s.PeerID == "" {
		err = errors.New("No peer ID is present")
	}
	return err
}

// LoadSettings will load the settings from disk.
func LoadSettings() (settings *Settings) {
	var data []byte
	var err error
	data, err = ioutil.ReadFile(settingsFilename)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil
	}
	return settings
}
