package health

import (
	"context"
	"hello-go/foundation/web"
	"net/http"
	"os"
)

type handlers struct {
	build string
}

func new(build string) *handlers {
	return &handlers{
		build,
	}
}

func (h *handlers) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *handlers) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status string `json:"up"`
		Build  string `json:"build"`
		Host   string `json:"host"`
	}{
		"up",
		h.build,
		host,
	}

	return web.Respond(ctx, w, r, data, http.StatusOK)
}
