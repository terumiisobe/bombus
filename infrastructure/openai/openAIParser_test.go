package openai

import (
	"bombus/domain/chatbot"
	"context"
	"reflect"
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

	options := []string{chatbot.ListColmeias.String(), chatbot.AddColmeiaForm.String()}

	t.Run("use deprecated model, should return external API request error", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, deprecatedModel)

		message := "some message"

		actionName, actionParams, err := parser.Parse(ctx, []string{}, message)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		if actionName != "" || actionParams != nil {
			t.Errorf("expected action name and params to be empty, got %+v and %+v", actionName, actionParams)
		}
	})

	t.Run("use incorrect api key, should return external API request error", func(t *testing.T) {
		parser := NewOpenAIParser(incorrectAPIKey, model)

		message := "some message"

		actionName, actionParams, err := parser.Parse(ctx, []string{}, message)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		if actionName != "" || actionParams != nil {
			t.Errorf("expected action name and params to be empty, got %+v and %+v", actionName, actionParams)
		}
	})

	t.Run("test list-like message with no filter, should return list_colmeia action", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, model)

		message := "listar"

		actionName, actionParams, err := parser.Parse(ctx, options, message)

		expActionName := chatbot.ListColmeias.String()
		expActionParams := map[string]string{}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if actionName != expActionName {
			t.Errorf("expected action name to be %s, got %s", expActionName, actionName)
		}
		if !reflect.DeepEqual(actionParams, expActionParams) {
			t.Errorf("expected action params to be %v, got %v", expActionParams, actionParams)
		}
	})

	t.Run("test list-like message with status filter, should return list_colmeia action", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, model)

		message := "listar colmeias com status 'em desenvolvimento'"

		actionName, actionParams, err := parser.Parse(ctx, options, message)

		expActionName := chatbot.ListColmeias.String()
		expActionParams := map[string]string{"status": "em desenvolvimento"}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if actionName != expActionName {
			t.Errorf("expected action name to be %s, got %s", expActionName, actionName)
		}
		if actionParams["status"] != expActionParams["status"] {
			t.Errorf("expected action params to be %v, got %v", expActionParams, actionParams)
		}
	})

	t.Run("test list-like message with status and species filter, should return list_colmeia action", func(t *testing.T) {
		parser := NewOpenAIParser(APIKey, model)

		message := "listar colmeias com status 'com mel' e espécie 'melipona bicolor'"

		actionName, actionParams, err := parser.Parse(ctx, options, message)

		expActionName := chatbot.ListColmeias.String()
		expActionParams := map[string]string{"status": "com mel", "species": "Melipona Bicolor"}

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if actionName != expActionName {
			t.Errorf("expected action name to be %s, got %s", expActionName, actionName)
		}
		if actionParams["status"] != expActionParams["status"] || actionParams["species"] != expActionParams["species"] {
			t.Errorf("expected action params to be %v, got %v", expActionParams, actionParams)
		}
	})
}
