package service

import (
	"bombus/domain"
	"bombus/domain/chatbot"
	"bombus/errs"
	"bombus/repository"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ChatbotServiceImplDefault struct {
	userInteractionStateMap map[string]chatbot.InteractionType
	stateLock               *sync.Mutex
	interactionRepo         repository.InteractionRepository

	colmeiaService ColmeiaService
}

func NewChatbotService(interactionRepo repository.InteractionRepository, cs ColmeiaService) ChatbotServiceImplDefault {
	return ChatbotServiceImplDefault{
		userInteractionStateMap: make(map[string]chatbot.InteractionType),
		stateLock:               new(sync.Mutex),
		interactionRepo:         interactionRepo,
		colmeiaService:          cs,
	}
}

func NewChatbotServiceCustomMap(r repository.InteractionRepository, m map[string]chatbot.InteractionType, cs ColmeiaService) ChatbotServiceImplDefault {
	return ChatbotServiceImplDefault{
		userInteractionStateMap: m,
		stateLock:               new(sync.Mutex),
		interactionRepo:         r,
		colmeiaService:          cs,
	}
}

func (cs ChatbotServiceImplDefault) GenerateReplyMessage(user, input string) string {
	return ""
}

func (cs ChatbotServiceImplDefault) processInteractionAndGenerateResponse(user, input string) string {
	userCurrentInteractionState := cs.userInteractionStateMap[user]

	//TODO: integrate GetNextInteraction, executeAction, GenerateMessage into interaction struct
	userNextInteractionState, err := GetNextInteraction(userCurrentInteractionState, input)
	if err != nil {
		// validation error
	}
	result, err := cs.executeAction(userNextInteractionState, input)
	response := cs.GenerateMessage(userNextInteractionState, result)

	cs.updateUserInteractionState(user, userNextInteractionState)
	return response
}

func GetNextInteraction(state chatbot.InteractionType, input string) (chatbot.InteractionType, *errs.AppError) {
	err := ValidateInput(state, input)

	if state == chatbot.MainMenu && input == "1" {
		return chatbot.ListColmeias, nil
	}
	if state == chatbot.MainMenu && input == "2" {
		return chatbot.AddColmeiaForm, nil
	}
	if state == chatbot.MainMenu && input == "3" {
		return chatbot.AddBatchColmeiaForm, nil
	}
	if state == chatbot.ListColmeias {
		return chatbot.Init, nil
	}
	if state == chatbot.AddColmeiaForm || state == chatbot.AddBatchColmeiaForm {
		if err != nil {
			return chatbot.AddFail, err
		}
		return chatbot.AddSuccess, nil
	}

	return chatbot.MainMenu, err
}

// impure function
func (cs ChatbotServiceImplDefault) executeAction(state chatbot.InteractionType, input string) (string, *errs.AppError) {
	if state == chatbot.ListColmeias {
		countBySpecies, err := cs.colmeiaService.CountBySpecies()
		if err != nil {
			return "", err
		}
		return convertMapToString(countBySpecies), nil
	}
	if state == chatbot.AddColmeiaForm {
		// TODO: call create colmeia
		return "", nil
	}
	if state == chatbot.AddBatchColmeiaForm {
		// TODO: call create batch colmeia
		return "", nil
	}
	return "", nil
}

// TODO: move to domain.colmeia
func convertToColmeia(s string) domain.Colmeia {
	return domain.Colmeia{}
}

// TODO: move to interaction ListColmeia
func convertMapToString(m map[string]int) string {
	lines := []string{}
	for idx, val := range m {
		lines = append(lines, fmt.Sprintf("%s: %d", idx, val))
	}
	return strings.Join(lines, "\n")
}

func (cs ChatbotServiceImplDefault) GenerateMessage(state chatbot.InteractionType, input string) string {
	if state == chatbot.AddColmeiaForm {
		err := ValidateInput(state, input)
		if err != nil {
			return cs.interactionRepo.GenerateText(chatbot.AddFail, err.Message)
		}
		return cs.interactionRepo.GetTextByType(chatbot.AddSuccess)
	}
	if state == chatbot.AddBatchColmeiaForm {
		err := ValidateInput(state, input)
		if err != nil {
			return cs.interactionRepo.GenerateText(chatbot.AddFail, err.Message)
		}
		return cs.interactionRepo.GetTextByType(chatbot.AddSuccess)
	}
	if state == chatbot.ListColmeias {
		return cs.interactionRepo.GetTextByType(chatbot.MainMenu)
	}
	if input == "1" {
		return cs.interactionRepo.GetTextByType(chatbot.ListColmeias)
	}
	if input == "2" {
		return cs.interactionRepo.GetTextByType(chatbot.AddColmeiaForm)
	}
	if input == "3" {
		return cs.interactionRepo.GetTextByType(chatbot.AddBatchColmeiaForm)
	}
	return cs.interactionRepo.GetTextByType(chatbot.MainMenu)
}

func ValidateInput(state chatbot.InteractionType, input string) *errs.AppError {
	valid := []string{"1", "2", "3"}
	if state == chatbot.MainMenu && !containsString(valid, input) {
		return errs.NewValidationError("Opção inválida.")
	}
	if state == chatbot.AddColmeiaForm || state == chatbot.AddBatchColmeiaForm {
		formValues := convertToFormValues(input)
		return ValidateForm(state, formValues)
	}
	return nil
}

func ValidateForm(formType chatbot.InteractionType, formValues []string) *errs.AppError {
	addColmeiaFormSizes := []int{3, 4}
	addBatchColmeiaFormSizes := []int{4, 5}
	if formType == chatbot.AddColmeiaForm && !containsInt(addColmeiaFormSizes, len(formValues)) {
		return errs.NewValidationError("Número incorreto de linhas.")
	}

	if formType == chatbot.AddBatchColmeiaForm && !containsInt(addBatchColmeiaFormSizes, len(formValues)) {
		return errs.NewValidationError("Número incorreto de linhas.")
	}

	validationPerFormValue := getValidationsPerFormValue(formType, len(formValues))

	invalidValues := []string{}
	for idx, val := range validationPerFormValue {
		isValid := val(formValues[idx])
		if !isValid {
			invalidValues = append(invalidValues, formValues[idx])
		}
	}
	if len(invalidValues) > 0 {
		return errs.NewValidationError(fmt.Sprintf("Dados inválidos (%s).", strings.Join(invalidValues, ", ")))
	}

	return nil
}

func getValidationsPerFormValue(interactiontype chatbot.InteractionType, formSize int) map[int]func(string) bool {

	validationPerFormValue := make(map[int]func(string) bool)
	if formSize == 3 {
		validationPerFormValue[0] = isValidSpecies
		validationPerFormValue[1] = isValidStartingDate
		validationPerFormValue[2] = isValidStatus
	}
	if formSize == 4 && interactiontype == chatbot.AddColmeiaForm {
		validationPerFormValue[0] = isValidQRCode
		validationPerFormValue[1] = isValidSpecies
		validationPerFormValue[2] = isValidStartingDate
		validationPerFormValue[3] = isValidStatus
	}
	if formSize == 4 && interactiontype == chatbot.AddBatchColmeiaForm {
		validationPerFormValue[0] = isValidQuantity
		validationPerFormValue[1] = isValidSpecies
		validationPerFormValue[2] = isValidStartingDate
		validationPerFormValue[3] = isValidStatus
	}
	if formSize == 5 {
		validationPerFormValue[0] = isValidQuantity
		validationPerFormValue[1] = isValidQRCode
		validationPerFormValue[2] = isValidSpecies
		validationPerFormValue[3] = isValidStartingDate
		validationPerFormValue[4] = isValidStatus

	}
	return validationPerFormValue
}

func isValidQuantity(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

func isValidQRCode(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

func isValidSpecies(v string) bool {
	err := domain.ValidateSpecies(v)
	return err == nil
}

func isValidStartingDate(v string) bool {
	_, err := time.Parse("02/01/2006", v)
	return err == nil
}
func isValidStatus(v string) bool {
	err := domain.ValidateStatus(v)
	return err == nil
}

func convertToFormValues(rawText string) []string {
	formSeparator := "\n"
	return convertStringToSlice(rawText, formSeparator)
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

func containsInt(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func containsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func (cs *ChatbotServiceImplDefault) updateUserInteractionState(user string, state chatbot.InteractionType) {

}

func (cs *ChatbotServiceImplDefault) setStateWithLock(user string, currentState chatbot.InteractionType) {
	log.Printf("Setting state from: %s to: %s", cs.userInteractionStateMap[user], currentState)
	cs.stateLock.Lock()
	cs.userInteractionStateMap[user] = currentState
	defer cs.stateLock.Unlock()
}

func (cs *ChatbotServiceImplDefault) clearStateWithLock(user string) {
	log.Printf("Cleanin up state from: %s", cs.userInteractionStateMap[user])
	cs.stateLock.Lock()
	defer cs.stateLock.Unlock()
	delete(cs.userInteractionStateMap, user)
}
