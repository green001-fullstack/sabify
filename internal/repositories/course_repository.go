package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"sabify/internal/models"
)

type CourseRepository struct {
	db *pgxpool.Pool
}

func NewCourseRepository(db *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(
	ctx context.Context,
	course *models.Course,
) error {

	query := `
		INSERT INTO courses (
			title,
			description,
			teacher_id
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		ctx,
		query,
		course.Title,
		course.Description,
		course.TeacherID,
	).Scan(
		&course.ID,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
}

func (r *CourseRepository) FindByID(
	ctx context.Context,
	id string,
) (*models.Course, error) {

	query := `
		SELECT
			id,
			title,
			description,
			teacher_id,
			created_at,
			updated_at
		FROM courses
		WHERE id = $1
	`

	var course models.Course

	err := r.db.QueryRow(ctx, query, id).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.TeacherID,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *CourseRepository) FindByTeacher(
	ctx context.Context,
	teacherID string,
) ([]models.Course, error) {

	query := `
		SELECT
			id,
			title,
			description,
			teacher_id,
			created_at,
			updated_at
		FROM courses
		WHERE teacher_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course

	for rows.Next() {
		var course models.Course

		if err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.TeacherID,
			&course.CreatedAt,
			&course.UpdatedAt,
		); err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, rows.Err()
}
