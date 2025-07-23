# Testing Guidelines

## Environment Variables and Configuration

### ❌ Bad Practices (Avoid These)

1. **Loading .env files in tests**
   ```go
   // DON'T DO THIS
   err := godotenv.Load("../../.env")
   if err != nil {
       t.Fatalf("Error loading .env file: %v", err)
   }
   ```

2. **Hardcoding API keys in tests**
   ```go
   // DON'T DO THIS
   APIKey := "sk-1234567890abcdef..."
   ```

3. **Making tests dependent on external files**
   ```go
   // DON'T DO THIS
   config := config.NewConfig() // Depends on .env file
   ```

### ✅ Good Practices

#### 1. Environment Variable Approach (Recommended for Integration Tests)

```go
func TestOpenAIParser_Parse(t *testing.T) {
    // Get API key from environment variable
    APIKey := os.Getenv("OPENAI_API_KEY")
    if APIKey == "" {
        t.Skip("OPENAI_API_KEY environment variable not set, skipping tests")
    }
    
    // Run your tests...
}
```

**Benefits:**
- Tests are environment-agnostic
- Works in CI/CD pipelines
- No dependency on local files
- Tests can be skipped when credentials aren't available

#### 2. Mock Configuration Approach (Recommended for Unit Tests)

```go
func TestSomeFunction(t *testing.T) {
    // Use test configuration
    testConfig := config.NewTestConfig()
    
    // Mock the configuration
    testConfig.OpenAIAPIKey = "test-key"
    
    // Run your tests with mock config...
}
```

#### 3. Helper Function Approach (Clean and Reusable)

```go
// test_helper.go
func getTestAPIKey(t *testing.T) string {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        t.Skip("OPENAI_API_KEY environment variable not set, skipping tests")
    }
    return apiKey
}

// In your test
func TestOpenAIParser_Parse(t *testing.T) {
    APIKey := getTestAPIKey(t)
    // Run your tests...
}
```

## Running Tests

### Local Development
```bash
# Set environment variable for integration tests
export OPENAI_API_KEY="your-actual-key"
go test ./...

# Or run specific test
go test ./infrastructure/openai -v
```

### CI/CD Pipeline
```bash
# Set environment variable in CI
OPENAI_API_KEY=${{ secrets.OPENAI_API_KEY }} go test ./...
```

### Skip Integration Tests
```bash
# Run only unit tests (no external dependencies)
go test -short ./...
```

## Test Categories

### Unit Tests
- No external dependencies
- Use mock configurations
- Fast execution
- Can run without environment setup

### Integration Tests
- Require external services (OpenAI API)
- Use environment variables
- Slower execution
- Should be skippable

### Example Test Structure
```go
func TestOpenAIParser_Parse(t *testing.T) {
    // Integration test - requires API key
    if testing.Short() {
        t.Skip("skipping integration test in short mode")
    }
    
    APIKey := getTestAPIKey(t)
    // ... test implementation
}

func TestOpenAIParser_Unit(t *testing.T) {
    // Unit test - no external dependencies
    testConfig := config.NewTestConfig()
    // ... test implementation
}
```

## Best Practices Summary

1. **Use environment variables** for integration tests
2. **Use mock configurations** for unit tests
3. **Skip tests gracefully** when dependencies aren't available
4. **Separate unit and integration tests** clearly
5. **Don't commit real credentials** to version control
6. **Use helper functions** for common test setup
7. **Make tests environment-agnostic** when possible 