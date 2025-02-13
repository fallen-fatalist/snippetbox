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

	data := NewTemplateData(r)
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
			app.clientError(w, http.StatusBadRequest)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := NewTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, viewPage, data)

}

var ErrMethodNotAllowed = errors.New("http method not allowed")

func (app *application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.Body = http.MaxBytesReader(w, r.Body, 4096)

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

		snippetID, errs := app.Service().SnippetService().CreateSnippet(form.Title, form.Content, form.Expires)
		if err, ok := errs["err"]; ok {
			app.serverError(w, r, err)
			return
		}

		if len(errs) > 0 {
			data := NewTemplateData(r)
			form.FieldErrors = errs
			data.Form = form

			app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", snippetID), http.StatusSeeOther)
	case http.MethodGet:
		data := NewTemplateData(r)
		data.Form = vo.SnippetCreateForm{}

		app.render(w, r, http.StatusOK, "create.html", data)
	default:
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

}
