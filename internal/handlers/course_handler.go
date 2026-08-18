package handlers

import (
	"net/http"

	"sabify/internal/services"
)

type CourseHandler struct {
	courseService *services.CourseService
}

func NewCourseHandler(courseService *services.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

// GET /courses
func (h *CourseHandler) ListCourses(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Courses page coming soon", http.StatusNotImplemented)
}

// GET /courses/{id}
func (h *CourseHandler) GetCourse(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Course page coming soon", http.StatusNotImplemented)
}

// POST /courses
func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Course creation coming soon", http.StatusNotImplemented)
}