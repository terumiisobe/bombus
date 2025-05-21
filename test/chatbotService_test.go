package test

import (
	"bombus/service"
	"testing"
)

var s service.ChatbotService

func TestChatbotService_GenerateMessage(t *testing.T) {
	t.Run("Empty strings", func(t *testing.T) {
		got := s.GenerateMessage("", "")
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
