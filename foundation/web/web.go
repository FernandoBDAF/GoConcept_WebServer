// Package web contains a small web framework extension.
package web

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

// type MockEncoder struct {
// 	Data        any
// 	ContentType string
// 	Err         error
// }

// func (e *MockEncoder) Encode() (data []byte, contentType string, err error) {
// 	data, err = json.Marshal(e.Data)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return data, e.ContentType, e.Err
// }

// HandlerFunc represents a function that handles a http request within our own
// little mini framework.
type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

// Logger represents a function that will be called to add information
// to the logs.
type Logger func(ctx context.Context, msg string, args ...any)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	log    Logger
	tracer trace.Tracer
	mux    *http.ServeMux
	otmux   http.Handler
	mw []MidFunc
	// origins []string
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(log Logger, tracer trace.Tracer, mw ...MidFunc) *App {
	// Create an OpenTelemetry HTTP Handler which wraps our router. This will start
	// the initial span and annotate it with information about the request/trusted.
	//
	// This is configured to use the W3C TraceContext standard to set the remote
	// parent if a client request includes the appropriate headers.
	// https://w3c.github.io/trace-context/

	mux := http.NewServeMux()

	return &App{
		log:    log,
		tracer: tracer,
		mux:    mux,
		otmux:  otelhttp.NewHandler(mux, "request"),
		mw:     mw,
	}
}

// ServeHTTP implements the http.Handler interface. It's the entry point for
// all http traffic and allows the opentelemetry mux to run first to handle
// tracing. The opentelemetry mux then calls the application mux to handle
// application traffic. This was set up in the NewApp function.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.otmux.ServeHTTP(w, r)
}

// HandlerFunc sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) HandlerFunc(method string, group string, path string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	// if a.origins != nil {
	// 	handlerFunc = wrapMiddleware([]MidFunc{a.corsHandler}, handlerFunc)
	// }

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTracer(r.Context(), a.tracer)
		ctx = setWriter(ctx, w)

		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(w.Header()))

		resp := handlerFunc(ctx, r)

		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}
	}

	finalPath := path
	if group != "" {
		finalPath = "/" + group + path
	}
	finalPath = fmt.Sprintf("%s %s", method, finalPath)

	fmt.Println("finalPath", finalPath)

	a.mux.HandleFunc(finalPath, h)
}
