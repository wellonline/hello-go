package health

import (
	"hello-go/foundation/web"
	"net/http"
)

// Routes adds specific routes for this group
func Routes(app *web.App) {
	hdl := new()
	app.HandleNoMiddleware(http.MethodGet, "/readiness", hdl.readiness)
	app.HandleNoMiddleware(http.MethodGet, "/liveness", hdl.liveness)
}
