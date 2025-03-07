package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fernandobdaf/GoConcept_WebServer/foundation/logger"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
)

// Logger writes information about the request to the logs.
func Logger(log *logger.Logger) web.MidFunc {
	m := func(next web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) web.MockEncoder {
			now := time.Now()

			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			log.Info(ctx, "request started", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr)

			resp := next(ctx, w, r)
			// err := isError(resp)

			var statusCode = http.StatusOK
			// if err != nil {
			// 	statusCode = errs.Internal

				// var v *errs.Error
				// if errors.As(err, &v) {
				// 	statusCode = v.Code
				// }
			// }

			log.Info(ctx, "request completed", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr,
				"statuscode", statusCode, "since", time.Since(now).String())

			return resp
		}
		return h
	}
	return m
}
