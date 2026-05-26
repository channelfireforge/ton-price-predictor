package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	APIEndpoint string  `yaml:"api_endpoint"`
	Alpha       float64 `yaml:"alpha"`
	Beta        float64 `yaml:"beta"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file failed: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("yaml unmarshal failed: %w", err)
	}

	if cfg.APIEndpoint == "" {
		cfg.APIEndpoint = "https://api.example.com/ton/price"
	}
	if cfg.Alpha == 0 {
		cfg.Alpha = 0.95
	}
	if cfg.Beta == 0 {
		cfg.Beta = 0.05
	}

	return &cfg, nil
}