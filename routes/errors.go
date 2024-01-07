package routes

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err     error  `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{Status: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{Status: http.StatusNotFound, Message: "Not found"}
	ErrBadRequest       = &ErrorResponse{Status: http.StatusBadRequest, Message: "Bad request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Status)
	return nil
}

func BadRequestRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Status:  http.StatusBadRequest,
		Message: err.Error(),
	}
}

func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}
