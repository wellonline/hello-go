package health

import (
	"context"
	"net/http"
)

type handlers struct{}

func new() *handlers {
	return &handlers{}
}

func (h *handlers) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}

func (h *handlers) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) {

}
