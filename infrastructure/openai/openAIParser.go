package openai

import (
	"bombus/domain/chatbot"
	"bombus/errs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type OpenAIParser struct {
	APIKey string
	model  string
}

func (o *OpenAIParser) Parse(ctx context.Context, message string) (*chatbot.Action, *errs.AppError) {
	//functionSchemas := GetColmeiaFunctionSchemas()

	payload := map[string]interface{}{
		"model": o.model,
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
		//"functions":     functionSchemas,
		//"function_call": "auto",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errs.NewJsonConversionError(fmt.Sprintf("failed to marshal payload: %s", err.Error()))
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("failed to create request: %s", err.Error()))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("failed to call OpenAI API: %s", err.Error()))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("failed to read response body: %s", err.Error()))
	}

	log.Printf("OpenAI API Response Body: %s", string(body))

	// Check for API errors first
	var errorResp struct {
		Error struct {
			Message string `json:"message"`
			Type    string `json:"type"`
			Code    string `json:"code"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &errorResp); err == nil && errorResp.Error.Message != "" {
		return nil, errs.NewExternalAPIRequestError(fmt.Sprintf("[OpenAI] %s (type: %s, code: %s)",
			errorResp.Error.Message, errorResp.Error.Type, errorResp.Error.Code))
	}

	var apiResp struct {
		Choices []struct {
			Message struct {
				FunctionCall struct {
					Name      string `json:"name"`
					Arguments string `json:"arguments"`
				} `json:"function_call"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, errs.NewJsonConversionError(fmt.Sprintf("failed to unmarshal OpenAI response: %s", err.Error()))
	}

	log.Printf("OpenAI API Response: %+v", apiResp)

	if len(apiResp.Choices) == 0 || apiResp.Choices[0].Message.FunctionCall.Name == "" {
		return nil, errs.NewUnexpectedError("no function_call in OpenAI response")
	}

	name := apiResp.Choices[0].Message.FunctionCall.Name
	params := make(map[string]string)
	if apiResp.Choices[0].Message.FunctionCall.Arguments != "" {
		_ = json.Unmarshal([]byte(apiResp.Choices[0].Message.FunctionCall.Arguments), &params)
	}

	return &chatbot.Action{
		Name:   chatbot.ActionName(name),
		Params: params,
	}, nil
}
