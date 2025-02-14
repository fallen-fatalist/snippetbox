package httpserver

import (
	"net/http"

	"github.com/fallen-fatalist/snippetbox/internal/server/vo"
)

func (app *application) UserSignup(w http.ResponseWriter, r *http.Request) {
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
	switch r.Method {
	case http.MethodPost:
	// 	r.Body = http.MaxBytesReader(w, r.Body, 4096)

	// 	err := r.ParseForm()
	// 	if err != nil {
	// 		app.clientError(w, http.StatusBadRequest)
	// 		return
	// 	}

	// 	form := vo.UserSignupForm{
	// 		Name:     r.PostForm.Get("name"),
	// 		Password: r.PostForm.Get("password"),
	// 		Email:    r.PostForm.Get("email"),
	// 	}

	// 	validator := app.Service().UserService().CreateUser(form.Name, form.Email, form.Password)
	// 	if err, exists := validator.FieldErrors["err"]; exists {
	// 		app.serverError(w, r, err)
	// 		return
	// 	}

	// 	if !validator.Valid() {
	// 		data := app.NewTemplateData(r)
	// 		form.FieldErrors = validator.FieldErrors
	// 		data.Form = form

	// 		app.render(w, r, http.StatusUnprocessableEntity, "create.html", data)
	// 		return
	// 	}

	// 	app.sessionManager.Put(r.Context(), "flash", "User successfully created!s")

	// case http.MethodGet:
	// 	data := app.NewTemplateData(r)
	// 	data.Form = vo.UserSignupForm{}
	// 	app.render(w, r, http.StatusOK, "Signup.html", data)
	// 	return
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
}
