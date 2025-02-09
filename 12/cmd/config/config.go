package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port int `json:"port"`
}

func LoadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
	 return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()
   
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
	 	return Config{}, fmt.Errorf("error decoding config file: %w", err)
	}
	return config, nil
}