package errors

import (
	goErrors "errors"
)

type ResponseError struct {
	Error string `json:"error"`
}

var ErrorMethodNotAllowed = goErrors.New("method not allowed")
