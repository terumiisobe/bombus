package domain

import "fmt"

type InteractionRepositoryStub struct {
	interactions []Interaction
}

func NewInteractionRepositoryStub() InteractionRepositoryStub {
	interactions := []Interaction{
		{Init, ""},
		{MainMenu, "Menu with the options: 1, 2, 3"},
		{ListColmeias, "Colmeias list"},
		{AddColmeiaForm, "Add Colmeia Form"},
		{AddBatchColmeiaForm, "Add Batch Colmeia Form"},
		{Success, "Success message"},
		{Fail, "Fail message, error is: %s"},
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.typeName == t {
			return interaction.defaultText
		}
	}
	return ""
}

func (s InteractionRepositoryStub) GenerateText(t InteractionType, additionalInfo string) string {
	for _, interaction := range s.interactions {
		if interaction.typeName == Fail {
			return fmt.Sprintf(interaction.defaultText, additionalInfo)
		}
		if interaction.typeName == t {
			return interaction.defaultText
		}
	}
	return ""
}
