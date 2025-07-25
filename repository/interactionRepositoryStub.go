package repository

import (
	"bombus/domain"
	"fmt"
)

type InteractionRepositoryStub struct {
	interactions []domain.Interaction
}

func NewInteractionRepositoryStub() InteractionRepositoryStub {
	interactions := []domain.Interaction{
		{domain.Init, ""},
		{domain.MainMenu, "Olá meliponicultor! O que deseja fazer?"},
		{domain.ListColmeias, "Esta é a sua listagem {com parâmetros x, y}. {lista}"},
		{domain.AddColmeiaForm, "Para adicionar uma colmeia, forneça as seguintes informações: número (opcional), espécie, status."},
		{domain.AddColmeiaValidation, " Uma nova colmeia será adicionada com as seguintes informações:{informações da colmeia}. Confirma?"},
		{domain.AddBatchColmeiaForm, "Add Batch Colmeia Form"},
		{domain.AddSuccess, "Colmeia criada com sucesso!"},
		{domain.AddFail, "Algumas informações não estão de acordo: {informações erradas/faltantes}."},
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t domain.InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.TypeName == t {
			return interaction.DefaultText
		}
	}
	return ""
}

func (s InteractionRepositoryStub) GenerateText(t domain.InteractionType, additionalInfo string) string {
	for _, interaction := range s.interactions {
		if interaction.TypeName == domain.AddFail {
			return fmt.Sprintf(interaction.DefaultText, additionalInfo)
		}
		if interaction.TypeName == t {
			return interaction.DefaultText
		}
	}
	return ""
}
