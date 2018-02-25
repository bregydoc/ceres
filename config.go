package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

// CreateConfigFile ...
func CreateConfigFile(path string, config *ConfigFile) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
