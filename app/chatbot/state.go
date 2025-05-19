package chatbot

type ChatbotState struct {
  userMenuStates map[string]string
  stateLock      sync.Mutex
}

func


func setState(user, currentState string) map[string]string {
	log.Printf("Setting state from: %s to: %s", userMenuStates[user], currentState)
	stateLock.Lock()

	userMenuStates[user] = currentState
	defer stateLock.Unlock()

	return userMenuStates
}

func clearState(user string)  map[string]string {
	log.Printf("Cleanin up state from: %s", userMenuStates[user])
	stateLock.Lock()
	defer stateLock.Unlock()
	delete(userMenuStates, user)

    return userMenuStates
}
