package config

import (
	"os"
)

// Config represetns Application Config
type Config struct {
	// Discord Bot Token
	Token string
	// Discord Bot Client ID
	ClientID string
}

// GetConfig is a function to generate app config struct.
func GetConfig() *Config {
	return &Config{
		Token:    os.Getenv("DISCORD_BOT_TOKEN"),
		ClientID: os.Getenv("543429901999407123"),
	}
}
