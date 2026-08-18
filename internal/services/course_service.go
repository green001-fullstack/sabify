package services

import (
	"context"
	"errors"
	"strings"

	"sabify/internal/models"
	"sabify/internal/repositories"
)

type CourseService struct {
	courseRepo *repositories.CourseRepository
}

func NewCourseService(
	courseRepo *repositories.CourseRepository,
) *CourseService {
	return &CourseService{
		courseRepo: courseRepo,
	}
}

func (s *CourseService) CreateCourse(
	ctx context.Context,
	title string,
	description string,
	teacherID string,
) (*models.Course, error) {

	title = strings.TrimSpace(title)

	if title == "" {
		return nil, errors.New("course title is required")
	}

	if teacherID == "" {
		return nil, errors.New("teacher ID is required")
	}

	course := &models.Course{
		Title:       title,
		Description: strings.TrimSpace(description),
		TeacherID:   teacherID,
	}

	if err := s.courseRepo.Create(ctx, course); err != nil {
		return nil, err
	}

	return course, nil
}

func (s *CourseService) GetCourse(
	ctx context.Context,
	id string,
) (*models.Course, error) {

	return s.courseRepo.FindByID(ctx, id)
}

func (s *CourseService) GetTeacherCourses(
	ctx context.Context,
	teacherID string,
) ([]models.Course, error) {

	return s.courseRepo.FindByTeacher(ctx, teacherID)
}
