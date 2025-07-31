package openai

import (
	"bombus/errs"
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
)

type OpenAIParser struct {
	client openai.Client
	model  shared.ChatModel
}

// NewOpenAIParser creates a new OpenAIParser instance
func NewOpenAIParser(apiKey string, model shared.ChatModel) *OpenAIParser {
	client := openai.NewClient(option.WithAPIKey(apiKey))

	return &OpenAIParser{
		client: client,
		model:  model,
	}
}

func (o *OpenAIParser) Parse(ctx context.Context, options []string, message string) (string, map[string]string, *errs.AppError) {

	params := createParams(o.model, options, message)
	resp, err := o.sendRequest(ctx, params)
	if err != nil {
		return "", nil, err
	}

	toolCall, err := getToolCallFromResponse(resp)
	if err != nil {
		return "", nil, err
	}

	actionName, actionParams, err := convertToolCallToAction(toolCall)
	if err != nil {
		return "", nil, err
	}

	return actionName, actionParams, nil
}

func (o *OpenAIParser) sendRequest(ctx context.Context, params openai.ChatCompletionNewParams) (*openai.ChatCompletion, *errs.AppError) {
	resp, err := o.client.Chat.Completions.New(ctx, params)
	if err != nil {
		return nil, errs.NewExternalAPIRequestError(fmt.Sprintf("[OpenAI] %s", err.Error()))
	}
	return resp, nil
}

func convertToolCallToAction(toolCall *openai.ChatCompletionMessageToolCall) (string, map[string]string, *errs.AppError) {
	params := make(map[string]string)
	if toolCall.Function.Arguments != "" {
		// Parse the arguments JSON string into a map
		argsMap := make(map[string]interface{})
		if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &argsMap); err != nil {
			return "", nil, errs.NewParsingError(fmt.Sprintf("failed to parse function arguments: %s", err.Error()))
		}

		// Convert interface{} values to strings
		for k, v := range argsMap {
			if str, ok := v.(string); ok {
				params[k] = str
			} else {
				// Convert other types to string representation
				params[k] = fmt.Sprintf("%v", v)
			}
		}
	}

	return toolCall.Function.Name, params, nil
}

func getToolCallFromResponse(response *openai.ChatCompletion) (*openai.ChatCompletionMessageToolCall, *errs.AppError) {
	toolCalls := response.Choices[0].Message.ToolCalls
	if len(toolCalls) == 0 {
		return nil, errs.NewParsingError("AI could not find a tool call in the response")
	}
	return &toolCalls[0], nil
}

func createParams(model shared.ChatModel, options []string, message string) openai.ChatCompletionNewParams {

	messages := []openai.ChatCompletionMessageParamUnion{
		//openai.DeveloperMessage("You are a bee hive management system assistant, always answer in brazilian portuguese and be friendly, but keep answers short."),
		openai.UserMessage(message),
	}

	tools := GetTools(options)

	return openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    model,
		Tools:    tools,
	}
}
