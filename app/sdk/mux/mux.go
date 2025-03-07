// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/fernandobdaf/GoConcept_WebServer/app/domain/checkapp"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/mid"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	// DB     *sqlx.DB
	Tracer trace.Tracer
	// SalesConfig
	// AuthConfig
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(config Config, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(config.Tracer, shutdown, mid.Otel(config.Tracer), mid.Logger(config.Log))

	checkapp.Routes(mux)

	return mux
}
