package utils

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Model   string                   `json:"model"`
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Errors  []map[string]interface{} `json:"errors"`
}

func NewError(code int, err error) APIError {
	e := APIError{}
	e.Code = code
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Message = v.Message.(string)
	default:
		e.Message = v.Error()
	}
	return e
}

func DBError(model interface{}, fieldName string, err error, dberr error) APIError {
	e := APIError{}

	modelName := reflect.TypeOf(model).Elem().Name()
	field, _ := reflect.ValueOf(model).Elem().Type().FieldByName(fieldName)
	tag, _ := field.Tag.Lookup("json")

	e.Code = http.StatusUnprocessableEntity
	e.Model = modelName
	e.Message = "Database error occurred."
	dbe := make(map[string]interface{})
	dbe["message"] = err.Error()
	dbe["field"] = tag
	dbe["details"] = dberr
	e.Errors = append(e.Errors, dbe)
	return e
}

func ValidatorError(vErrors []ValidationError) APIError {
	e := APIError{}
	e.Code = http.StatusUnprocessableEntity
	e.Model = "temp"
	e.Message = "Validation error occurred."
	for _, vErr := range vErrors {
		ve := make(map[string]interface{})
		ve["field"] = vErr.Field
		ve["condition"] = vErr.Condition
		ve["message"] = vErr.Message
		ve["conditionParameters"] = vErr.ConditionParameters
		ve["receivedValue"] = vErr.ReceivedValue
		e.Errors = append(e.Errors, ve)
	}
	return e
}

func Unauthorized() APIError {
	e := APIError{}
	e.Code = http.StatusUnauthorized
	e.Model = "temp"
	e.Message = "Access forbidden."
	return e
}

func ResourceNotFound() APIError {
	e := APIError{}
	e.Code = http.StatusNotFound
	e.Model = "temp"
	e.Message = "Resource not found."
	return e
}
