package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Submission struct {
	ID             string
	QuizID         string
	StudentID      string
	Score          int
	TotalQuestions int
	SubmittedAt    time.Time
}

type SubmissionModel struct {
	DB *pgxpool.Pool
}

func (m *SubmissionModel) Insert(ctx context.Context, submission *Submission) error {
	query := `
		INSERT INTO submissions (quiz_id, student_id, score, total_questions)
		VALUES ($1, $2, $3, $4)
		RETURNING id, submitted_at
	`

	return m.DB.QueryRow(
		ctx, query,
		submission.QuizID, submission.StudentID,
		submission.Score, submission.TotalQuestions,
	).Scan(&submission.ID, &submission.SubmittedAt)
}

func (m *SubmissionModel) FindByStudent(ctx context.Context, studentID string) ([]Submission, error) {
	query := `
		SELECT id, quiz_id, student_id, score, total_questions, submitted_at
		FROM submissions
		WHERE student_id = $1
		ORDER BY submitted_at DESC
	`

	rows, err := m.DB.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submissions []Submission

	for rows.Next() {
		var s Submission
		if err := rows.Scan(
			&s.ID, &s.QuizID, &s.StudentID,
			&s.Score, &s.TotalQuestions, &s.SubmittedAt,
		); err != nil {
			return nil, err
		}
		submissions = append(submissions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return submissions, nil
}

func (m *SubmissionModel) FindByQuiz(ctx context.Context, quizID string) ([]Submission, error) {
	query := `
		SELECT id, quiz_id, student_id, score, total_questions, submitted_at
		FROM submissions
		WHERE quiz_id = $1
		ORDER BY submitted_at DESC
	`

	rows, err := m.DB.Query(ctx, query, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submissions []Submission

	for rows.Next() {
		var s Submission
		if err := rows.Scan(
			&s.ID, &s.QuizID, &s.StudentID,
			&s.Score, &s.TotalQuestions, &s.SubmittedAt,
		); err != nil {
			return nil, err
		}
		submissions = append(submissions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return submissions, nil
}
