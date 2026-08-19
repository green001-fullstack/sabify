package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Quiz struct {
	ID          string
	CourseID    string
	Title       string
	Description string
	CreatedAt   time.Time
}

type QuizModel struct {
	DB *pgxpool.Pool
}

func (m *QuizModel) Insert(ctx context.Context, quiz *Quiz) error {
	query := `
		INSERT INTO quizzes (course_id, title, description)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return m.DB.QueryRow(
		ctx, query,
		quiz.CourseID, quiz.Title, quiz.Description,
	).Scan(&quiz.ID, &quiz.CreatedAt)
}

func (m *QuizModel) FindByID(ctx context.Context, id string) (*Quiz, error) {
	var quiz Quiz

	query := `
		SELECT id, course_id, title, description, created_at
		FROM quizzes
		WHERE id = $1
	`

	err := m.DB.QueryRow(ctx, query, id).Scan(
		&quiz.ID, &quiz.CourseID, &quiz.Title,
		&quiz.Description, &quiz.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (m *QuizModel) FindByCourse(ctx context.Context, courseID string) ([]Quiz, error) {
	query := `
		SELECT id, course_id, title, description, created_at
		FROM quizzes
		WHERE course_id = $1
		ORDER BY created_at DESC
	`

	rows, err := m.DB.Query(ctx, query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []Quiz

	for rows.Next() {
		var quiz Quiz
		if err := rows.Scan(
			&quiz.ID, &quiz.CourseID, &quiz.Title,
			&quiz.Description, &quiz.CreatedAt,
		); err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (m *QuizModel) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM quizzes WHERE id = $1`

	result, err := m.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrNoRecord
	}

	return nil
}
