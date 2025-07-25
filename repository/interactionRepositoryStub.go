package repository

import (
	"bombus/domain/chatbot"
	"fmt"
)

type InteractionRepositoryStub struct {
	interactions []chatbot.Interaction
}

func NewInteractionRepositoryStub() InteractionRepositoryStub {
	interactions := []chatbot.Interaction{
		chatbot.CreateMainMenuInteraction(),
		chatbot.CreateListColmeiasInteraction(),
		chatbot.CreateAddColmeiaInteraction(),
		chatbot.CreateAddColmeiaValidationInteraction(),
		chatbot.CreateAddBatchColmeiaInteraction(),
		chatbot.CreateAddColmeiaSuccessInteraction(),
		chatbot.CreateAddColmeiaFailInteraction(),
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t chatbot.InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.Name == t {
			return interaction.DefaultText
		}
	}
	return ""
}

func (s InteractionRepositoryStub) GenerateText(t chatbot.InteractionType, additionalInfo string) string {
	for _, interaction := range s.interactions {
		if interaction.Name == chatbot.AddColmeiaFail {
			return fmt.Sprintf(interaction.DefaultText, additionalInfo)
		}
		if interaction.Name == t {
			return interaction.DefaultText
		}
	}
	return ""
}
