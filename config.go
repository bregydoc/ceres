package main

import (
	"encoding/json"
	"io/ioutil"
)

// GetConfigFromFile ...
func GetConfigFromFile(path string) (*ConfigFile, error) {
	config := new(ConfigFile)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
