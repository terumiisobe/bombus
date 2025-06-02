package test

import (
	"bombus/domain"
	"bombus/errs"
	"bombus/service"
	"reflect"
	"testing"
)

func TestChatbotService_GetNextInteraction(t *testing.T) {

	type want struct {
		result domain.InteractionType
		err    *errs.AppError
	}

	const (
		anyText                      = "anything"
		validTextAddColmeiaForm      = "123 \n1 \n01/05/2020 \n1"
		validTextAddBatchColmeiaForm = "1000 \n123 \n1 \n01/05/2020 \n1"
	)

	tests := []struct {
		name        string
		firstInput  domain.InteractionType
		secondInput string
		want        want
	}{
		{
			name:        "Current: Init | Input: any | Should return: MainMenu",
			firstInput:  domain.Init,
			secondInput: anyText,
			want: want{
				result: domain.MainMenu,
				err:    nil,
			},
		},
		{
			name:        "Current: MainMenu | Input: 1 | Should return: ListColmeias",
			firstInput:  domain.MainMenu,
			secondInput: "1",
			want: want{
				result: domain.ListColmeias,
				err:    nil,
			},
		},
		{
			name:        "Current: MainMenu | Input: 2 | Should return: AddColmeiaForm",
			firstInput:  domain.MainMenu,
			secondInput: "2",
			want: want{
				result: domain.AddColmeiaForm,
				err:    nil,
			},
		},
		{
			name:        "Current: MainMenu | Input: 3 | Should return: AddBatchColmeiaForm",
			firstInput:  domain.MainMenu,
			secondInput: "3",
			want: want{
				result: domain.AddBatchColmeiaForm,
				err:    nil,
			},
		},
		{
			name:        "Current: MainMenu | Input: any | Should return: MainMenu",
			firstInput:  domain.MainMenu,
			secondInput: anyText,
			want: want{
				result: domain.MainMenu,
				err:    errs.NewValidationError("Opção inválida."),
			},
		},
		{
			name:        "Current: ListColmeias | Input: any | Should return: Init",
			firstInput:  domain.ListColmeias,
			secondInput: anyText,
			want: want{
				result: domain.Init,
				err:    nil,
			},
		},
		{
			name:        "Current: AddColmeiaForm | Input:  valid text | Should return: Success",
			firstInput:  domain.AddColmeiaForm,
			secondInput: validTextAddColmeiaForm,
			want: want{
				result: domain.Success,
				err:    nil,
			},
		},
		{
			name:        "Current: AddColmeiaForm | Input:  invalid text | Should return: Fail",
			firstInput:  domain.AddColmeiaForm,
			secondInput: anyText,
			want: want{
				result: domain.Fail,
				err:    errs.NewValidationError("Número incorreto de linhas."),
			},
		},
		{
			name:        "Current: AddBatchColmeiaForm | Input:  valid text | Should return: Success",
			firstInput:  domain.AddBatchColmeiaForm,
			secondInput: validTextAddBatchColmeiaForm,
			want: want{
				result: domain.Success,
				err:    nil,
			},
		},
		{
			name:        "Current: AddBatchColmeiaForm | Input:  invalid text | Should return: Fail",
			firstInput:  domain.AddBatchColmeiaForm,
			secondInput: anyText,
			want: want{
				result: domain.Fail,
				err:    errs.NewValidationError("Número incorreto de linhas."),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotErr := service.GetNextInteraction(tc.firstInput, tc.secondInput)

			// 1) compare errors:
			if !errs.IsEqual(gotErr, tc.want.err) {
				t.Fatalf("unexpected error: got %v, want %v", gotErr, tc.want.err)
			}

			// 2) only if no error do we assert the result:
			if tc.want.err == nil {
				if !reflect.DeepEqual(gotResult, tc.want.result) {
					t.Errorf("unexpected result: got %#v, want %#v", gotResult, tc.want.result)
				}
			}
		})
	}
}
func TestChatbotService_ValidateInput(t *testing.T) {

	const validAddColmeiaFormValues string = "123 \n1 \n01/05/2020 \n1"
	const validAddBatchColmeiaFormValues string = "10 \n123 \n1 \n01/05/2020 \n1"
	const invalidInput string = "some invalid text"

	t.Run("Current: init | Input: any, should return nil", func(t *testing.T) {
		got := service.ValidateInput(domain.Init, "anything")
		var want *errs.AppError = nil

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want nil", got)
		}
	})
	t.Run("Current: MainMenu | Input: invalid, should return error", func(t *testing.T) {
		got := service.ValidateInput(domain.Init, invalidInput)
		want := errs.NewValidationError("Opção inválida.")

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Current: MainMenu | Input: valid, should return nil", func(t *testing.T) {
		valid := []string{"1", "2", "3"}

		for _, v := range valid {
			got := service.ValidateInput(domain.Init, v)
			if got != nil {
				t.Errorf("got %q want nil", got)
			}
		}
	})
	t.Run("Current: ListColmeias | Input: anything, should return nil", func(t *testing.T) {
		got := service.ValidateInput(domain.ListColmeias, invalidInput)
		var want *errs.AppError = nil

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want nil", got)
		}
	})
	t.Run("Current: AddColmeiaForm | Input: invalid, should return error", func(t *testing.T) {
		got := service.ValidateInput(domain.AddColmeiaForm, invalidInput)

		if got == nil {
			t.Errorf("got %q want error", got)
		}
	})
	t.Run("Current: AddColmeiaForm | Input: valid, should return nil", func(t *testing.T) {
		got := service.ValidateInput(domain.AddColmeiaForm, validAddColmeiaFormValues)
		var want *errs.AppError = nil

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want nil", got)
		}
	})
	t.Run("Current: AddColmeiaBatchForm | Input: invalid, should return error", func(t *testing.T) {
		got := service.ValidateInput(domain.AddBatchColmeiaForm, invalidInput)

		if got == nil {
			t.Errorf("got %q want error", got)
		}
	})
	t.Run("Current: AddColmeiaBatchForm | Input: valid, should return nil", func(t *testing.T) {
		got := service.ValidateInput(domain.AddBatchColmeiaForm, validAddBatchColmeiaFormValues)
		var want *errs.AppError = nil

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want nil", got)
		}
	})
}
func TestChatbotService_ValidateForm(t *testing.T) {
	const (
		wrongValue     = "this is a wrong value"
		wrongEnumValue = "1000"
		quantity       = "1"
		QRCode         = "123"
		species        = "1"
		startingDate   = "01/02/2020"
		status         = "1"
	)

	t.Run("[AddColmeiaForm] Insufficient fields", func(t *testing.T) {
		formValues := []string{
			QRCode,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Excessive fields", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			QRCode,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Wrong data (all wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			wrongValue,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value, this is a wrong value, this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Wrong data (first wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Wrong data (second wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			species,
			wrongValue,
			status,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Wrong data (third wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			species,
			startingDate,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Valid data (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		if got != nil {
			t.Errorf("ValidateForm() = %v, want nil", got)
		}
	})

	t.Run("[AddColmeiaForm] Wrong data (all wrong) (all fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			wrongValue,
			wrongValue,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value, this is a wrong value, this is a wrong value, this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Wrong data (first wrong) (all fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Wrong enum data", func(t *testing.T) {
		formValues := []string{
			QRCode,
			wrongEnumValue,
			startingDate,
			wrongEnumValue,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (1000, 1000).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddColmeiaForm] Valid data  (all fields)", func(t *testing.T) {
		formValues := []string{
			QRCode,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddColmeiaForm, formValues)
		if got != nil {
			t.Errorf("ValidateForm() = %v, want nil", got)
		}
	})
	t.Run("[AddBatchColmeiaForm] Insufficient fields", func(t *testing.T) {
		formValues := []string{
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Excessive fields", func(t *testing.T) {
		formValues := []string{
			quantity,
			QRCode,
			species,
			startingDate,
			status,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Número incorreto de linhas.")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong data (all wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			wrongValue,
			wrongValue,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value, this is a wrong value, this is a wrong value, this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong data (first wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong data (second wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			quantity,
			wrongValue,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong data (third wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			quantity,
			species,
			wrongValue,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})
	t.Run("[AddBatchColmeiaForm] Wrong data (forth wrong) (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			quantity,
			species,
			startingDate,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Valid data (only mandatory fields)", func(t *testing.T) {
		formValues := []string{
			quantity,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		if got != nil {
			t.Errorf("ValidateForm() = %v, want nil", got)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong data (all wrong) (all fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			wrongValue,
			wrongValue,
			wrongValue,
			wrongValue,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value, this is a wrong value, this is a wrong value, this is a wrong value, this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong data (first wrong) (all fields)", func(t *testing.T) {
		formValues := []string{
			wrongValue,
			QRCode,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (this is a wrong value).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Wrong enum data", func(t *testing.T) {
		formValues := []string{
			quantity,
			QRCode,
			wrongEnumValue,
			startingDate,
			wrongEnumValue,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		want := errs.NewValidationError("Dados inválidos (1000, 1000).")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ValidateForm() = %v, want %v", got, want)
		}
	})

	t.Run("[AddBatchColmeiaForm] Valid data  (all fields)", func(t *testing.T) {
		formValues := []string{
			quantity,
			QRCode,
			species,
			startingDate,
			status,
		}

		got := service.ValidateForm(domain.AddBatchColmeiaForm, formValues)
		if got != nil {
			t.Errorf("ValidateForm() = %v, want nil", got)
		}
	})
}

// func TestChatbotService_GenerateMessage(t *testing.T) {

// 	r := domain.NewInteractionRepositoryStub()

// 	usr := "+5512345"

// 	t.Run("Empty strings, return menu", func(t *testing.T) {
// 		s := service.NewChatbotService(r)

// 		got := s.GenerateOutputMessageTDD("", "")
// 		want := r.GetTextByType(domain.MainMenu)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User without state, types anything, return menu", func(t *testing.T) {
// 		s := service.NewChatbotService(r)

// 		got := s.GenerateOutputMessageTDD(usr, "something")
// 		want := r.GetTextByType(domain.MainMenu)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User MainMenu state, types anything, return menu", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.MainMenu,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)

// 		got := s.GenerateOutputMessageTDD(usr, "something")
// 		want := r.GetTextByType(domain.MainMenu)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User MainMenu state, types 1, return list", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.MainMenu,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)

// 		got := s.GenerateOutputMessageTDD(usr, "1")
// 		want := r.GetTextByType(domain.ListColmeias)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User MainMenu state, types 2, return add form", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.MainMenu,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)

// 		got := s.GenerateOutputMessageTDD(usr, "2")
// 		want := r.GetTextByType(domain.AddColmeiaForm)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User MainMenu state, types 3, return add batch form", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.MainMenu,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)

// 		got := s.GenerateOutputMessageTDD(usr, "3")
// 		want := r.GetTextByType(domain.AddBatchColmeiaForm)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User ListColmeias state, types anything, return MainMenu", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.ListColmeias,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)

// 		got := s.GenerateOutputMessageTDD(usr, "anything")
// 		want := r.GetTextByType(domain.MainMenu)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User AddColmeiaForm state, types valid input, return success", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.AddColmeiaForm,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)
// 		validText := "123 \n1 \n01/05/2020 \n1"
// 		got := s.GenerateOutputMessageTDD(usr, validText)
// 		want := r.GetTextByType(domain.Success)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User AddColmeiaForm state, types invalid input, return fail with error message", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.AddColmeiaForm,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)
// 		invalidText := "\n1000 \n01/05/2020 \n1"
// 		got := s.GenerateOutputMessageTDD(usr, invalidText)
// 		want := r.GenerateText(domain.Fail, "Dados inválidos (1000).")

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User AddBatchColmeiaForm state, types valid input (5 values), return success", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.AddBatchColmeiaForm,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)
// 		validText := "10\n123 \n1 \n01/05/2020 \n1"
// 		got := s.GenerateOutputMessageTDD(usr, validText)
// 		want := r.GetTextByType(domain.Success)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("User AddBatchColmeiaForm state, types invalid input (5 values), return fail", func(t *testing.T) {
// 		m := map[string]domain.InteractionType{
// 			usr: domain.AddBatchColmeiaForm,
// 		}
// 		s := service.NewChatbotServiceCustomMap(r, m)
// 		validText := "12\n123 \n10 \n01/05/2020 \n1"
// 		got := s.GenerateOutputMessageTDD(usr, validText)
// 		want := r.GenerateText(domain.Fail, "Dados inválidos (10).")

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }
