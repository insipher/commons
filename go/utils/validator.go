package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Validator defines the structure of the Validator object
type Validator struct {
	validator *validator.Validate
}

// ValidationError represents a structured validation error object to be used on responses
type ValidationError struct {
	Message             string
	Field               string
	Condition           string
	ConditionParameters string
	ReceivedValue       interface{}
}

// NewValidator returns create a new validator object
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate runs validation over a given object. Validation rules are defined on the model (struct) object.
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func GetValidationErrors(err validator.ValidationErrors) []ValidationError {
	var vErrors []ValidationError

	for _, e := range err {
		msg := getValidationMessage(e.Field(), e.ActualTag())
		vErrors = append(vErrors, ValidationError{
			Message:             msg,
			Field:               e.Field(),
			Condition:           e.ActualTag(),
			ConditionParameters: e.Param(),
			ReceivedValue:       e.Value(),
		})
	}

	return vErrors
}

func getValidationMessage(field string, tag string) string {
	var msg string

	switch tag {
	case "required":
		msg = fmt.Sprintf("%s is a required field.", field)
	case "email":
		msg = fmt.Sprintf("%s must be a valid email.", field)
	case "passwd":
		msg = fmt.Sprintf("%s must be a valid password.", field)
	}

	return msg
}
