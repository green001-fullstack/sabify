package handlers

import (
	"html/template"
	"net/http"
)

type HomeHandler struct {
	template *template.Template
}

func NewHomeHandler() *HomeHandler {
	tmpl := template.Must(
		template.ParseFiles("templates/layouts/base.html", "templates/layouts/public.html", "templates/pages/home/index.html"),
	)

	return &HomeHandler{
		template: tmpl,
	}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := h.template.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Unable to load homepage", http.StatusInternalServerError)
	}
}
