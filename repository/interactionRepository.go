package repository

import "bombus/domain/chatbot"

type InteractionRepository interface {
	GetTextByType(chatbot.InteractionType) string
	GenerateText(chatbot.InteractionType, string) string
}
