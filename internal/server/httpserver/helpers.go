package httpserver

import (
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(
		err.Error(),
		slog.String("method", r.Method),
		slog.String("uri", r.URL.RequestURI()),
		slog.String("trace", string(debug.Stack())),
	)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

var ErrPageNotExist = errors.New("page you want to get does not exist")

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, ErrPageNotExist, http.StatusNotFound)
}
