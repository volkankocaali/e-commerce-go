package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"unicode"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

type Validator struct {
	validate *validator.Validate
}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// ValidateStruct performs validation on a struct and returns validation errors
func (v *Validator) ValidateStruct(s interface{}) []ValidationError {
	var validationErrors []ValidationError

	// Register custom validation functions
	err := v.validate.RegisterValidation("customValidation", customValidation)
	if err != nil {
		return nil
	}

	if err := v.validate.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := formatFieldName(err.Field())
			message := err.Tag()

			switch err.Tag() {
			case "min":
				message = customMinMessage(field, err.ActualTag(), err.Param())
			case "max":
				message = customMaxMessage(field, err.ActualTag(), err.Param())
			case "required":
				message = field + " is required"
			}
			validationError := ValidationError{
				Field: field,
				Tag:   message,
			}

			validationErrors = append(validationErrors, validationError)
		}
	}

	return validationErrors
}

// formatFieldName is a helper function to format field names
func formatFieldName(fieldName string) string {
	var result strings.Builder
	for i, c := range fieldName {
		if i > 0 && unicode.IsUpper(c) && unicode.IsLower(rune(fieldName[i-1])) {
			result.WriteRune(' ')
		}
		result.WriteRune(c)
	}
	return result.String()
}

// customValidation is a custom validation function example
func customValidation(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", fl.Field().String())
	return match
}

// customMinMessage is a custom error message for "min" validation
func customMinMessage(field string, tag string, param string) string {
	return field + " must be at least " + param + " characters long."
}

// customMaxMessage is a custom error message for "max" validation
func customMaxMessage(field string, tag string, param string) string {
	return field + " must be at most " + param + " characters long."
}
