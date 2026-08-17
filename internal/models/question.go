package models

import "time"

type Question struct {
	ID            string    `json:"id"`
	QuizID        string    `json:"quiz_id"`
	QuestionText  string    `json:"question_text"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       string    `json:"option_d"`
	CorrectAnswer string    `json:"correct_answer"`
	CreatedAt     time.Time `json:"created_at"`
}
