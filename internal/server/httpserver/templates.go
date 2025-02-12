package httpserver

import (
	"html/template"
	"path/filepath"
)

type templateData struct {
	Snippets []viewSnippet
	Snippet  viewSnippet
}

type viewSnippet struct {
	ID      int
	Title   string
	Content string
	Created string
	Expires string
}

// Page names
const (
	homePage = "home.html"
	viewPage = "view.html"
)

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
