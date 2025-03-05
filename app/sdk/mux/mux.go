// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/app/domain/checkapp"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI() http.Handler {
	mux := http.NewServeMux()

	checkapp.Routes(mux)

	return mux
}
