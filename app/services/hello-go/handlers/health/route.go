package health

import (
	"hello-go/foundation/web"
	"net/http"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
}

// Routes adds specific routes for this group
func Routes(app *web.App, cfg Config) {
	hdl := new(cfg.Build)
	app.HandleNoMiddleware(http.MethodGet, "/readiness", hdl.readiness)
	app.HandleNoMiddleware(http.MethodGet, "/liveness", hdl.liveness)
}
