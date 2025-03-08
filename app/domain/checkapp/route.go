package checkapp

import (
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// "github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
// "github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
// "github.com/jmoiron/sqlx"

// Config contains all the mandatory systems required by handlers.
type Config struct {
	// Build string
	// Log   *logger.Logger
	// DB    *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	// const version = "v1"

	// api := newApp(cfg.Build, cfg.Log, cfg.DB)
	api := newApp()

	app.HandlerFunc(http.MethodGet, "", "/readiness", api.readiness)
	app.HandlerFunc(http.MethodGet, "", "/liveness", api.liveness)
	app.HandlerFunc(http.MethodGet, "", "/test-error", api.testError)
	app.HandlerFunc(http.MethodGet, "", "/test-panic", api.testPanic)
}
