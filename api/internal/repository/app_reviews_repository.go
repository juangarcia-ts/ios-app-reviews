package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AppReview struct {
	Id          string    `json:"id" db:"id"`
	AppId       string    `json:"app_id" db:"app_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	Author      string    `json:"author" db:"author"`
	Rating      int       `json:"rating" db:"rating"`
	SubmittedAt time.Time `json:"submitted_at" db:"submitted_at"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type AppReviewsRepository struct {
	db *sqlx.DB
}

func NewAppReviewsRepository(db *sqlx.DB) *AppReviewsRepository {
	return &AppReviewsRepository{db: db}
}

func (r *AppReviewsRepository) FindAll(appId string, limit int, offset int, startsAt time.Time, endsAt time.Time) ([]AppReview, error) {
	appReviews := make([]AppReview, 0)
	query := "SELECT * FROM app_reviews WHERE app_id = $1 AND submitted_at BETWEEN $2 AND $3 ORDER BY submitted_at DESC LIMIT $4 OFFSET $5"
	err := r.db.Select(&appReviews, query, appId, startsAt, endsAt, limit, offset)
	return appReviews, err
}

func (r *AppReviewsRepository) Count(appId string, startsAt time.Time, endsAt time.Time) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM app_reviews WHERE app_id = $1 AND submitted_at BETWEEN $2 AND $3"
	err := r.db.Get(&count, query, appId, startsAt, endsAt)
	return count, err
}

func (r *AppReviewsRepository) FindById(id string) (*AppReview, error) {
	appReview := AppReview{}
	query := "SELECT * FROM app_reviews WHERE id = $1"
	err := r.db.Get(&appReview, query, id)
	return &appReview, err
}

func (r *AppReviewsRepository) Create(
	appId string,
	id string,
	title string,
	content string,
	author string,
	rating int,
	submittedAt time.Time,
) (*AppReview, error) {
	query := `INSERT INTO app_reviews ("app_id", "id", "title", "content", "author", "rating", "submitted_at") VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO NOTHING`
	_, err := r.db.Exec(query, appId, id, title, content, author, rating, submittedAt)

	if err != nil {
		return nil, err
	}

	return r.FindById(id)
}
