package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(code int, err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["code"] = code
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func DuplicateError(code int, err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["code"] = code
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func NewValidatorError(vErrors []ValidationError) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["code"] = http.StatusUnprocessableEntity
	e.Errors["validationErrors"] = vErrors
	return e
}

func Unauthorized() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["code"] = http.StatusUnauthorized
	e.Errors["body"] = "Access forbidden."
	return e
}

func ResourceNotFound() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["code"] = http.StatusNotFound
	e.Errors["body"] = "Resource not found."
	return e
}
