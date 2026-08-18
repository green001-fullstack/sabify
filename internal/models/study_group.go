package models

import "time"

type StudyGroup struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CourseID  string    `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
}
