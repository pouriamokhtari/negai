package helpers

import (
	"strings"

	validator "github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) []string {
	var errors []string
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := strings.Join([]string{err.Field(), err.ActualTag()}, ".")
			errors = append(errors, message)
		}
	}
	return errors
}
