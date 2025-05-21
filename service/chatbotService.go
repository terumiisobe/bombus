package service

import (
	"log"
	"sync"
)

type ChatbotService interface {
	GenerateMessage(string, string) string
}

type DefaultChatbotService struct {
	userStatesMap map[string]string
	stateLock     *sync.Mutex
}

func NewChatbotService() DefaultChatbotService {
	return DefaultChatbotService{
		userStatesMap: make(map[string]string),
		stateLock:     new(sync.Mutex),
	}
}

func (cs DefaultChatbotService) GenerateMessage(from, body string) string {

	currentState := cs.userStatesMap[from]

	switch currentState {
	case "":
		cs.setStateWithLock(from, "main_menu")
		return "Olá, meliponicultor! O que você quer fazer?\n Digite o número correspondente:\n 1. Listar contagem de colmeias\n 2. Adicionar colmeia\n 3. Adicionar várias colmeias\n 4. (em desenvolvimento) Escrever formulário de visita técnica\n 5. (em desenvolvimento) Listar tarefas pendentes"
	case "main_menu":
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

func (cs *DefaultChatbotService) setStateWithLock(user, currentState string) {
	log.Printf("Setting state from: %s to: %s", cs.userStatesMap[user], currentState)
	cs.stateLock.Lock()
	cs.userStatesMap[user] = currentState
	defer cs.stateLock.Unlock()
}

func (cs *DefaultChatbotService) clearStateWithLock(user string) {
	log.Printf("Cleanin up state from: %s", cs.userStatesMap[user])
	cs.stateLock.Lock()
	defer cs.stateLock.Unlock()
	delete(cs.userStatesMap, user)
}
