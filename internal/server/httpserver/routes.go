package httpserver

import "net/http"

func (app *application) Routes() http.Handler {

	mux := http.NewServeMux()

	// Fileserver
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(app.cfg.StaticDir())})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Endpoints
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view/{id}", app.SnippetView)
	mux.HandleFunc("/snippet/create", app.SnippetCreate)

	// Middlewares
	// The request flow goes from bottom to up
	panicCatchedMux := app.recoverPanic(mux)
	loggedMux := app.requestLog(panicCatchedMux)
	secureHeadersAppliedMux := secureHeaders(loggedMux)

	//				//\\
	//| from here to ||
	return secureHeadersAppliedMux
}
