package api

import (
	serviceErrors "assets/internal/errors"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func JSONResponseError(w http.ResponseWriter, err error) {

	msg := err.Error()
	statusCode := http.StatusInternalServerError

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		msg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		statusCode = http.StatusBadRequest

	case errors.Is(err, io.ErrUnexpectedEOF):
		msg = fmt.Sprintf("Request body contains badly-formed JSON")
		statusCode = http.StatusBadRequest

	case errors.As(err, &unmarshalTypeError):
		msg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
		statusCode = http.StatusBadRequest

	case errors.Is(err, io.EOF):
		msg = "Request body must not be empty"
		statusCode = http.StatusBadRequest

	case errors.Is(err, serviceErrors.ErrorMethodNotAllowed):
		statusCode = http.StatusMethodNotAllowed

	case errors.Is(err, serviceErrors.InvalidLoginPasswordError):
	case errors.Is(err, serviceErrors.UnauthorizedError):
		statusCode = http.StatusUnauthorized

	case errors.Is(err, serviceErrors.NotFoundError):
		statusCode = http.StatusNotFound

	}

	data, err := json.Marshal(map[string]string{"error": msg})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}
