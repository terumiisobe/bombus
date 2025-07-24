package service

type ChatbotService interface {
	GenerateReplyMessage(string, string) string
}
