package mid

import (
	"context"
	"net/http"

	"github.com/fernandobdaf/GoConcept_WebServer/foundation/otel"
	"github.com/fernandobdaf/GoConcept_WebServer/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

// Otel starts the otel tracing and stores the trace id in the context.
func Otel(tracer trace.Tracer) web.MidFunc {
	m := func(next web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) web.MockEncoder {
			ctx = otel.InjectTracing(ctx, tracer)

			return next(ctx, w, r)
		}

		return h
	}

	return m
}
