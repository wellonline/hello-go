package web

import (
	"context"
	"log/slog"
	"net/http"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	mux *http.ServeMux
}

func NewApp() *App {
	mux := http.NewServeMux()
	return &App{
		mux: mux,
	}
}

func (a *App) HandleNoMiddleware(method string, path string, handler Handler) {
	h := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			slog.Error(err.Error())
		}
	}
	a.mux.HandleFunc(path, h)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
