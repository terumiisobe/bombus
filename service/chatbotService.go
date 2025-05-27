package service

import (
	"bombus/domain"
	"bombus/errs"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
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
		err := ValidateText(currentUserInteractionState, input)
		if err != nil {
			return cs.interactionRepo.GenerateText(domain.Fail, err.Message)
		}
		return cs.interactionRepo.GetTextByType(domain.Success)
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

	formValues := convertToFormValues(text)
	formSizes := []int{3, 4}
	if !contains(formSizes, len(formValues)) {
		return errs.NewValidationError("Número incorreto de linhas.")
	}

	invalidValues := []string{}
	var err error

	for idx, val := range formValues {
		switch idx {
		case 0:
			err = domain.ValidateStatus(val)
		case 1:
			_, err = time.Parse("02/01/2006", val)
		case 2:
			err = domain.ValidateSpecies(val)
		case 3:
			_, err = strconv.Atoi(val)
		}

		if err != nil {
			invalidValues = append(invalidValues, val)
		}
	}

	if len(invalidValues) > 0 {
		return errs.NewValidationError(fmt.Sprintf("Dados inválidos (%s).", strings.Join(reverse(invalidValues), ", ")))
	}

	return nil
}

func convertToFormValues(rawText string) []string {
	formSeparator := "\n"
	return reverse(convertStringToSlice(rawText, formSeparator))
}

func convertStringToSlice(s string, separator string) []string {
	splittedStrings := strings.Split(s, separator)
	var slice []string
	for _, split := range splittedStrings {
		trimmed := strings.TrimSpace(split)
		if trimmed != "" {
			slice = append(slice, trimmed)
		}
	}
	return slice
}

func reverse(s []string) []string {
	rev := make([]string, len(s))
	for i, v := range s {
		rev[len(s)-1-i] = v
	}
	return rev
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
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
