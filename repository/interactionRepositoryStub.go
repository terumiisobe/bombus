package repository

import (
	"fmt"

	"bombus/domain"
)

type InteractionRepositoryStub struct {
	interactions []domain.Interaction
}

func NewInteractionRepositoryStub() InteractionRepositoryStub {
	interactions := []domain.Interaction{
		{domain.Init, ""},
		{domain.MainMenu, "Menu with the options: 1, 2, 3"},
		{domain.ListColmeias, "Colmeias list"},
		{domain.AddColmeiaForm, "Add Colmeia Form"},
		{domain.AddBatchColmeiaForm, "Add Batch Colmeia Form"},
		{domain.Success, "Success message"},
		{domain.Fail, "Fail message, error is: %s"},
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t domain.InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.Type == t {
			return interaction.DefaultText
		}
	}
	return ""
}

func (s InteractionRepositoryStub) GenerateText(t domain.InteractionType, additionalInfo string) string {
	for _, interaction := range s.interactions {
		if interaction.Type == domain.Fail {
			return fmt.Sprintf(interaction.DefaultText, additionalInfo)
		}
		if interaction.Type == t {
			return interaction.DefaultText
		}
	}
	return ""
}
