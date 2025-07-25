package repository

import "bombus/domain"

type InteractionRepository interface {
	GetTextByType(domain.InteractionType) string
	GenerateText(domain.InteractionType, string) string
}
