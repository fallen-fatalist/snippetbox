package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fallen-fatalist/snippetbox/internal/repository"
	"github.com/fallen-fatalist/snippetbox/internal/server/vo"
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

	data := app.NewTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, homePage, data)
}

func (app *application) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	snippet, err := app.Service().SnippetService().GetSnippetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNoRecord) {
			app.notFound(w)
		} else if service.IsServiceError(err) {
			app.clientError(w, http.StatusBadRequest)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	flash := app.sessionManager.PopString(r.Context(), "flash")

	data := app.NewTemplateData(r)
	data.Snippet = snippet
	data.Flash = flash

	app.render(w, r, http.StatusOK, viewPage, data)

}

var ErrMethodNotAllowed = errors.New("http method not allowed")

func (app *application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.Body = http.MaxBytesReader(w, r.Body, int64(service.SnippetFormBytesLimit))

		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		expires, err := strconv.Atoi(r.PostForm.Get("expires"))
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		form := vo.SnippetCreateForm{
			Title:   r.PostForm.Get("title"),
			Content: r.PostForm.Get("content"),
			Expires: expires,
		}

		snippetID, validator := app.Service().SnippetService().CreateSnippet(form.Title, form.Content, form.Expires)
		if err, exists := validator.FieldErrors["err"]; exists {
			app.serverError(w, r, err)
			return
		}

		if !validator.Valid() {
			data := app.NewTemplateData(r)
			form.FieldErrors = validator.FieldErrors
			data.Form = form

			app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
			return
		}

		app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", snippetID), http.StatusSeeOther)
	case http.MethodGet:
		data := app.NewTemplateData(r)
		data.Form = vo.SnippetCreateForm{}

		app.render(w, r, http.StatusOK, "create.html", data)
	default:
		w.Header().Set("Allow", http.MethodPost+""+http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

}
