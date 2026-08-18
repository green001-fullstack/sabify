package main

import (
	"net/http"
)

func (app *application) studentCourses(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "My Courses"

	app.render(w, http.StatusOK, "student-courses.html", data)
}

func (app *application) studentCourseDetail(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Course Detail"

	app.render(w, http.StatusOK, "student-course-detail.html", data)
}

func (app *application) studentQuizzes(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Available Quizzes"

	app.render(w, http.StatusOK, "student-quizzes.html", data)
}

func (app *application) submitQuiz(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Quiz Results"

	app.render(w, http.StatusOK, "student-results.html", data)
}

func (app *application) studentResults(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "My Results"

	app.render(w, http.StatusOK, "student-results.html", data)
}

func (app *application) studentStudyGroups(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Study Groups"

	app.render(w, http.StatusOK, "student-study-groups.html", data)
}
