package test

import (
	"bombus/domain"
	"bombus/errs"
	"bombus/service"
	"reflect"
	"testing"
)

func TestChatbotService_GenerateMessage(t *testing.T) {

	r := domain.NewInteractionRepositoryStub()

	usr := "+5512345"

	t.Run("Empty strings, return menu", func(t *testing.T) {
		s := service.NewChatbotService(r)

		got := s.GenerateOutputMessageTDD("", "")
		want := r.GetTextByType(domain.MainMenu)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User without state, types anything, return menu", func(t *testing.T) {
		s := service.NewChatbotService(r)

		got := s.GenerateOutputMessageTDD(usr, "something")
		want := r.GetTextByType(domain.MainMenu)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User MainMenu state, types anything, return menu", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.MainMenu,
		}
		s := service.NewChatbotServiceCustomMap(r, m)

		got := s.GenerateOutputMessageTDD(usr, "something")
		want := r.GetTextByType(domain.MainMenu)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User MainMenu state, types types 1, return list", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.MainMenu,
		}
		s := service.NewChatbotServiceCustomMap(r, m)

		got := s.GenerateOutputMessageTDD(usr, "1")
		want := r.GetTextByType(domain.ListColmeias)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User MainMenu state, types 2, return add", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.MainMenu,
		}
		s := service.NewChatbotServiceCustomMap(r, m)

		got := s.GenerateOutputMessageTDD(usr, "2")
		want := r.GetTextByType(domain.AddColmeiaForm)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User MainMenu state, types 3, return add batch", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.MainMenu,
		}
		s := service.NewChatbotServiceCustomMap(r, m)

		got := s.GenerateOutputMessageTDD(usr, "3")
		want := r.GetTextByType(domain.AddBatchColmeiaForm)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User ListColmeias state, types anything, return MainMenu", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.ListColmeias,
		}
		s := service.NewChatbotServiceCustomMap(r, m)

		got := s.GenerateOutputMessageTDD(usr, "something")
		want := r.GetTextByType(domain.MainMenu)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User ListColmeias state, types 1, return MainMenu", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.ListColmeias,
		}
		s := service.NewChatbotServiceCustomMap(r, m)

		got := s.GenerateOutputMessageTDD(usr, "1")
		want := r.GetTextByType(domain.MainMenu)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User AddColmeiaForm state, types valid input, return success", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.AddColmeiaForm,
		}
		s := service.NewChatbotServiceCustomMap(r, m)
		validText := "123 \n1 \n01/05/2020 \npet"
		got := s.GenerateOutputMessageTDD(usr, validText)
		want := r.GetTextByType(domain.Success)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestChatbotService_ValidateText(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "")
		want := errs.NewValidationError("Texto vazio.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Any text, 1 line", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something")
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Any text, 3 lines (mandatory fields quantity)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something\nsomething\nsomething")
		want := errs.NewValidationError("Dados inválidos (something, something, something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Any text, 4 lines (mandatory+optional fields quantity)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something\nsomething\nsomething\nsomething")
		want := errs.NewValidationError("Dados inválidos (something, something, something, something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Any text, 5 lines", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something\nsomething\nsomething\nsomething\nsomething")
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Valid text, 3 lines", func(t *testing.T) {
		validText := "1\n" + "01/02/2020\n" + "1\n"

		got := service.ValidateText(domain.AddColmeiaForm, validText)

		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	})
}
