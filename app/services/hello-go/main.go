package main

import (
	"context"
	"fmt"
	"hello-go/app/services/hello-go/handlers/health"
	"hello-go/foundation/web"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var build = "develop"

func main() {
	log := slog.New(slog.Default().Handler())
	log.Info("start", "version", build)

	ctx := context.Background()
	if err := run(ctx, log); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context, log *slog.Logger) error {
	log.Info("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	cfg := struct {
		BuildVersion string
	}{
		build,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	app := web.NewApp()
	health.Routes(app, health.Config{cfg.BuildVersion})
	api := http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		Handler:           app,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Info("startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		log.Info("shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info("shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			closeErr := api.Close()
			return fmt.Errorf("could not stop server gracefully: %w, %w", err, closeErr)
		}
	}

	return nil
}
