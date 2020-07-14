package utils

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type APIError struct {
	Model   string        `json:"model"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Errors  []interface{} `json:"errors"`
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
	e.Errors = append(e.Errors, err)
	return e
}

func DBError(model interface{}, fieldName string, err error, dberr error) APIError {
	e := APIError{}
	e.Code = http.StatusUnprocessableEntity
	e.Model = getModelName(model)
	e.Message = "Database error occurred."
	dbe := make(map[string]interface{})
	dbe["message"] = err.Error()
	dbe["field"] = getFieldTag(model, fieldName)
	dbe["details"] = dberr
	e.Errors = append(e.Errors, dbe)
	return e
}

func ValidatorError(model interface{}, vErrors []ValidationError) APIError {
	e := APIError{}
	e.Code = http.StatusUnprocessableEntity
	e.Model = getModelName(model)
	e.Message = "Validation error occurred."
	for _, vErr := range vErrors {
		ve := make(map[string]interface{})
		ve["field"] = getFieldTag(model, vErr.Field)
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
	e.Message = "Access forbidden."
	return e
}

func ResourceNotFound() APIError {
	e := APIError{}
	e.Code = http.StatusNotFound
	e.Message = "Resource not found."
	return e
}

func getModelName(model interface{}) string {
	return reflect.TypeOf(model).Elem().Name()
}

func getFieldTag(model interface{}, fieldName string) string {
	field, _ := reflect.ValueOf(model).Elem().Type().FieldByName(fieldName)
	tag, _ := field.Tag.Lookup("json")
	return tag
}
