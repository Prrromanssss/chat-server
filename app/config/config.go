package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	GRPC     Server   `validate:"required" yaml:"grpc"`
	Postgres Database `validate:"required" yaml:"postgres"`
}

type Server struct {
	Host string `validate:"required" yaml:"host"`
	Port string `validate:"required" yaml:"port"`
}

type Database struct {
	Host     string `validate:"required" yaml:"host"`
	Port     string `validate:"required" yaml:"port"`
	User     string `validate:"required" yaml:"user"`
	Password string `validate:"required" yaml:"password"`
	DBName   string `validate:"required" yaml:"dbname"`
	SSLMode  string `validate:"required" yaml:"sslmode"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
