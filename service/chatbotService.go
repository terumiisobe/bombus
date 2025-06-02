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

type ChatbotService interface{}

type ChatbotServiceImpl struct {
	userInteractionStateMap map[string]domain.InteractionType
	stateLock               *sync.Mutex
	interactionRepo         domain.InteractionRepository

	colmeiaService ColmeiaService
}

func NewChatbotService(interactionRepo domain.InteractionRepository, cs ColmeiaService) ChatbotServiceImpl {
	return ChatbotServiceImpl{
		userInteractionStateMap: make(map[string]domain.InteractionType),
		stateLock:               new(sync.Mutex),
		interactionRepo:         interactionRepo,
		colmeiaService:          cs,
	}
}

func NewChatbotServiceCustomMap(r domain.InteractionRepository, m map[string]domain.InteractionType, cs ColmeiaService) ChatbotServiceImpl {
	return ChatbotServiceImpl{
		userInteractionStateMap: m,
		stateLock:               new(sync.Mutex),
		interactionRepo:         r,
		colmeiaService:          cs,
	}
}

func (cs ChatbotServiceImpl) processInteractionAndGenerateResponse(user, input string) string {
	userCurrentInteractionState := cs.userInteractionStateMap[user]

	userNextInteractionState, err := GetNextInteraction(userCurrentInteractionState, input)
	if err != nil {
		// validation error
	}
	result := cs.executeAction(userNextInteractionState, input)
	response := GenerateMessage(userNextInteractionState, result)

	cs.updateUserInteractionState(user, userNextInteractionState)
	return response
}

func GetNextInteraction(state domain.InteractionType, input string) (domain.InteractionType, *errs.AppError) {
	err := ValidateInput(state, input)

	if state == domain.MainMenu && input == "1" {
		return domain.ListColmeias, nil
	}
	if state == domain.MainMenu && input == "2" {
		return domain.AddColmeiaForm, nil
	}
	if state == domain.MainMenu && input == "3" {
		return domain.AddBatchColmeiaForm, nil
	}
	if state == domain.ListColmeias {
		return domain.Init, nil
	}
	if state == domain.AddColmeiaForm || state == domain.AddBatchColmeiaForm {
		if err != nil {
			return domain.Fail, err
		}
		return domain.Success, nil
	}

	return domain.MainMenu, err
}

// impure function
func (cs ChatbotServiceImpl) executeAction(state domain.InteractionType, input string) string {
	// if state == domain.ListColmeias {
	// 	cs.colmeiaService.GetAllColmeia()
	// }
}

func GenerateMessage(state domain.InteractionType, input string) string {
	// state := cs.userInteractionStateMap[user]
	// if state == domain.AddColmeiaForm {
	// 	err := ValidateInput(state, input)
	// 	if err != nil {
	// 		return cs.interactionRepo.GenerateText(domain.Fail, err.Message)
	// 	}
	// 	return cs.interactionRepo.GetTextByType(domain.Success)
	// }
	// if state == domain.AddBatchColmeiaForm {
	// 	err := ValidateInput(state, input)
	// 	if err != nil {
	// 		return cs.interactionRepo.GenerateText(domain.Fail, err.Message)
	// 	}
	// 	return cs.interactionRepo.GetTextByType(domain.Success)
	// }
	// if state == domain.ListColmeias {
	// 	return cs.interactionRepo.GetTextByType(domain.MainMenu)
	// }
	// if input == "1" {
	// 	return cs.interactionRepo.GetTextByType(domain.ListColmeias)
	// }
	// if input == "2" {
	// 	return cs.interactionRepo.GetTextByType(domain.AddColmeiaForm)
	// }
	// if input == "3" {
	// 	return cs.interactionRepo.GetTextByType(domain.AddBatchColmeiaForm)
	// }
	// return cs.interactionRepo.GetTextByType(domain.MainMenu)
	return ""
}

func ValidateInput(state domain.InteractionType, input string) *errs.AppError {
	valid := []string{"1", "2", "3"}
	if state == domain.MainMenu && !containsString(valid, input) {
		return errs.NewValidationError("Opção inválida.")
	}
	if state == domain.AddColmeiaForm || state == domain.AddBatchColmeiaForm {
		formValues := convertToFormValues(input)
		return ValidateForm(state, formValues)
	}
	return nil
}

func ValidateForm(formType domain.InteractionType, formValues []string) *errs.AppError {
	addColmeiaFormSizes := []int{3, 4}
	addBatchColmeiaFormSizes := []int{4, 5}
	if formType == domain.AddColmeiaForm && !containsInt(addColmeiaFormSizes, len(formValues)) {
		return errs.NewValidationError("Número incorreto de linhas.")
	}

	if formType == domain.AddBatchColmeiaForm && !containsInt(addBatchColmeiaFormSizes, len(formValues)) {
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

func getValidationsPerFormValue(interactiontype domain.InteractionType, formSize int) map[int]func(string) bool {

	validationPerFormValue := make(map[int]func(string) bool)
	if formSize == 3 {
		validationPerFormValue[0] = isValidSpecies
		validationPerFormValue[1] = isValidStartingDate
		validationPerFormValue[2] = isValidStatus
	}
	if formSize == 4 && interactiontype == domain.AddColmeiaForm {
		validationPerFormValue[0] = isValidQRCode
		validationPerFormValue[1] = isValidSpecies
		validationPerFormValue[2] = isValidStartingDate
		validationPerFormValue[3] = isValidStatus
	}
	if formSize == 4 && interactiontype == domain.AddBatchColmeiaForm {
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

func (cs *ChatbotServiceImpl) updateUserInteractionState(user string, state domain.InteractionType) {

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
