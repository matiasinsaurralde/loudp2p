package loudp2p

import(
  "log"
  "encoding/json"
  "io/ioutil"
)

const(
  SettingsFilename = "settings.json"
)

type Settings struct {
  PubKey []byte
  PrivKey []byte
  PeerId string
}

func(s *Settings) Persist() (err error) {
  log.Println("Writing settings to disk.")
  var data []byte
  data, err = json.Marshal(s)
  err = ioutil.WriteFile(SettingsFilename, data, 0700)
  return err
}

func LoadSettings() (settings *Settings) {
  var data []byte
  var err error
  data, err = ioutil.ReadFile(SettingsFilename)
  if err != nil {
    return nil
  }
  err = json.Unmarshal(data, &settings)
  if err != nil {
    return nil
  }
  return settings
}
