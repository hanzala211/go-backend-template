package main

import (
	"net/http"
)

type CheckHealthPayload struct {
	Log string `json:"log" validate:"required"`
}

type WriteTokenHealhtPayload struct {
	UserID string `json:"user_id" validate:"required"`
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

func (app *application) writeTokenHealth(w http.ResponseWriter, r *http.Request) {
	var payload WriteTokenHealhtPayload
	err := app.DecodeStruct(w, r, &payload)
	if err != nil {
		return
	}
	token, err := app.jwtAuthenticator.GenerateToken(payload.UserID)
	if err != nil {
		app.writeJSONError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}
	app.writeJSON(w, http.StatusOK, map[string]any{
		"token": token,
	})
}

func (app *application) checkTokenHealth(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	if userId == "" {
		app.writeJSONError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	app.writeJSON(w, http.StatusOK, map[string]any{
		"user_id": userId,
	})
}
