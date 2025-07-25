package domain

type Interaction struct {
	TypeName    InteractionType
	DefaultText string
}

type InteractionType int

const (
	Init InteractionType = iota
	MainMenu
	ListColmeias
	AddColmeiaForm
	AddColmeiaValidation
	AddBatchColmeiaForm
	AddSuccess
	AddFail
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
	case AddColmeiaValidation:
		return "AddColmeiaValidation"
	case AddBatchColmeiaForm:
		return "AddBatchColmeiaForm"
	case AddSuccess:
		return "Success"
	case AddFail:
		return "Fail"
	default:
		return "Unknown"
	}
}
