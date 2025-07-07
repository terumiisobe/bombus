package domain

import "bombus/errs"

type Interaction struct {
	Type        InteractionType
	DefaultText string
	ValidateIn  func(in string) *errs.AppError
}

func CreateInteractionInit() Interaction {
	return Interaction{
		Type:        Init,
		DefaultText: "",
		ValidateIn: func(in string) *errs.AppError {
			return nil
		},
	}
}

func CreateInteractionMainMenu() Interaction {
	return Interaction{
		Type:        MainMenu,
		DefaultText: "MainMenu",
		ValidateIn: func(in string) *errs.AppError {
			return nil
		},
	}
}

func CreateInteractionListColmeias() Interaction {
	return Interaction{
		Type:        ListColmeias,
		DefaultText: "Listagem de colmeias",
		ValidateIn: func(in string) *errs.AppError {
			if in != "1" {
				return errs.NewValidationError("%s não é válido para listagem de colmeias")
			}
			return nil
		},
	}
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
