package domain_test

import (
	"bombus/domain"
	"bombus/errs"
	"reflect"
	"testing"
)

func TestCreateInteraction(t *testing.T) {

	const anyText = "anytext"
	//TODO: refactor ListColmeia tests creating const for inputs

	t.Run("Init validation accepts anytext.", func(t *testing.T) {
		interaction := domain.CreateInteractionInit()
		err := interaction.ValidateIn(anyText)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
	t.Run("MainMenu validation accepts anytext.", func(t *testing.T) {
		interaction := domain.CreateInteractionMainMenu()
		err := interaction.ValidateIn(anyText)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
	t.Run("ListColmeias validation accepts 1.", func(t *testing.T) {
		interaction := domain.CreateInteractionListColmeias()
		err := interaction.ValidateIn("1")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
	t.Run("ListColmeias validation accepts nothing other than 1.", func(t *testing.T) {
		interaction := domain.CreateInteractionListColmeias()
		gotErr := interaction.ValidateIn(anyText)
		expectedErr := errs.NewValidationError("%s não é válido para listagem de colmeias")
		if !reflect.DeepEqual(gotErr, expectedErr) {
			t.Errorf("Expected nil, got %v", gotErr)
		}
	})
}
