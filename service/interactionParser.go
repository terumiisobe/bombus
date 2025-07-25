package service

import (
	"bombus/errs"
	"context"
)

type InteractionParser interface {
	Parse(ctx context.Context, options []string, message string) (string, map[string]string, *errs.AppError)
}
