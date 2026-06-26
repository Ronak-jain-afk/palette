package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Defaults Defaults `yaml:"defaults"`
}

type Defaults struct {
	Mood   string `yaml:"mood"`
	Count  int    `yaml:"count"`
	Scheme string `yaml:"scheme"`
}

func Path() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(configDir, "palette")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

func Load() (*Config, error) {
	p, err := Path()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Defaults: Defaults{
			Mood:   "dark",
			Count:  5,
			Scheme: "analogous",
		},
	}

	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	if cfg.Defaults.Mood == "" {
		cfg.Defaults.Mood = "dark"
	}
	if cfg.Defaults.Count == 0 {
		cfg.Defaults.Count = 5
	}
	if cfg.Defaults.Scheme == "" {
		cfg.Defaults.Scheme = "analogous"
	}

	return cfg, nil
}

func Save(cfg *Config) error {
	p, err := Path()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(p, data, 0644)
}
