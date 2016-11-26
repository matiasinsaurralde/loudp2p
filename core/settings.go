package loudp2p

import (
	"encoding/json"
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
