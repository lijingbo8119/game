package util

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func ValidateStruct(s interface{}) error {
	var (
		err error
	)
	if err = validate.Struct(s); err == nil {
		return nil
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		return err.(validator.ValidationErrors)[0]
	}
	if _, ok := err.(*validator.ValidationErrors); ok {
		return (*err.(*validator.ValidationErrors))[0]
	}
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}
	return err
}

func init() {
	validate = validator.New()
}
