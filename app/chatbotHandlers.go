package app

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"bombus/app/chatbot"
)

var (
	userMenuStates = make(map[string]string)
	stateLock      sync.Mutex
)

func NewChatbotHandler() *ChatbotHandler {
    handler := ChatbotHandler{}
    return &handler
}
