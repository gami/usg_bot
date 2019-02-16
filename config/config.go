package config

import (
	"fmt"
	"os"
	"strings"
)

// Config represetns Application Config
type Config struct {
	// Discord Bot Token
	Token string
	// Discord Bot Client ID
	ClientID string
	// Tournamet Spread Sheet ID
	TournamentSheetIDs map[string]string
	// Google Service Asccount JWT json file Path
	GoogleJWTPath string
}

// GetConfig is a function to generate app config struct.
func GetConfig() *Config {

	return &Config{
		Token:              os.Getenv("DISCORD_BOT_TOKEN"),
		ClientID:           os.Getenv("DISCORD_BOT_CLIENT_ID"),
		TournamentSheetIDs: makeTournamentSheetIDs(),
		GoogleJWTPath:      os.Getenv("GOOGLE_SERVICE_ACCOUNT_JWT"),
	}
}

func makeTournamentSheetIDs() map[string]string {
	keyStr := os.Getenv("TOURNAMENT_SHEET_KEYS")
	keys := strings.Split(keyStr, ",")

	sheetIds := make(map[string]string)

	for _, key := range keys {
		sheetID := os.Getenv(fmt.Sprintf("TOURNAMENT_SHEET_ID_%v", key))
		sheetIds[key] = sheetID
	}
	return sheetIds
}
