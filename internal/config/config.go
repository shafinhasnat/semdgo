package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type LetsEncryptConfig struct {
	Enabled bool   `json:"enabled"`
	Domain  string `json:"domain"`
	Email   string `json:"email"`
}

type Config struct {
	Port             int               `json:"port"`
	ContentPath      string            `json:"content_path"`
	ContentEntrypoint string           `json:"content_entrypoint"`
	LetsEncrypt      LetsEncryptConfig `json:"letsencrypt"`
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:              80,
		ContentPath:       "/var/semdgo/content",
		ContentEntrypoint: "README.md",
	}

	data, err := os.ReadFile("semdgo.json")
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("invalid semdgo.json: %w", err)
	}

	if cfg.Port == 0 {
		cfg.Port = 80
	}
	if cfg.ContentPath == "" {
		cfg.ContentPath = "/var/semdgo/content"
	}
	if cfg.ContentEntrypoint == "" {
		cfg.ContentEntrypoint = "README.md"
	}

	return cfg, nil
}

func Validate(cfg *Config) error {
	info, err := os.Stat(cfg.ContentPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("content_path %q does not exist", cfg.ContentPath)
		}
		return fmt.Errorf("content_path %q: %w", cfg.ContentPath, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("content_path %q is not a directory", cfg.ContentPath)
	}

	entrypoint := filepath.Join(cfg.ContentPath, cfg.ContentEntrypoint)
	if _, err := os.Stat(entrypoint); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("content_entrypoint %q does not exist in %q", cfg.ContentEntrypoint, cfg.ContentPath)
		}
		return fmt.Errorf("content_entrypoint %q: %w", entrypoint, err)
	}

	if cfg.LetsEncrypt.Enabled && cfg.LetsEncrypt.Domain == "" {
		return fmt.Errorf("letsencrypt.domain must be set when letsencrypt.enabled is true")
	}

	return nil
}
