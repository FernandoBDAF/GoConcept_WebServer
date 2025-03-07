package mid

import (
	"context"
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// Logger writes information about the request to the logs.
func Logger(log *logger.Logger) web.MidFunc {
	m := func(next web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) web.MockEncoder {

			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			encoder := next(ctx, w, r)

			log.Info(ctx, "request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			return encoder
		}
		return h
	}
	return m
}
