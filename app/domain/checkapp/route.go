package checkapp

import "github.com/fernandobdaf/GoConcept_WebServer/foundation/web"

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

	app.HandlerFunc("/readiness", api.readiness)
	app.HandlerFunc("/liveness", api.liveness)
}
