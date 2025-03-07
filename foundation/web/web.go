// Package web contains a small web framework extension.
package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type MockEncoder struct {
	Data        any
	ContentType string
	Err         error
}

func (e *MockEncoder) Encode() (data []byte, contentType string, err error) {
	data, err = json.Marshal(e.Data)
	if err != nil {
		return nil, "", err
	}

	return data, e.ContentType, e.Err
}

// HandlerFunc represents a function that handles a http request within our own
// little mini framework.
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) MockEncoder

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	shutdown chan os.Signal
	mw       []MidFunc
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...MidFunc) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		shutdown: shutdown,
		mw:       mw,
	}
}

// HandlerFunc sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) HandlerFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {

		// PUT ANY CODE WE WANT TO RUN BEFORE THE HANDLER HERE

		if encode := handlerFunc(r.Context(), w, r); encode.Err != nil {
			// ERROR HANDLING HERE
			fmt.Println(encode.Err)
			return
		}

		// PUT ANY CODE WE WANT TO RUN AFTER THE HANDLER HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
