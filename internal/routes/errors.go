package routes

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"-"`
	Error  string `json:"error"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.Status)
	json.NewEncoder(w).Encode(e)
	return nil
}

func ApplicationError(message string) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		Error:  message,
	}
}

func BadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusBadRequest,
		Error:  message,
	}
}

func MethodNotAllowedError() *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusMethodNotAllowed,
		Error:  "Method Not Allowed",
	}
}

func NotFoundError() *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusNotFound,
		Error:  "Not Found",
	}
}
