package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"sabify/internal/models"
)

type SubmissionRepository struct {
	db *pgxpool.Pool
}

func NewSubmissionRepository(db *pgxpool.Pool) *SubmissionRepository {
	return &SubmissionRepository{db: db}
}

func (r *SubmissionRepository) Create(
	ctx context.Context,
	submission *models.Submission,
) error {

	query := `
		INSERT INTO submissions (
			quiz_id,
			student_id,
			score,
			total_questions
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, submitted_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		submission.QuizID,
		submission.StudentID,
		submission.Score,
		submission.TotalQuestions,
	).Scan(
		&submission.ID,
		&submission.SubmittedAt,
	)
}
