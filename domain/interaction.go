package domain

type Interaction struct {
	typeName InteractionType
	text     string
}

type InteractionRepository interface {
	GetTextByType(InteractionType) string
	GenerateText(InteractionType, string) string
}

type InteractionType int

const (
	Init InteractionType = iota
	MainMenu
	ListColmeias
	AddColmeiaForm
	AddBatchColmeiaForm
	Success
	Fail
)
