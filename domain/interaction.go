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

func (i InteractionType) String() string {
	switch i {
	case Init:
		return "Init"
	case MainMenu:
		return "MainMenu"
	case ListColmeias:
		return "ListColmeias"
	case AddColmeiaForm:
		return "AddColmeiaForm"
	case AddBatchColmeiaForm:
		return "AddBatchColmeiaForm"
	case Success:
		return "Success"
	case Fail:
		return "Fail"
	default:
		return "Unknown"
	}
}
