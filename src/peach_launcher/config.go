package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Launcher struct {
		MaxClients     int    `json:"max_clients"`
		LogLevel       string `json:"log_level"`
		CoordinatorURL string `json:"coordinator"`
	} `json:"launcher,omitempty"`
	Coordinator struct {
		Launch              bool   `json:"launch"`
		Port                string `json:"port"`
		DBCredentials       string `json:"dbc"`
		SpotifyClientID     string `json:"spotify_client_id"`
		SpotifyClientSecret string `json:"spotify_client_secret"`
		CertType            string `json:"cert_type"`
		Domain              string `json:"domain"`
	} `json:"coordinator,omitempty"`
	Secret          string `json:"secret"`
	RedactSensitive bool   `json:"redact_sensitive"`
}

func (l *Launcher) loadJson() error {
	f, err := os.Open("launchcfg.json")
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &l.Config)

	return nil
}
