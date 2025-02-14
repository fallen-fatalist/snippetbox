package httpserver

import (
	"net/http"

	"github.com/fallen-fatalist/snippetbox/internal/server/vo"
)

func (app *application) UserSignup(w http.ResponseWriter, r *http.Request) {
	if app.isAuthenticated(r) {
		app.sessionManager.Put(r.Context(), "flash", "User already logged in!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	switch r.Method {
	case http.MethodPost:
		r.Body = http.MaxBytesReader(w, r.Body, 4096)

		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		form := vo.UserSignupForm{
			Name:     r.PostForm.Get("name"),
			Password: r.PostForm.Get("password"),
			Email:    r.PostForm.Get("email"),
		}

		_, validator := app.Service().UserService().CreateUser(form.Name, form.Email, form.Password)
		if err, exists := validator.FieldErrors["err"]; exists {
			app.serverError(w, r, err)
			return
		}

		if !validator.Valid() {
			data := app.NewTemplateData(r)
			form.FieldErrors = validator.FieldErrors
			data.Form = form

			app.render(w, r, http.StatusUnprocessableEntity, "signup.html", data)
			return
		}

		app.sessionManager.Put(r.Context(), "flash", "User successfully created!")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)

	case http.MethodGet:
		data := app.NewTemplateData(r)
		data.Form = vo.UserSignupForm{}
		app.render(w, r, http.StatusOK, "signup.html", data)
		return
	default:
		w.Header().Set("Allow", http.MethodPost+" "+http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return

	}
}

func (app *application) UserLogin(w http.ResponseWriter, r *http.Request) {
	if app.isAuthenticated(r) {
		app.sessionManager.Put(r.Context(), "flash", "User already logged in!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	switch r.Method {
	case http.MethodGet:
		data := app.NewTemplateData(r)
		data.Form = vo.UserLoginForm{}
		app.render(w, r, http.StatusOK, "login.html", data)
		return
	case http.MethodPost:
		r.Body = http.MaxBytesReader(w, r.Body, 4096)

		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}

		form := vo.UserLoginForm{
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}

		userID, validator := app.Service().UserService().Authenticate(form.Email, form.Password)
		if !validator.Valid() {
			data := app.NewTemplateData(r)
			form.FieldErrors = validator.FieldErrors
			form.NonFieldErrors = validator.NonFieldErrors
			data.Form = form
			app.render(w, r, http.StatusUnprocessableEntity, "login.html", data)
		}

		err = app.sessionManager.RenewToken(r.Context())
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		app.sessionManager.Put(r.Context(), "authenticatedUserID", userID)
		app.sessionManager.Put(r.Context(), "flash", "User successfully logined!")
		http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
	default:
		w.Header().Set("Allow", http.MethodPost+" "+http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return

	}
}

func (app *application) UserLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")
	app.sessionManager.Put(r.Context(), "flash", "You've been successfully logged out!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
}
