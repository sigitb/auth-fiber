package user

import (
	"base-fiber/app/utils/rules"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateRegister(registerUser InputRegister) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(registerUser)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateLogin(LoginUser InputLogin) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(LoginUser)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateVerification(Verif InputVerification) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(Verif)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateEmail(FindEmail interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate.RegisterValidation("in", func(fl validator.FieldLevel) bool {
        allowedValues := strings.Split(fl.Param(), "+")
        return rules.RuleIn(allowedValues, fl.Field().String())
    })
	err := validate.Struct(FindEmail)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

