package config

import "github.com/google/wire"

// Set is a Wire provider set that gives configuration for this app.
var Set = wire.NewSet(
	NewConfig,
)

// Config
type Config struct {
	DBUser     string
	DBPassword string
	DBAddr     string
	DBName     string
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		DBUser:     "",
		DBPassword: "",
		DBAddr:     "",
		DBName:     "",
	}
}
