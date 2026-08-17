package main

import (
	"log"
	"net/http"

	"sabify/internal/database"
	"sabify/internal/handlers"
	"sabify/internal/repositories"
	"sabify/internal/routes"
	"sabify/internal/services"
)

func main() {

	// -------------------------
	// Database
	// -------------------------

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	log.Println(" Successfully Connected to PostgreSQL")

	// -------------------------
	// Repositories
	// -------------------------

	userRepository := repositories.NewUserRepository(db)
	courseRepository := repositories.NewCourseRepository(db)

	// -------------------------
	// Services
	// -------------------------

	authService := services.NewAuthService(userRepository)
	courseService := services.NewCourseService(courseRepository)

	// -------------------------
	// Handlers
	// -------------------------

	authHandler := handlers.NewAuthHandler(authService)
	courseHandler := handlers.NewCourseHandler(courseService)
	homeHandler := handlers.NewHomeHandler()

	// -------------------------
	// Router
	// -------------------------

	router := routes.SetupRouter(
		homeHandler,
		authHandler,
		courseHandler,
	)

	// -------------------------
	// Server
	// -------------------------

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println(" SABIFY running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed:", err)
	}
}
