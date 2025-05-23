package domain

type InteractionRepositoryStub struct {
	interactions []Interaction
}

func NewInteractionRepositoryStub() InteractionRepositoryStub {
	interactions := []Interaction{
		{MainMenu, "Menu with the options: 1, 2, 3"},
		{ListColmeias, "Colmeias list"},
		{AddColmeiaForm, "Add Colmeia Form"},
		{AddBatchColmeiaForm, "Add Batch Colmeia Form"},
		{Success, "Success message"},
		{Fail, "Fail message"},
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.typeName == t {
			return interaction.text
		}
	}
	return ""
}
