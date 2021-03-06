package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Account holds the settings for one account
type Account struct {
	/*
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
	*/
	Instance     string
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
}

// Config holds chirp's config settings
type Config struct {
	Account []Account
	Style   string
}

const (
	configFile = "chirp.conf"
)

// LoadConfig returns the current config as a Config struct
func LoadConfig() Config {
	_, err := os.Stat(configFile)
	if err != nil {
		SaveConfig(Config{
			Style: "Material",
			Account: []Account{
				{
					Instance:     "your instance",
					Username:     "your username",
					Password:     "your password",
					ClientID:     "your client id",
					ClientSecret: "your client secret",
				},
			},
		})
		log.Fatal("Config file is missing, but a template was created for you! Please edit ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("Could not decode config file: ", err)
	}

	return config
}

// SaveConfig stores the current config
func SaveConfig(config Config) {
	f, err := os.Create(configFile)
	if err != nil {
		log.Fatal("Could not open config file: ", err)
	}
	if err := toml.NewEncoder(f).Encode(config); err != nil {
		log.Fatal("Could not encode config: ", err)
	}
}
