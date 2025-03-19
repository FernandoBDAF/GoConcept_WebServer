package checkapp

import (
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/auth"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/authclient"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/mid"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	// DB    *sqlx.DB
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	// api := newApp(cfg.Build, cfg.Log, cfg.DB)

	// api := newApp(cfg.Build, cfg.Log, cfg.DB)
	authen := mid.Authenticate(cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newApp(cfg.Build, cfg.Log)

	app.HandlerFuncNoMid(http.MethodGet, version, "/readiness", api.readiness)
	app.HandlerFuncNoMid(http.MethodGet, version, "/liveness", api.liveness)
	app.HandlerFunc(http.MethodGet, version, "/test-error", api.testError)
	app.HandlerFunc(http.MethodGet, version, "/test-panic", api.testPanic)
	app.HandlerFunc(http.MethodGet, version, "/test-auth", api.testAuth, authen, ruleAdmin)
}
