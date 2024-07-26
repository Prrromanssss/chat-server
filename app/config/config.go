package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

// Config holds the configuration for the application, including server and database settings.
type Config struct {
	GRPC     Server   `validate:"required" yaml:"grpc"`     // gRPC server configuration
	Postgres Database `validate:"required" yaml:"postgres"` // PostgreSQL database configuration
}

// Server holds the configuration for the gRPC server.
type Server struct {
	Host string `validate:"required" yaml:"host"` // gRPC server host
	Port string `validate:"required" yaml:"port"` // gRPC server port
}

// Database holds the configuration for the PostgreSQL database.
type Database struct {
	Host     string `validate:"required" yaml:"host"`     // Database host
	Port     string `validate:"required" yaml:"port"`     // Database port
	User     string `validate:"required" yaml:"user"`     // Database user
	Password string `validate:"required" yaml:"password"` // Database password
	DBName   string `validate:"required" yaml:"dbname"`   // Database name
	SSLMode  string `validate:"required" yaml:"sslmode"`  // SSL mode for database connection
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
