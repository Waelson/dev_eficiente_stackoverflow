package api

import (
	"fmt"
	"net/http"
)

type Error interface {
	Message() string
	Error() string
	Status() int
}

type apiError struct {
	ErrorMessage  string `json:"message"`
	ErrorStatus   int    `json:"status"`
	ErrorOriginal string `json:"-"`
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func (e apiError) Error() string {
	return fmt.Sprintf("Message: %s;Status: %d", e.ErrorMessage, e.ErrorStatus)
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func NewDatabaseError(err error) Error {
	return apiError{
		ErrorMessage:  err.Error(),
		ErrorStatus:   http.StatusInternalServerError,
		ErrorOriginal: err.Error(),
	}
}

func NewServiceError(message string) Error {
	return apiError{
		ErrorMessage:  message,
		ErrorStatus:   http.StatusBadRequest,
		ErrorOriginal: message,
	}
}

func NewInternalError(err error) Error {
	return apiError{
		ErrorMessage:  err.Error(),
		ErrorStatus:   http.StatusBadRequest,
		ErrorOriginal: err.Error(),
	}
}

func NewApiError(message string, error string, status int) Error {
	return apiError{
		ErrorMessage:  message,
		ErrorStatus:   status,
		ErrorOriginal: error,
	}
}
