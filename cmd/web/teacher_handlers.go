package main

import (
	"net/http"
	"strings"

	"sabify/internal/models"
	"sabify/internal/validator"
)

func (app *application) teacherCourses(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "My Courses"

	app.render(w, http.StatusOK, "teacher/courses.html", data)
}

func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.Form.Get("title")
	description := r.Form.Get("description")

	v := validator.New()
	v.CheckField(validator.NotBlank(title), "title", "This field cannot be blank")
	v.CheckField(validator.MaxChars(title, 255), "title", "This field must not be more than 255 characters long")

	if !v.Valid() {
		data := app.newTemplateData(r)
		data.Title = "Create Course"
		data.Form = map[string]string{
			"title":       title,
			"description": description,
		}
		data.FormErrors = v.GetFieldErrors()
		app.render(w, http.StatusUnprocessableEntity, "teacher/create-course.html", data)
		return
	}

	teacherID := app.session.GetString(r.Context(), "authenticatedUserID")

	course := &models.Course{
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		TeacherID:   teacherID,
	}

	err = app.models.Courses.Insert(r.Context(), course)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r.Context(), "flash", "Course created successfully!")
	http.Redirect(w, r, "/teacher/courses", http.StatusSeeOther)
}

func (app *application) teacherCourseDetail(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Course Detail"

	app.render(w, http.StatusOK, "teacher/course-detail.html", data)
}

func (app *application) teacherQuizzes(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "My Quizzes"

	app.render(w, http.StatusOK, "teacher/quizzes.html", data)
}

func (app *application) createQuiz(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Create Quiz"

	app.render(w, http.StatusOK, "teacher/create-quiz.html", data)
}

func (app *application) teacherSubmissions(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Title = "Submissions"

	app.render(w, http.StatusOK, "teacher/submissions.html", data)
}
