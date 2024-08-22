package validation

import (
	"encoding/json"
	"errors"
	"github.com/enioaires/crud-go/src/configuration/rest_err"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate *validator.Validate
	transl   ut.Translator
)

func init() {
	Validate = validator.New()

	eng := en.New()
	unicode := ut.New(eng, eng)
	var found bool
	transl, found = unicode.GetTranslator("en")
	if !found {
		panic("Translator not found")
	}

	if err := entranslation.RegisterDefaultTranslations(Validate, transl); err != nil {
		panic("Failed to register translations: " + err.Error())
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Field type invalid")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsCauses []rest_err.Causes

		for _, err := range jsonValidationError {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Field:   err.Field(),
				Message: err.Translate(transl),
			})
		}
		return rest_err.NewRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error validating request")
	}
}
