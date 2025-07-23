package config

// TestConfig provides a mock configuration for testing
type TestConfig struct {
	OpenAIAPIKey string
	AccountSID   string
	AuthToken    string
	FromNumber   string
	DBUser       string
	DBPass       string
}

// NewTestConfig creates a test configuration with default values
func NewTestConfig() *TestConfig {
	return &TestConfig{
		OpenAIAPIKey: "test-openai-key",
		AccountSID:   "test-account-sid",
		AuthToken:    "test-auth-token",
		FromNumber:   "test-from-number",
		DBUser:       "test-db-user",
		DBPass:       "test-db-pass",
	}
}

// GetOpenAIAPIKey returns the test OpenAI API key
func (tc *TestConfig) GetOpenAIAPIKey() string {
	return tc.OpenAIAPIKey
}

// GetTwilioAccountSID returns the test Twilio Account SID
func (tc *TestConfig) GetTwilioAccountSID() string {
	return tc.AccountSID
}

// GetTwilioAuthToken returns the test Twilio Auth Token
func (tc *TestConfig) GetTwilioAuthToken() string {
	return tc.AuthToken
}

// GetTwilioSenderNumber returns the test Twilio sender number
func (tc *TestConfig) GetTwilioSenderNumber() string {
	return tc.FromNumber
}

// GetDBUser returns the test database user
func (tc *TestConfig) GetDBUser() string {
	return tc.DBUser
}

// GetDBPass returns the test database password
func (tc *TestConfig) GetDBPass() string {
	return tc.DBPass
}
