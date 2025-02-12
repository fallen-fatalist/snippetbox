package httpserver

import "net/http"

func (app *application) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	// Fileserver
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(app.cfg.StaticDir())})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Endpoints
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.SnippetView)
	mux.HandleFunc("/snippet/create", app.SnippetCreate)

	return mux
}
