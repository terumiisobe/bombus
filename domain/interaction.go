package domain

type Interaction struct {
	Type        InteractionType
	DefaultText string
}

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
