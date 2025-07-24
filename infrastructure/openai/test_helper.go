package openai

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// getTestAPIKey returns the API key for testing, skipping the test if not available
func getTestAPIKey(t *testing.T) string {
	// Load .env file for tests
	if err := godotenv.Load("../../.env"); err != nil {
		// Don't fail the test if .env doesn't exist, just log it
		t.Logf("Could not load .env file: %v", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("OPENAI_API_KEY environment variable not set, skipping tests")
	}
	return apiKey
}
