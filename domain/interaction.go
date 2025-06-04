package domain

type interaction interface {
	GenerateText(string) string
}

type Interaction struct {
	typeName    InteractionType
	defaultText string
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
