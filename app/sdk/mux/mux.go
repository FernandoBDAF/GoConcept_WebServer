// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/fernandobdaf/GoConcept_WebServer/app/domain/checkapp"
	"github.com/fernandobdaf/GoConcept_WebServer/app/sdk/mid"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal, log *logger.Logger) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log))

	checkapp.Routes(mux)

	return mux
}
