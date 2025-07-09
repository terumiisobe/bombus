package openai

import (
	"context"
	"testing"
)

func TestOpenAIParser_Parse(t *testing.T) {
	parser := &OpenAIParser{
		APIKey: "my-key",
		model:  "gpt-4o-mini"}
	ctx := context.Background()

	t.Run("use deprecated model, should return external API request error", func(t *testing.T) {
		deprecatedModel := "gpt-3.5-turbo-0613"
		parser.model = deprecatedModel

		message := "some message"

		action, err := parser.Parse(ctx, message)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		if action != nil {
			t.Errorf("expected action to be nil, got %+v", action)
		}
	})

	t.Run("should return success", func(t *testing.T) {
		// Reset to valid model
		parser.model = "gpt-4o-mini"
		message := "listar"

		action, err := parser.Parse(ctx, message)

		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if action == nil {
			t.Errorf("expected action to not be nil")
		}
	})
}
