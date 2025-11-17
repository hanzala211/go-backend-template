package main

import (
	"context"
	"net/http"
)

func (app *application) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			app.writeJSONError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		userId, err := app.jwtAuthenticator.ValidateToken(token)
		if err != nil {
			app.writeJSONError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
