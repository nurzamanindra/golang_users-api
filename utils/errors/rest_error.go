package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
