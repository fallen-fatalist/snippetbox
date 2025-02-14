package httpserver

import "net/http"

func (app *application) Routes() http.Handler {

	mux := http.NewServeMux()

	// Fileserver
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(app.cfg.StaticDir())})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Snippet Endpoints
	mux.Handle("/", app.sessionManager.LoadAndSave(http.HandlerFunc(app.Home)))
	mux.Handle("/snippet/view/{id}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.SnippetView)))
	mux.Handle("/snippet/create", app.sessionManager.LoadAndSave(http.HandlerFunc(app.SnippetCreate)))

	// User Endpoints
	mux.Handle("/user/signup", app.sessionManager.LoadAndSave(http.HandlerFunc(app.UserSignup)))
	mux.Handle("/user/login", app.sessionManager.LoadAndSave(http.HandlerFunc(app.UserSignup)))
	mux.Handle("/user/logout", app.sessionManager.LoadAndSave(http.HandlerFunc(app.UserSignup)))

	// Middlewares
	// The request flow goes from bottom to up
	panicCatchedMux := app.recoverPanic(mux)
	loggedMux := app.requestLog(panicCatchedMux)
	secureHeadersAppliedMux := secureHeaders(loggedMux)

	//				//\\
	//| from here to ||
	return secureHeadersAppliedMux
}
