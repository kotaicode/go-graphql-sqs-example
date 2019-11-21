package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Info().Msg("Shutting down")
	os.Exit(0)
}

// MustGet will return the env or panic if it is not present
func MustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic().Msgf("ENV missing, key: " + k)
	}
	return v
}
