package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

const (
	configPath = "../../config-local.yaml"
)

type Config struct {
	GRPC Server `validate:"required" yaml:"grpc"`
}

type Server struct {
	Host string `validate:"required" yaml:"host"`
	Port string `validate:"required" yaml:"port"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
