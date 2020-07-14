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

func DBError(code int, model string, field string, err error, dberr error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["code"] = code
	e.Errors["body"] = err.Error()
	e.Errors["model"] = model
	e.Errors["field"] = field
	e.Errors["details"] = dberr
	return e
}

func ValidatorError(vErrors []ValidationError) Error {
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
