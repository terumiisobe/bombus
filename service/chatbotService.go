package service

import (
	"bombus/domain"
	"bombus/errs"
	"fmt"
	"log"
	"strings"
	"sync"
)

type ChatbotService interface {
	GenerateMessage(string, string) string
}

type ChatbotServiceImpl struct {
	userInteractionStateMap map[string]domain.InteractionType
	stateLock               *sync.Mutex
	interactionRepo         domain.InteractionRepository
}

func NewChatbotService(interactionRepo domain.InteractionRepository) ChatbotServiceImpl {
	return ChatbotServiceImpl{
		userInteractionStateMap: make(map[string]domain.InteractionType),
		stateLock:               new(sync.Mutex),
		interactionRepo:         interactionRepo,
	}
}

func NewChatbotServiceCustomMap(r domain.InteractionRepository, m map[string]domain.InteractionType) ChatbotServiceImpl {
	return ChatbotServiceImpl{
		userInteractionStateMap: m,
		stateLock:               new(sync.Mutex),
		interactionRepo:         r,
	}
}

func (cs ChatbotServiceImpl) GenerateOutputMessageTDD(user, input string) string {
	currentUserInteractionState := cs.userInteractionStateMap[user]
	if currentUserInteractionState == domain.AddColmeiaForm {
		ValidateText(currentUserInteractionState, input)
	}
	if currentUserInteractionState == domain.ListColmeias {
		return cs.interactionRepo.GetTextByType(domain.MainMenu)
	}
	if input == "1" {
		return cs.interactionRepo.GetTextByType(domain.ListColmeias)
	}
	if input == "2" {
		return cs.interactionRepo.GetTextByType(domain.AddColmeiaForm)
	}
	if input == "3" {
		return cs.interactionRepo.GetTextByType(domain.AddBatchColmeiaForm)
	}
	return cs.interactionRepo.GetTextByType(domain.MainMenu)
}

func ValidateText(interactionType domain.InteractionType, text string) *errs.AppError {
	if text == "" {
		return errs.NewValidationError("Texto vazio.")
	}
	formValues := convertRawTextIntoSlice(text)
	if len(formValues) != 3 && len(formValues) != 4 {
		return errs.NewValidationError("Número incorreto de linhas.")
	}
	if len(formValues) == 3 {
		return errs.NewValidationError(fmt.Sprintf("Dados inválidos (%s, %s, %s).", formValues[0], formValues[1], formValues[2]))
	}
	if len(formValues) == 4 {
		return errs.NewValidationError(fmt.Sprintf("Dados inválidos (%s, %s, %s, %s).", formValues[0], formValues[1], formValues[2], formValues[3]))
	}
	return nil
}

func convertRawTextIntoSlice(rawText string) []string {
	lines := strings.Split(rawText, "\n")
	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func (cs ChatbotServiceImpl) GenerateMessage(from, body string) string {

	currentState := cs.userInteractionStateMap[from]

	switch currentState {
	case 0:
		cs.setStateWithLock(from, domain.MainMenu)
		return "Olá, meliponicultor! O que você quer fazer?\n Digite o número correspondente:\n 1. Listar contagem de colmeias\n 2. Adicionar colmeia\n 3. Adicionar várias colmeias\n 4. (em desenvolvimento) Escrever formulário de visita técnica\n 5. (em desenvolvimento) Listar tarefas pendentes"
	case domain.MainMenu:
		switch body {
		case "1":
			// TODO: send list (getColmeias)
			cs.clearStateWithLock(from)
			return "Aqui está sua lista de colmeias!"
		case "2":
			// TODO: create colmeia (createColmeia)
			cs.clearStateWithLock(from)
			return "Vamos criar um registro de colmeia!"
		case "3":
			// TODO: create colmeia in batch (?)
			cs.clearStateWithLock(from)
			return "Vamos criar várias registro de colmeia!"
		}
	}

	return "Não entendi"
}

func (cs *ChatbotServiceImpl) setStateWithLock(user string, currentState domain.InteractionType) {
	log.Printf("Setting state from: %s to: %s", cs.userInteractionStateMap[user], currentState)
	cs.stateLock.Lock()
	cs.userInteractionStateMap[user] = currentState
	defer cs.stateLock.Unlock()
}

func (cs *ChatbotServiceImpl) clearStateWithLock(user string) {
	log.Printf("Cleanin up state from: %s", cs.userInteractionStateMap[user])
	cs.stateLock.Lock()
	defer cs.stateLock.Unlock()
	delete(cs.userInteractionStateMap, user)
}
