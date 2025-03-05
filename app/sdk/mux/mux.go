// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"encoding/json"
	"net/http"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI() http.Handler {
	mux := http.NewServeMux()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
	})

	mux.Handle("GET /test", h)


	return mux
}
