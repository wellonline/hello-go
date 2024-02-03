package main

import (
	"hello-go/handler"
	"log"
	"log/slog"
	"net/http"
	"time"
)

var build = "develop"

func main() {
	logger := slog.New(slog.Default().Handler())
	logger.Info("start", "version", build)

	h := handler.New()
	srv := http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           h,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
