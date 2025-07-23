package config

import "os"

// ConfigProvider defines the interface for configuration
type ConfigProvider interface {
	GetOpenAIAPIKey() string
	GetTwilioAccountSID() string
	GetTwilioAuthToken() string
	GetTwilioSenderNumber() string
	GetDBUser() string
	GetDBPass() string
}

type Config struct {
	AccountSID   string
	AuthToken    string
	FromNumber   string
	OpenAIAPIKey string
	DBUser       string
	DBPass       string
}

// NewConfig creates a new config from environment variables
func NewConfig() *Config {
	return &Config{
		AccountSID:   os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:    os.Getenv("TWILIO_AUTHTOKEN"),
		FromNumber:   os.Getenv("TWILIO_SENDER_NUMBER"),
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASS"),
	}
}

// GetOpenAIAPIKey returns the OpenAI API key
func (c *Config) GetOpenAIAPIKey() string {
	return c.OpenAIAPIKey
}

// GetTwilioAccountSID returns the Twilio Account SID
func (c *Config) GetTwilioAccountSID() string {
	return c.AccountSID
}

// GetTwilioAuthToken returns the Twilio Auth Token
func (c *Config) GetTwilioAuthToken() string {
	return c.AuthToken
}

// GetTwilioSenderNumber returns the Twilio sender number
func (c *Config) GetTwilioSenderNumber() string {
	return c.FromNumber
}

// GetDBUser returns the database user
func (c *Config) GetDBUser() string {
	return c.DBUser
}

// GetDBPass returns the database password
func (c *Config) GetDBPass() string {
	return c.DBPass
}
