// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"context"
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/auth"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/authclient"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/mid"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

// StaticSite represents a static site to run.
// type StaticSite struct {
// 	react      bool
// 	static     embed.FS
// 	staticDir  string
// 	staticPath string
// }

// Options represent optional parameters.
type Options struct {
	corsOrigin []string
	// sites      []StaticSite
}

// WithCORS provides configuration options for CORS.
func WithCORS(origins []string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origins
	}
}

type SalesConfig struct {
	AuthClient *authclient.Client
}

// AuthConfig contains auth service specific config.
type AuthConfig struct {
	Auth *auth.Auth
}

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	// DB     *sqlx.DB
	Tracer trace.Tracer
	SalesConfig
	AuthConfig
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg Config)
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config, routeAdder RouteAdder, options ...func(opts *Options)) http.Handler {
	logger := func(ctx context.Context, msg string, args ...any) {
		cfg.Log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		logger,
		cfg.Tracer,
		mid.Otel(cfg.Tracer),
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	)

	var opts Options
	for _, option := range options {
		option(&opts)
	}

	if len(opts.corsOrigin) > 0 {
		app.EnableCORS(opts.corsOrigin)
	}

	routeAdder.Add(app, cfg)

	// for _, site := range opts.sites {
	// 	if site.react {
	// 		app.FileServerReact(site.static, site.staticDir, site.staticPath)
	// 	} else {
	// 		app.FileServer(site.static, site.staticDir, site.staticPath)
	// 	}
	// }

	return app
}
