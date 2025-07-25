package chatbot

type Interaction struct {
	Name        InteractionType   `json:"name"`
	DefaultText string            `json:"default_text"`
	Params      map[string]string `json:"params"`
}

type InteractionType string

const (
	Init                 InteractionType = "init"
	MainMenu             InteractionType = "main_menu"
	ListColmeias         InteractionType = "list_colmeias"
	AddColmeiaForm       InteractionType = "add_colmeia"
	AddColmeiaValidation InteractionType = "add_colmeia_validation"
	AddBatchColmeia      InteractionType = "add_batch_colmeia"
	AddColmeiaSuccess    InteractionType = "add_success"
	AddColmeiaFail       InteractionType = "add_fail"
)

func CreateMainMenuInteraction() Interaction {
	return Interaction{
		Name:        MainMenu,
		DefaultText: "Olá meliponicultor! O que deseja fazer?",
		Params:      map[string]string{},
	}
}

func CreateListColmeiasInteraction() Interaction {
	return Interaction{
		Name:        ListColmeias,
		DefaultText: "Esta é a sua listagem: %s",
		Params:      map[string]string{},
	}
}

func CreateAddColmeiaInteraction() Interaction {
	return Interaction{
		Name:        AddColmeiaForm,
		DefaultText: "Para adicionar uma colmeia, forneça as seguintes informações: número (opcional), espécie, status.",
		Params:      map[string]string{},
	}
}

func CreateAddColmeiaValidationInteraction() Interaction {
	return Interaction{
		Name:        AddColmeiaValidation,
		DefaultText: "Uma nova colmeia será adicionada com as seguintes informações: %s. Confirma?",
		Params:      map[string]string{},
	}
}

func CreateAddBatchColmeiaInteraction() Interaction {
	return Interaction{
		Name:        AddBatchColmeia,
		DefaultText: "Para adicionar uma colmeia, forneça as seguintes informações: número (opcional), espécie, status.",
		Params:      map[string]string{},
	}
}

func CreateAddColmeiaSuccessInteraction() Interaction {
	return Interaction{
		Name:        AddColmeiaSuccess,
		DefaultText: "Colmeia criada com sucesso!",
		Params:      map[string]string{},
	}
}

func CreateAddColmeiaFailInteraction() Interaction {
	return Interaction{
		Name:        AddColmeiaFail,
		DefaultText: "Algumas informações não estão de acordo: %s.",
		Params:      map[string]string{},
	}
}
