package handler

import (
	"errors"
	"net/http"
)

func Validation(w http.ResponseWriter, r *http.Request) error {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		return errors.New("Content-Type должен быть application/json")
	}

	return nil
}
