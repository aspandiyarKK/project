package main

import (
	"errors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
	"os"
	"os/signal"
	"project/pkg/logger"
	"project/pkg/repository"
	"syscall"

	"project/internal/rest"
)

var version = "0.0.1"

var pgDSN = os.Getenv("PG_DSN")

func main() {
	log := logger.NewLogger()
	log.Infof("Starting server on port 8080")
	pg, err := repository.NewPG(log, pgDSN)
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}
	r := rest.NewRouter(log, version, pg)
	go func() {
		if err = r.Run("localhost:8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("Error starting server: %v", err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	<-sigCh
	log.Info("Shutting down")
}
