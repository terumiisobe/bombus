package service

import (
	"bombus/domain"
	"bombus/repository"
	"sync"
)

type ChatbotServiceImplAI struct {
	userInteractionStateMap map[string]domain.InteractionType
	stateLock               *sync.Mutex
	interactionRepo         repository.InteractionRepository

	colmeiaService ColmeiaService
}

func (cs ChatbotServiceImplAI) GenerateReplyMessage(user, input string) string {
	return ""
}
