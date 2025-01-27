package handler

import (
	"errors"
	"net/http"
)

func Validation(w http.ResponseWriter, r *http.Request) error {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		return errors.New("Content-Type header is not application/json")
	}

	return nil
}
