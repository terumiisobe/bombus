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
		{chatbot.Init, ""},
		{chatbot.MainMenu, "Olá meliponicultor! O que deseja fazer?"},
		{chatbot.ListColmeias, "Esta é a sua listagem {com parâmetros x, y}. {lista}"},
		{chatbot.AddColmeiaForm, "Para adicionar uma colmeia, forneça as seguintes informações: número (opcional), espécie, status."},
		{chatbot.AddColmeiaValidation, " Uma nova colmeia será adicionada com as seguintes informações:{informações da colmeia}. Confirma?"},
		{chatbot.AddBatchColmeiaForm, "Add Batch Colmeia Form"},
		{chatbot.AddSuccess, "Colmeia criada com sucesso!"},
		{chatbot.AddFail, "Algumas informações não estão de acordo: {informações erradas/faltantes}."},
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t chatbot.InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.TypeName == t {
			return interaction.DefaultText
		}
	}
	return ""
}

func (s InteractionRepositoryStub) GenerateText(t chatbot.InteractionType, additionalInfo string) string {
	for _, interaction := range s.interactions {
		if interaction.TypeName == chatbot.AddFail {
			return fmt.Sprintf(interaction.DefaultText, additionalInfo)
		}
		if interaction.TypeName == t {
			return interaction.DefaultText
		}
	}
	return ""
}
