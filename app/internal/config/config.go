package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	Memory struct {
		SQLitePath string `yaml:"sqlite_path"`
	} `yaml:"memory"`

	Tools struct {
		TailLogDefault string `yaml:"tail_log_default"`
	} `yaml:"tools"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	if c.Server.Addr == "" {
		return nil, errors.New("server.addr is required")
	}
	if c.Memory.SQLitePath == "" {
		return nil, errors.New("memory.sqlite_path is required")
	}
	return &c, nil
}
