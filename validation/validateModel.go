package validation

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidateCriteria(model interface{}) *ValidationError {
	validate := validator.New()

	err := validate.Struct(model)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	errors := make([]ValidationErrorDetail, 0)

	for _, fieldErr := range validationErrors {
		errors = append(errors, ValidationErrorDetail{
			Field:   getJSONFieldName(model, fieldErr.Field()),
			Message: validationMessage(fieldErr),
		})
	}

	return &ValidationError{Errors: errors}
}

func getJSONFieldName(model interface{}, fieldName string) string {
	t := reflect.TypeOf(model)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := fieldName

	if field, ok := t.FieldByName(fieldName); ok {
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			name = strings.Split(jsonTag, ",")[0]
		}
	}

	return camelToTitle(name)
}

func validationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "field is required"
	case "min":
		return fmt.Sprintf("minimum length is %s", err.Param())
	case "max":
		return fmt.Sprintf("maximum length is %s", err.Param())
	case "email":
		return "invalid email format"
	default:
		return "invalid field value"
	}
}

func camelToTitle(input string) string {
	if input == "" {
		return ""
	}

	var result []rune
	runes := []rune(input)

	for i, r := range runes {
		// jika huruf besar dan bukan index 0 → tambah spasi
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, ' ')
		}
		result = append(result, r)
	}

	return strings.Title(string(result))
}

func ValidateUUID(id string) error {
	_, err := uuid.Parse(id)
	return err
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []ValidationErrorDetail `json:"errors"`
}

func (e ValidationError) Error() string {
	return "validation error"
}
