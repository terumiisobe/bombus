package repository

import (
	"fmt"

	"bombus/domain"
	"bombus/errs"
)

type InteractionRepositoryStub struct {
	interactions []domain.Interaction
}

func NewInteractionRepositoryStub() InteractionRepositoryStub {
	interactions := []domain.Interaction{
		{domain.Init, "", func(in string) *errs.AppError {
			return nil
		}},
		{domain.MainMenu, "Menu with the options: 1, 2, 3", func(in string) *errs.AppError {
			return nil
		}},
		{domain.ListColmeias, "Colmeias list", func(in string) *errs.AppError {
			return nil
		}},
		{domain.AddColmeiaForm, "Add Colmeia Form", func(in string) *errs.AppError {
			return nil
		}},
		{domain.AddBatchColmeiaForm, "Add Batch Colmeia Form", func(in string) *errs.AppError {
			return nil
		}},
		{domain.Success, "Success message", func(in string) *errs.AppError {
			return nil
		}},
		{domain.Fail, "Fail message, error is: %s", func(in string) *errs.AppError {
			return nil
		}},
	}

	return InteractionRepositoryStub{interactions}
}

func (s InteractionRepositoryStub) GetTextByType(t domain.InteractionType) string {
	for _, interaction := range s.interactions {
		if interaction.Type == t {
			return interaction.DefaultText
		}
	}
	return ""
}

func (s InteractionRepositoryStub) GenerateText(t domain.InteractionType, additionalInfo string) string {
	for _, interaction := range s.interactions {
		if interaction.Type == domain.Fail {
			return fmt.Sprintf(interaction.DefaultText, additionalInfo)
		}
		if interaction.Type == t {
			return interaction.DefaultText
		}
	}
	return ""
}
