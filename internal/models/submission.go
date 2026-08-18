package models

import "time"

type Submission struct {
	ID             string    `json:"id"`
	QuizID         string    `json:"quiz_id"`
	StudentID      string    `json:"student_id"`
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	SubmittedAt    time.Time `json:"submitted_at"`
}
