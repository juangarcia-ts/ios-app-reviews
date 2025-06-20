package service

import (
	"time"

	"ios-app-reviews-viewer.com/m/internal/repository"
)

type AppReviewsService struct {
	repository *repository.AppReviewsRepository
}

func NewAppReviewsService(repository *repository.AppReviewsRepository) *AppReviewsService {
	return &AppReviewsService{repository: repository}
}

func (s *AppReviewsService) FindAll(appId string, limit int, offset int) ([]repository.AppReview, error) {
	return s.repository.FindAll(appId, limit, offset)
}

func (s *AppReviewsService) Create(appId string, title string, content string, author string, rating int, submittedAt time.Time) (*repository.AppReview, error) {
	return s.repository.Create(appId, title, content, author, rating, submittedAt)
}
