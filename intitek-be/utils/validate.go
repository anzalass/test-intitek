package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(input interface{}) error {
	err := validate.Struct(input)
	if err != nil {
		errMsg := "Validation errors:\n"
		for _, e := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("Field '%s' failed on '%s' rule\n", e.Field(), e.Tag())
		}
		return fmt.Errorf(errMsg)
	}
	return nil
}
