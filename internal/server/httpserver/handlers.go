package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/service"
)

const (
	DefaultHomeSnippetShown = 10
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.Service().SnippetService().LatestSnippets(DefaultHomeSnippetShown)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := NewTemplateData()
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, homePage, data)
}

func (app *application) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	snippet, err := app.Service().SnippetService().GetSnippetByID(int64(id))
	if err != nil {
		if errors.Is(err, repository.ErrNoRecord) {
			app.notFound(w)
		} else if service.IsServiceError(err) {
			app.clientError(w, err, http.StatusBadRequest)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := NewTemplateData()
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, viewPage, data)

}

var ErrMethodNotAllowed = errors.New("http method not allowed")

func (app *application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
	}
	switch r.Method {
	case http.MethodPost:
		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expires := 7

		snippetID, err := app.Service().SnippetService().CreateSnippet(title, content, expires)
		if err != nil {
			app.serverError(w, r, err)
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", snippetID), http.StatusSeeOther)
	case http.MethodGet:
		w.Write([]byte("Display the form for creating a new snippet..."))
	default:
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, ErrMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

}
