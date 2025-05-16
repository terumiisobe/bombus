package config

import "os"

type Config struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

func NewConfig() *Config {
	return &Config{
		AccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTHTOKEN"),
		FromNumber: os.Getenv("TWILIO_SENDER_NUMBER"),
	}
}
