package model

import (
	"fmt"
	"net/http"
)

type AppError struct {
	code    int
	message string
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Status() int {
	return e.code
}

func NewBadRequestError(reason string) *AppError {
	return &AppError{
		code:    http.StatusBadRequest,
		message: fmt.Sprintf("Bad request. %v", reason),
	}
}

func NewUnauthorizedError(isTokenProvided bool) *AppError {
	if isTokenProvided {
		return &AppError{
			code:    http.StatusUnauthorized,
			message: fmt.Sprintf("Unauthorized. Authentication failed."),
		}
	}

	return &AppError{
		code:    http.StatusUnauthorized,
		message: fmt.Sprintf("Unauthorized. Authorization token needed but not provided."),
	}
}

func NewForbiddenError() *AppError {
	return &AppError{
		code:    http.StatusForbidden,
		message: fmt.Sprintf("Forbidden. You are not authorized to access this resource."),
	}
}

func NewNotFoundError(name string) *AppError {
	return &AppError{
		code:    http.StatusNotFound,
		message: fmt.Sprintf("Not found. %v not found", name),
	}
}

func NewConflictError(name string) *AppError {
	return &AppError{
		code:    http.StatusConflict,
		message: fmt.Sprintf("Conflict. %v already exists", name),
	}
}

func NewUnsupportedMediaTypeError(reason string) *AppError {
	return &AppError{
		code:    http.StatusUnsupportedMediaType,
		message: fmt.Sprintf("Unsupported media type. %v", reason),
	}
}

func NewInternalServerError() *AppError {
	return &AppError{
		code:    http.StatusInternalServerError,
		message: fmt.Sprintf("Internal server error."),
	}
}
