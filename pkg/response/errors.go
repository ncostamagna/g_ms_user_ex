package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ErrorResponse) GetData() interface{} {
	return e.Errors
}

func InternalServerError(msg string) Response {
	return makeErrorResponse(msg, nil, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return makeErrorResponse(msg, nil, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return makeErrorResponse(msg, nil, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return makeErrorResponse(msg, nil, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return makeErrorResponse(msg, nil, http.StatusBadRequest)
}

func InvalidInput(msg string, errors interface{}) Response {
	return makeErrorResponse(msg, errors, http.StatusBadRequest)
}

func makeErrorResponse(msg string, errors interface{}, status int) Response {
	return &ErrorResponse{
		Status:  status,
		Message: msg,
		Errors:  errors,
	}
}
