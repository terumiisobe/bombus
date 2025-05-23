package domain

type Interaction struct {
	typeName InteractionType
	text     string
}

type InteractionRepository interface {
	GetTextByType(InteractionType) string
}

type InteractionType int

const (
	MainMenu InteractionType = iota + 1
	ListColmeias
	AddColmeiaForm
	AddBatchColmeiaForm
	Success
	Fail
)

var InteractionDependencyMap = map[InteractionType]map[string]InteractionType{
	MainMenu: {
		"1": ListColmeias,
		"2": AddColmeiaForm,
		"3": AddBatchColmeiaForm},
	ListColmeias:        {},
	AddColmeiaForm:      {},
	AddBatchColmeiaForm: {},
	Success:             {},
	Fail:                {},
}
