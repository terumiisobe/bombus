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
	t.Run("User MainMenu state, types 2, return add form", func(t *testing.T) {
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
	t.Run("User MainMenu state, types 3, return add batch form", func(t *testing.T) {
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

		got := s.GenerateOutputMessageTDD(usr, "anything")
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
		validText := "123 \n1 \n01/05/2020 \n1"
		got := s.GenerateOutputMessageTDD(usr, validText)
		want := r.GetTextByType(domain.Success)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User AddColmeiaForm state, types invalid input, return fail with error message", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.AddColmeiaForm,
		}
		s := service.NewChatbotServiceCustomMap(r, m)
		invalidText := "\n1000 \n01/05/2020 \n1"
		got := s.GenerateOutputMessageTDD(usr, invalidText)
		want := r.GenerateText(domain.Fail, "Dados inválidos (1000).")

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("User AddBatchColmeiaForm state, types valid input (5 values), return success", func(t *testing.T) {
		m := map[string]domain.InteractionType{
			usr: domain.AddBatchColmeiaForm,
		}
		s := service.NewChatbotServiceCustomMap(r, m)
		validText := "10\n123 \n1 \n01/05/2020 \n1"
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
	// InteractionType = AddColmeiaForm
	t.Run("[AddColmeiaForm] Wrong data quantity (less)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something")
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data quantity (more)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something\nsomething\nsomething\nsomething\nsomething")
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (all wrong in type), right data amount (only mandatory fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something\nsomething\nsomething")
		want := errs.NewValidationError("Dados inválidos (something, something, something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (1 wrong in type, 2 right), right data amount (only mandatory fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something\n01/02/2020\n1")
		want := errs.NewValidationError("Dados inválidos (something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (2 wrong in type, 1 right), right data amount (only mandatory fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "1\nsomething\nsome other thing")
		want := errs.NewValidationError("Dados inválidos (something, some other thing).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Valid data, right data amount (only mandatory fields)", func(t *testing.T) {
		validText := "1\n" + "01/02/2020\n" + "1\n"

		got := service.ValidateText(domain.AddColmeiaForm, validText)

		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (all wrong in type), right data amount (all fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "something A\nsomething B\nsomething C\nsomething D")
		want := errs.NewValidationError("Dados inválidos (something A, something B, something C, something D).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (1 wrong in type, 3 right), right data amount (all fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "123\n1\n01/02/2020\nsomething")
		want := errs.NewValidationError("Dados inválidos (something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (2 wrong in type, 2 right), right data amount (all fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "123\nsomething A\n01/02/2020\nsomething B")
		want := errs.NewValidationError("Dados inválidos (something A, something B).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Wrong data (2 wrong in enum, 2 right), right data amount (all fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddColmeiaForm, "123\n100\n01/02/2020\n50")
		want := errs.NewValidationError("Dados inválidos (100, 50).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddColmeiaForm] Valid data, right data amount (all fields)", func(t *testing.T) {
		validText := "123\n" + "2\n" + "01/02/2020\n" + "1\n"

		got := service.ValidateText(domain.AddColmeiaForm, validText)

		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	})

	// InteractionType = AddBatchColmeiaForm
	t.Run("[AddBatchColmeiaForm] Wrong data quantity (less)", func(t *testing.T) {
		got := service.ValidateText(domain.AddBatchColmeiaForm, "something\nsomething\nsomething")
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Wrong data quantity (more)", func(t *testing.T) {
		got := service.ValidateText(domain.AddBatchColmeiaForm, "something\nsomething\nsomething\nsomething\nsomething\nsomething")
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Wrong data (all wrong in type), right data amount (only mandatory fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddBatchColmeiaForm, "something\nsomething\nsomething\nsomething")
		want := errs.NewValidationError("Dados inválidos (something, something, something, something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Wrong data (1 wrong in type, 3 right), right data amount (only mandatory fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddBatchColmeiaForm, "something\n1\n01/02/2020\n1")
		want := errs.NewValidationError("Dados inválidos (something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Valid data, right data amount (only mandatory fields)", func(t *testing.T) {
		validText := "10\n" + "1\n" + "01/02/2020\n" + "1\n"

		got := service.ValidateText(domain.AddBatchColmeiaForm, validText)

		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	})
	t.Run("[AddBatchColmeiaForm] Wrong data (all wrong in type), right data amount (all fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddBatchColmeiaForm, "something A\nsomething B\nsomething C\nsomething D\nsomething E")
		want := errs.NewValidationError("Dados inválidos (something A, something B, something C, something D, something E).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Wrong data (1 wrong in type, 4 right), right data amount (all fields)", func(t *testing.T) {
		got := service.ValidateText(domain.AddBatchColmeiaForm, "something\n12345\n1\n01/02/2020\n1")
		want := errs.NewValidationError("Dados inválidos (something).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Valid data, right data amount (all fields)", func(t *testing.T) {
		validText := "3\n" + "123\n" + "2\n" + "01/02/2020\n" + "1\n"

		got := service.ValidateText(domain.AddBatchColmeiaForm, validText)

		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	})
}
