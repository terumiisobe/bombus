package openai

import (
	"bombus/domain"

	"github.com/openai/openai-go"
)

func GetAllTools() []openai.ChatCompletionToolParam {
	return []openai.ChatCompletionToolParam{
		GetColmeiaListToolParams(),
		GetColmeiaAddToolParams(),
	}
}

func GetColmeiaListToolParams() openai.ChatCompletionToolParam {
	return openai.ChatCompletionToolParam{
		Function: openai.FunctionDefinitionParam{
			Name:        "list_colmeia",
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
						"enum":        domain.GetAllSpecies(),
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
			Name:        "add_colmeia",
			Description: openai.String("Add a new bee hive"),
			Parameters: openai.FunctionParameters{
				"type": "object",
				"properties": map[string]interface{}{
					"colmeia_id": map[string]interface{}{
						"type":        "integer",
						"description": "ID of bee hive (optional)",
					},
					"species": map[string]interface{}{
						"type":        "string",
						"enum":        domain.GetAllSpecies(),
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
