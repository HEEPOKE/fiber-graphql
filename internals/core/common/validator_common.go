package common

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateCommon[T any](payload T) string {
	var errorMessages []string
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s %s %s", err.StructField(), err.Tag(), err.Param())
			errorMessages = append(errorMessages, message)
		}
	}
	return strings.Join(errorMessages, "; ")
}
