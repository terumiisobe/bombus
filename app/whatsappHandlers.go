package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var (
	userMenuStates = make(map[string]string)
	stateLock      sync.Mutex
)

type WhatsappHandler struct {
}

func (wh *WhatsappHandler) webhookHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	from := r.FormValue("From")
	body := r.FormValue("Body")

	state := userMenuStates[from]

	log.Printf("From: %s | Message: %s \n", from, body)

	// first interaction: send menu with all the options
	switch state {
	case "":
		setState(from, "main_menu")
		sendReply(from, "Olá, meliponicultor! O que você quer fazer?\n Digite o número correspondente:\n 1. Listar contagem de colmeias\n 2. Adicionar colmeia\n 3. Adicionar várias colmeias\n 4. (em desenvolvimento) Escrever formulário de visita técnica\n 5. (em desenvolvimento) Listar tarefas pendentes")
	case "main_menu":
		switch body {
		case "1":
			// TODO: send list (getColmeias)
			sendReply(from, "Aqui está sua lista de colmeias!")
			clearState(from)
		case "2":
			// TODO: create colmeia (createColmeia)
			sendReply(from, "Vamos criar um registro de colmeia!")
			clearState(from)
		case "3":
			// TODO: create colmeia in batch (?)
			sendReply(from, "Vamos criar várias registro de colmeia!")
			clearState(from)
		}
	}
}

func sendReply(to, message string) error {
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", AppConfig.AccountSID)

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", AppConfig.FromNumber)
	msgData.Set("Body", message)

	req, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(msgData.Encode()))
	req.SetBasicAuth(AppConfig.AccountSID, AppConfig.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("Sent to %s: %s\n", to, message)
	return nil
}

func setState(user, currentState string) {
	log.Printf("Setting state from: %s to: %s", userMenuStates[user], currentState)
	stateLock.Lock()
	defer stateLock.Unlock()
	userMenuStates[user] = currentState
}

func clearState(user string) {
	log.Printf("Cleanin up state from: %s", userMenuStates[user])
	stateLock.Lock()
	defer stateLock.Unlock()
	delete(userMenuStates, user)
}
