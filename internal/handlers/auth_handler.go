package handlers

import (
	"html/template"
	"net/http"

	"sabify/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// GET /register
func (h *AuthHandler) ShowRegister(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "Register page coming soon", http.StatusNotImplemented)
	tmpl := template.Must(
		template.ParseFiles("templates/auth/register.html", "templates/auth/login.html"),
	)
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, "Unable to load register page", http.StatusInternalServerError)
	}
}

// POST /register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// need to connect this to the register template next.
	http.Error(w, "Registration coming soon", http.StatusNotImplemented)
}

// GET /login
func (h *AuthHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Login page coming soon", http.StatusNotImplemented)
}

// POST /login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// need to  connect this to the login template next.
	http.Error(w, "Login coming soon", http.StatusNotImplemented)
}
