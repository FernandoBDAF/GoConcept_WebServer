package authapp

import (
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/auth"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/mid"
	// "github.com/fernandobdaf/GoConcept_WebServer/business/domain/userbus"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	// UserBus *userbus.Business
	Auth    *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	bearer := mid.Bearer(cfg.Auth)
	// basic := mid.Basic(cfg.Auth, cfg.UserBus)

	api := newApp(cfg.Auth)

	// app.HandlerFunc(http.MethodGet, version, "/auth/token/{kid}", api.token, basic)
	app.HandlerFunc(http.MethodGet, version, "/auth/authenticate", api.authenticate, bearer)
	app.HandlerFunc(http.MethodPost, version, "/auth/authorize", api.authorize)
}
