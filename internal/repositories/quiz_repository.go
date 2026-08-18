package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"sabify/internal/models"
)

type QuizRepository struct {
	db *pgxpool.Pool
}

func NewQuizRepository(db *pgxpool.Pool) *QuizRepository {
	return &QuizRepository{db: db}
}

func (r *QuizRepository) Create(
	ctx context.Context,
	quiz *models.Quiz,
) error {

	query := `
		INSERT INTO quizzes (
			course_id,
			title,
			description
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		quiz.CourseID,
		quiz.Title,
		quiz.Description,
	).Scan(
		&quiz.ID,
		&quiz.CreatedAt,
	)
}

func (r *QuizRepository) AddQuestion(
	ctx context.Context,
	question *models.Question,
) error {

	query := `
		INSERT INTO questions (
			quiz_id,
			question_text,
			option_a,
			option_b,
			option_c,
			option_d,
			correct_answer
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		question.QuizID,
		question.QuestionText,
		question.OptionA,
		question.OptionB,
		question.OptionC,
		question.OptionD,
		question.CorrectAnswer,
	).Scan(
		&question.ID,
		&question.CreatedAt,
	)
}

func (r *QuizRepository) FindQuestions(
	ctx context.Context,
	quizID string,
) ([]models.Question, error) {

	query := `
		SELECT
			id,
			quiz_id,
			question_text,
			option_a,
			option_b,
			option_c,
			option_d,
			correct_answer,
			created_at
		FROM questions
		WHERE quiz_id = $1
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(ctx, query, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question

	for rows.Next() {
		var question models.Question

		if err := rows.Scan(
			&question.ID,
			&question.QuizID,
			&question.QuestionText,
			&question.OptionA,
			&question.OptionB,
			&question.OptionC,
			&question.OptionD,
			&question.CorrectAnswer,
			&question.CreatedAt,
		); err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, rows.Err()
}
