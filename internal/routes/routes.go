package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"sabify/internal/handlers"
)

func SetupRouter(
	homeHandler *handlers.HomeHandler,
	authHandler *handlers.AuthHandler,
	courseHandler *handlers.CourseHandler,
) *chi.Mux {

	r := chi.NewRouter()

	r.Handle(
		"/static/*",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	r.Get("/", homeHandler.Home)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("SABIFY is running"))
	})

	r.Get("/register", authHandler.ShowRegister)
	r.Post("/register", authHandler.Register)

	r.Get("/login", authHandler.ShowLogin)
	r.Post("/login", authHandler.Login)

	r.Get("/courses", courseHandler.ListCourses)
	r.Get("/courses/{id}", courseHandler.GetCourse)
	r.Post("/courses", courseHandler.CreateCourse)

	return r
}
