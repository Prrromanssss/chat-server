package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

// Config holds the configuration for the application, including server and database settings.
type Config struct {
	GRPC     Server   `validate:"required" yaml:"grpc"`
	Postgres Database `validate:"required" yaml:"postgres"`
}

// Server holds the configuration for the gRPC server.
type Server struct {
	Port string `validate:"required" yaml:"port"`
	Host string `validate:"required" yaml:"host"`
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

// Database holds the configuration for the PostgreSQL database.
type Database struct {
	Host     string `validate:"required" yaml:"host"`
	Port     string `validate:"required" yaml:"port"`
	User     string `validate:"required" yaml:"user"`
	Password string `validate:"required" yaml:"password"`
	DBName   string `validate:"required" yaml:"dbname"`
	SSLMode  string `validate:"required" yaml:"sslmode"`
}

func (d *Database) DSN() string {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.DBName,
		d.SSLMode,
	)

	return connStr
}

// LoadConfig reads and parses the configuration from a file specified by the CONFIG_PATH environment variable.
func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Check if the configuration file exists.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	// Read and parse the configuration file into the Config struct.
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
