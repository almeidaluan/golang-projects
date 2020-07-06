package utils

import "net/http"

type RestError struct {

	Message string `json:"message"`
	Status int `json:"status"`
	Error string  `json:"error"`
}

func BadRequestError(message string, error string) *RestError{
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		Error: error,
	}
}

func NotFoundError(message string, error string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   error,
	}
}

func NewBadRequestError(message string, error string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   error,
	}
}

func InternalServerError(message string, error string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   error,
	}
}
