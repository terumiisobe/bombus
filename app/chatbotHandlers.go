package app

import (
	"bombus/service"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// deals with webhook response
// uses userState
// uses messageGenerator

type ChatbotHandler struct {
	chatbotService service.ChatbotService
}

func (wh *ChatbotHandler) handle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	from := r.FormValue("From")
	body := r.FormValue("Body")

	log.Printf("From: %s | Message: %s \n", from, body)

	// TODO: uncomment and update GenerateMessage()
	//replyMessage := wh.chatbotService.GenerateMessage(from, body)
	replyMessage := ""
	sendReply(from, replyMessage)
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
