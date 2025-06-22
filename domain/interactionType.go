package domain

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
