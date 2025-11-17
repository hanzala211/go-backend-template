package main

import (
	"net/http"
)

type CheckHealthPayload struct {
	Log string `json:"log" validate:"required"`
}

func (app *application) checkHealth(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Server is running")
	app.writeJSON(w, http.StatusOK, "OK")
}

func (app *application) writeCheckHealth(w http.ResponseWriter, r *http.Request) {
	var payload CheckHealthPayload
	err := app.DecodeStruct(w, r, &payload)
	if err != nil {
		return
	}
	app.logger.Infow("Server is running", "payload", payload)
	app.writeJSON(w, http.StatusOK, map[string]any{
		"log": payload.Log,
	})
}
