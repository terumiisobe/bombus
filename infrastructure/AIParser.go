package infrastructure

import (
	"bombus/domain/chatbot"
	"bombus/errs"
	"context"
)

type AIParser interface {
	Parse(ctx context.Context, message string) (*chatbot.Action, *errs.AppError)
}
