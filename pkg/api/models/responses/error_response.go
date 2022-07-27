package models

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

// ErrorResponse
//
// A error response model when an error occured.
//
// swagger:model
type ErrorResponse struct {
	// Error message
	Message string `json:"title"`
	// Error code
	Code int `json:"code"`
	// Error details
	Errors []InnerErrorResponse `json:"errors"`
}

type InnerErrorResponse struct {
	// Field name
	FieldName string `json:"fieldName"`
	// Error message
	ErrorMessage string `json:"errorMessage"`
	// Attempted value
	AttemptedValue interface{} `json:"attempedValue"`
}

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    code,
	}
}

func NewErrorResponseFromNotFoundResource() *ErrorResponse {
	return &ErrorResponse{
		Message: "Resource not found",
		Code:    http.StatusNotFound,
	}
}

func NewErrorResponseFromBadRequest(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
		Code:    http.StatusBadRequest,
	}
}

func NewErrorResponseFromError(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	}
}

func NewErrorResponseFromValidationErrors(validationErrors validator.ValidationErrors) *ErrorResponse {
	errors := make([]InnerErrorResponse, len(validationErrors))

	for i, fe := range validationErrors {
		errors[i] = InnerErrorResponse{
			FieldName:      fe.Field(),
			AttemptedValue: fe.Value(),
			ErrorMessage:   "Invalid field.",
		}
	}

	errorResponse := &ErrorResponse{
		Message: "Request body validation failed.",
		Code:    http.StatusBadRequest,
		Errors:  errors,
	}

	return errorResponse
}

func (er *ErrorResponse) AsJson() (string, error) {
	json, err := json.Marshal(&er)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
