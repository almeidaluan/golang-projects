package utils

import "net/http"

type RestError struct {

	Message string `json:"message"`
	Status int `json:"status"`
	Error error  `json:"error"`
}

func BadRequestError(message string, error error) *RestError{
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		Error: error,
	}
}
