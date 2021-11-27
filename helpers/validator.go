package helpers

import (
	validator "github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
}

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}
