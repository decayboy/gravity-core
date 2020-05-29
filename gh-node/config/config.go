package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	NodeUrl        string
	GHNodeURL      string
	Timeout        int
	NebulaId       string
	ChainType      string
	NebulaContract string
}

func Load(filename string) (Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	if err := json.Unmarshal(file, &config); err != nil {
		return Config{}, err
	}
	return config, err
}
