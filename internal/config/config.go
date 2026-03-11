package config

import (
	"encoding/json"
	"os"
)

type LetsEncryptConfig struct {
	Enabled bool   `json:"enabled"`
	Domain  string `json:"domain"`
	Email   string `json:"email"`
}

type Config struct {
	Port        int               `json:"port"`
	LetsEncrypt LetsEncryptConfig `json:"letsencrypt"`
}

func Load() (*Config, error) {
	cfg := &Config{
		Port: 80,
	}

	data, err := os.ReadFile("semdgo.json")
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	if cfg.Port == 0 {
		cfg.Port = 80
	}

	return cfg, nil
}
