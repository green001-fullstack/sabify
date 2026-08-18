package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Material struct {
	ID          string
	CourseID    string
	Title       string
	Description string
	FileURL     string
	CreatedAt   time.Time
}

type MaterialModel struct {
	DB *pgxpool.Pool
}

func (m *MaterialModel) Insert(ctx context.Context, material *Material) error {
	query := `
		INSERT INTO materials (course_id, title, description, file_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	return m.DB.QueryRow(
		ctx, query,
		material.CourseID, material.Title, material.Description, material.FileURL,
	).Scan(&material.ID, &material.CreatedAt)
}

func (m *MaterialModel) FindByCourse(ctx context.Context, courseID string) ([]Material, error) {
	query := `
		SELECT id, course_id, title, description, file_url, created_at
		FROM materials
		WHERE course_id = $1
		ORDER BY created_at DESC
	`

	rows, err := m.DB.Query(ctx, query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []Material

	for rows.Next() {
		var mat Material
		if err := rows.Scan(
			&mat.ID, &mat.CourseID, &mat.Title,
			&mat.Description, &mat.FileURL, &mat.CreatedAt,
		); err != nil {
			return nil, err
		}
		materials = append(materials, mat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return materials, nil
}
