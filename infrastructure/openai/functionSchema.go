package openai

import (
	"bombus/domain"
	"bombus/domain/chatbot"
	"fmt"
	"slices"

	"github.com/openai/openai-go"
)

var (
	colmeiaListToolName          = chatbot.ListColmeias.String()
	colmeiaAddToolName           = chatbot.AddColmeiaForm.String()
	colmeiaAddValidationToolName = chatbot.AddColmeiaValidation.String()
)

func GetTools(options []string) []openai.ChatCompletionToolParam {
	tools := []openai.ChatCompletionToolParam{}
	if slices.Contains(options, colmeiaListToolName) {
		tools = append(tools, GetColmeiaListToolParams())
	}
	if slices.Contains(options, colmeiaAddToolName) {
		tools = append(tools, GetColmeiaAddToolParams())
	}
	if slices.Contains(options, colmeiaAddValidationToolName) {
		tools = append(tools, GetColmeiaAddValidationToolParams())
	}
	fmt.Println(tools)
	return tools
}

func GetColmeiaListToolParams() openai.ChatCompletionToolParam {
	return openai.ChatCompletionToolParam{
		Function: openai.FunctionDefinitionParam{
			Name:        colmeiaListToolName,
			Description: openai.String("List bee hives, in case of doubt, return all bee hives."),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"status": map[string]interface{}{
						"type":        "string",
						"enum":        domain.GetAllStatus(),
						"description": "Status of bee hive (optional)",
					},
					"species": map[string]interface{}{
						"type":        "string",
						"description": "Species of bee hive (optional)",
					},
				},
				"required": []string{},
			},
		},
	}
}

func GetColmeiaAddToolParams() openai.ChatCompletionToolParam {
	return openai.ChatCompletionToolParam{
		Function: openai.FunctionDefinitionParam{
			Name:        colmeiaAddToolName,
			Description: openai.String("Add a new bee hive"),
		},
	}
}

func GetColmeiaAddValidationToolParams() openai.ChatCompletionToolParam {
	return openai.ChatCompletionToolParam{
		Function: openai.FunctionDefinitionParam{
			Name:        colmeiaAddValidationToolName,
			Description: openai.String("Validate the parameters of a new bee hive"),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"colmeia_id": map[string]interface{}{
						"type":        "integer",
						"description": "ID of bee hive (optional)",
					},
					"species": map[string]interface{}{
						"type":        "string",
						"description": "Species of bee hive (required)",
					},
					"starting_date": map[string]interface{}{
						"type":        "string",
						"description": "Starting date of bee hive (optional)",
					},
					"status": map[string]interface{}{
						"type":        "string",
						"enum":        domain.GetAllStatus(),
						"description": "Status of bee hive (optional)",
					},
				},
				"required": []string{"species"},
			},
		},
	}
}
