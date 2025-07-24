package service

import (
	"bombus/domain"
	"sync"
)

type ChatbotServiceImplAI struct {
	userInteractionStateMap map[string]domain.InteractionType
	stateLock               *sync.Mutex
	interactionRepo         domain.InteractionRepository

	colmeiaService ColmeiaService
}
