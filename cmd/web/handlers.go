package main

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"sabify/internal/models"
	"sabify/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	data := app.newTemplateData(r)
	data.Title = "Home"

	app.render(w, http.StatusOK, "index.html", data)
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	err := app.models.Users.DB.Ping(ctx)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","database":"connected"}`))
}

func (app *application) showRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Register"

	app.render(w, http.StatusOK, "register.html", data)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	role := r.Form.Get("role")

	v := validator.New()
	v.CheckField(validator.NotBlank(name), "name", "This field cannot be blank")
	v.CheckField(validator.NotBlank(email), "email", "This field cannot be blank")
	v.CheckField(validator.MinChars(password, 8), "password", "This field must be at least 8 characters long")
	v.CheckField(validator.PermittedValue(role, "student", "teacher"), "role", "This field must be either student or teacher")

	if !v.Valid() {
		data := app.newTemplateData(r)
		data.Title = "Register"
		data.Form = map[string]string{
			"name":  name,
			"email": email,
		}
		data.FormErrors = v.GetFieldErrors()
		app.render(w, http.StatusUnprocessableEntity, "register.html", data)
		return
	}

	exists, err := app.models.Users.Exists(r.Context(), email)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if exists {
		data := app.newTemplateData(r)
		data.Title = "Register"
		data.Form = map[string]string{
			"name":  name,
			"email": email,
		}
		data.FormErrors = map[string]string{"email": "An account with this email already exists"}
		app.render(w, http.StatusUnprocessableEntity, "register.html", data)
		return
	}

	user := &models.User{
		Name:  strings.TrimSpace(name),
		Email: strings.ToLower(strings.TrimSpace(email)),
		Role:  role,
	}

	err = app.models.Users.Insert(r.Context(), user, password)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r.Context(), "flash", "Your registration was successful. Please log in.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) showLoginForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Login"

	app.render(w, http.StatusOK, "login.html", data)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	v := validator.New()
	v.CheckField(validator.NotBlank(email), "email", "This field cannot be blank")
	v.CheckField(validator.NotBlank(password), "password", "This field cannot be blank")

	if !v.Valid() {
		data := app.newTemplateData(r)
		data.Title = "Login"
		data.Form = map[string]string{"email": email}
		data.FormErrors = v.GetFieldErrors()
		app.render(w, http.StatusUnprocessableEntity, "login.html", data)
		return
	}

	user, err := app.models.Users.Authenticate(r.Context(), email, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			data := app.newTemplateData(r)
			data.Title = "Login"
			data.Form = map[string]string{"email": email}
			data.FormErrors = map[string]string{"email": "Invalid email or password"}
			app.render(w, http.StatusUnauthorized, "login.html", data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r.Context(), "authenticatedUserID", user.ID)
	app.session.Put(r.Context(), "userRole", user.Role)
	app.session.Put(r.Context(), "flash", "You have been logged in successfully!")

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r.Context(), "authenticatedUserID")
	app.session.Remove(r.Context(), "userRole")
	app.session.Put(r.Context(), "flash", "You have been logged out.")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := app.session.GetInt64(r.Context(), "authenticatedUserID")

		if userID == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := app.session.GetString(r.Context(), "userRole")

			if userRole != role {
				app.clientError(w, http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func (app *application) dashboard(w http.ResponseWriter, r *http.Request) {
	userRole := app.session.GetString(r.Context(), "userRole")

	switch userRole {
	case "teacher":
		http.Redirect(w, r, "/teacher/courses", http.StatusSeeOther)
	case "student":
		http.Redirect(w, r, "/student/courses", http.StatusSeeOther)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
