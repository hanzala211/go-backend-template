package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hanzala211/cms/internal/store"
	"go.uber.org/zap"
)

type application struct {
	config config
	logger *zap.SugaredLogger
	db     *sql.DB
	store  *store.Storage
}

type config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func (app *application) serve() *http.Server {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.checkHealth)
		r.Post("/health", app.writeCheckHealth)
	})
	return &http.Server{
		Addr:    app.config.addr,
		Handler: r,
	}
}
