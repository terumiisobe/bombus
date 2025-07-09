package chatbot

type ActionName string

const (
	ListColmeia ActionName = "list_colmeia"
	AddColmeia  ActionName = "add_colmeia"
)

type Action struct {
	Name   ActionName        `json:"name"`
	Params map[string]string `json:"params"`
}
