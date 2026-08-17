package handlers

import (
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
	http.Error(w, "Register page coming soon", http.StatusNotImplemented)
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
