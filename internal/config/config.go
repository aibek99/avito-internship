package config

import "os"

// Config is
type Config struct {
	ServerAddress string
	Host          string
	Port          string
	User          string
	Password      string
	DBName        string
}

// LoadConfig is
func LoadConfig() (*Config, error) {
	dbConfig := Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		Host:          os.Getenv("POSTGRES_HOST"),
		User:          os.Getenv("POSTGRES_USERNAME"),
		Password:      os.Getenv("POSTGRES_PASSWORD"),
		DBName:        os.Getenv("POSTGRES_DATABASE"),
		Port:          os.Getenv("POSTGRES_PORT"),
	}
	return &dbConfig, nil
}
