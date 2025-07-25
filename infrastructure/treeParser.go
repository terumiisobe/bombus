package infrastructure

import (
	"bombus/domain/chatbot"
	"bombus/errs"
	"context"
)

type TreeParser struct {
}

func NewTreeParser() *TreeParser {
	return &TreeParser{}
}

func (p *TreeParser) Parse(ctx context.Context, message string) (*chatbot.Interaction, *errs.AppError) {
	return nil, nil
}
