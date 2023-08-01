package handler

import "net/http"

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	return mux
}
