package helpers

import (
	validator "github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) []string {
	var errors []string
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field())
		}
	}
	return errors
}
