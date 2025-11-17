package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hanzala211/cms/internal/db"
	"github.com/hanzala211/cms/internal/env"
	"github.com/hanzala211/cms/internal/store"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}
	cfg := config{
		addr: ":" + env.GetEnv("ADDR", "4000"),
		dbConfig: dbConfig{
			host:     env.GetEnv("DB_HOST", "localhost"),
			port:     env.GetEnv("DB_PORT", "5432"),
			user:     env.GetEnv("DB_USER", "postgres"),
			password: env.GetEnv("DB_PASSWORD", "postgres"),
			dbname:   env.GetEnv("DB_NAME", "postgres"),
		},
	}

	loggerProd, _ := zap.NewDevelopment()
	defer loggerProd.Sync()
	db := db.New(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.dbConfig.user,
		cfg.dbConfig.password,
		cfg.dbConfig.host,
		cfg.dbConfig.port,
		cfg.dbConfig.dbname,
	))
	logger := loggerProd.Sugar()
	userStore := store.NewUserStruct(db)
	storage := store.NewStorage(userStore)
	app := application{
		config: cfg,
		logger: logger,
		db:     db,
		store:  storage,	
	}
	srv := app.serve()
	serverErr := make(chan error, 1) // one buffer because multiple errors can cause the channel to miss the real err
	go func() {
		logger.Infow("server started", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			serverErr <- err
		}
		close(serverErr)
	}()
	stopSign := make(chan os.Signal, 1)
	signal.Notify(stopSign, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-serverErr:
		logger.Fatalw("server error", zap.Error(err))
	case <-stopSign:
		logger.Info("shutdown signal received, gracefully shutting down server")
	}
	shutDownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutDownCtx); err != nil {
		logger.Fatalw("server shutdown error", zap.Error(err))
	}
	logger.Info("server stopped gracefully")
}
