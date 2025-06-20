package service

import (
	"fmt"
	"time"

	"ios-app-reviews-viewer.com/m/internal/client"
	"ios-app-reviews-viewer.com/m/internal/repository"
)

type AppReviewsService struct {
	repository     *repository.AppReviewsRepository
	appStoreClient *client.AppStoreClient
}

func NewAppReviewsService(repository *repository.AppReviewsRepository, appStoreClient *client.AppStoreClient) *AppReviewsService {
	return &AppReviewsService{
		repository:     repository,
		appStoreClient: appStoreClient,
	}
}

func (s *AppReviewsService) FindAllWithCount(appId string, limit int, offset int, startsAt time.Time, endsAt time.Time) ([]repository.AppReview, int, error) {
	count, err := s.repository.Count(appId, startsAt, endsAt)

	if err != nil {
		return nil, 0, err
	}

	reviews, err := s.repository.FindAll(appId, limit, offset, startsAt, endsAt)

	if err != nil {
		return nil, count, err
	}

	return reviews, count, nil
}

func (s *AppReviewsService) Create(appId string, appStoreReview client.AppStoreReview) (*repository.AppReview, error) {
	return s.repository.Create(
		appId,
		appStoreReview.Id,
		appStoreReview.Title,
		appStoreReview.Content,
		appStoreReview.Author,
		appStoreReview.Rating,
		appStoreReview.SubmittedAt,
	)
}

// Fetches up to 10 pages of reviews at a time during initial sync
const (
	MAX_PAGE_LIMIT = 10
)

func (s *AppReviewsService) SyncAppReviews(appId string, moreRecentThan *time.Time) error {
	page := 1

	for {
		reviews, err := s.appStoreClient.GetReviews(appId, page)

		if err != nil {
			return err
		}

		recentReviews := filterReviews(reviews, moreRecentThan)

		if len(recentReviews) == 0 {
			break
		}

		for _, review := range recentReviews {
			newReview, err := s.Create(appId, review)

			if err != nil {
				fmt.Printf("Error saving review %s: %v\n", review.Id, err)
				return err
			}

			fmt.Printf("[App ID: %s] App review successfully saved: %s\n", appId, newReview.Id)
		}

		if page >= MAX_PAGE_LIMIT {
			break
		}

		page++
	}

	fmt.Printf("[App ID: %s] App reviews successfully synced\n", appId)
	return nil
}

func filterReviews(reviews []client.AppStoreReview, moreRecentThan *time.Time) []client.AppStoreReview {
	recentReviews := []client.AppStoreReview{}

	for _, review := range reviews {
		if moreRecentThan == nil || review.SubmittedAt.After(*moreRecentThan) {
			recentReviews = append(recentReviews, review)
		}
	}

	return recentReviews
}
