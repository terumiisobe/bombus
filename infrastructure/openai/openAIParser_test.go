package openai

import (
	"bombus/domain/chatbot"
	"context"
	"testing"

	"github.com/openai/openai-go/shared"
)

func TestOpenAIParser_Parse(t *testing.T) {

	model := shared.ChatModelGPT4oMini
	deprecatedModel := shared.ChatModelGPT3_5Turbo16k

	// Get API key using helper function
	APIKey := getTestAPIKey(t)

	incorrectAPIKey := "incorrect-key"

	ctx := context.Background()

	t.Run("use deprecated model, should return external API request error", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, deprecatedModel)

		message := "some message"

		action, err := parser.Parse(ctx, message)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		if action != nil {
			t.Errorf("expected action to be nil, got %+v", action)
		}
	})

	t.Run("use incorrect api key, should return external API request error", func(t *testing.T) {
		parser := NewOpenAIParser(incorrectAPIKey, model)

		message := "some message"

		action, err := parser.Parse(ctx, message)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		if action != nil {
			t.Errorf("expected action to be nil, got %+v", action)
		}
	})

	t.Run("test list-like message with no filter, should return list_colmeia action", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, model)

		message := "listar"

		action, err := parser.Parse(ctx, message)

		expAction := chatbot.Action{
			Name:   chatbot.ListColmeia,
			Params: map[string]string{},
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if action.Name != expAction.Name {
			t.Errorf("expected action name to be %s, got %s", expAction.Name, action.Name)
		}
	})

	t.Run("test list-like message with status filter, should return list_colmeia action", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, model)

		message := "listar colmeias com status 'em desenvolvimento'"

		action, err := parser.Parse(ctx, message)

		expAction := chatbot.Action{
			Name:   chatbot.ListColmeia,
			Params: map[string]string{"status": "em desenvolvimento"},
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if action.Name != expAction.Name {
			t.Errorf("expected action name to be %s, got %s", expAction.Name, action.Name)
		}
		if action.Params["status"] != expAction.Params["status"] {
			t.Errorf("expected action params to be %v, got %v", expAction.Params, action.Params)
		}
	})

	t.Run("test list-like message with status and species filter, should return list_colmeia action", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, model)

		message := "listar colmeias com status 'com mel' e espécie 'melipona bicolor'"

		action, err := parser.Parse(ctx, message)

		expAction := chatbot.Action{
			Name:   chatbot.ListColmeia,
			Params: map[string]string{"status": "com mel", "species": "Melipona Bicolor"},
		}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if action.Name != expAction.Name {
			t.Errorf("expected action name to be %s, got %s", expAction.Name, action.Name)
		}
		if action.Params["status"] != expAction.Params["status"] || action.Params["species"] != expAction.Params["species"] {
			t.Errorf("expected action params to be %v, got %v", expAction.Params, action.Params)
		}
	})
}
