package main

import (
	"errors"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	if err := app.toJson(w, http.StatusOK, &payload); err != nil {
		app.errorJson(w, errors.New("internal server error"), http.StatusInternalServerError)
	}
}
