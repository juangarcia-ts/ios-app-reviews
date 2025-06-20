package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type AppReview struct {
	Id           string    `json:"id" db:"id"`
	AppId        string    `json:"app_id" db:"app_id"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	Author       string    `json:"author" db:"author"`
	Rating       int       `json:"rating" db:"rating"`
	SubmittedAt  time.Time `json:"submitted_at" db:"submitted_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type AppReviewsRepository struct {
	db *sqlx.DB
}

func NewAppReviewsRepository(db *sqlx.DB) *AppReviewsRepository {
	return &AppReviewsRepository{db: db}
}

func (r *AppReviewsRepository) FindAll(limit int, offset int) ([]AppReview, error) {
	appReviews := make([]AppReview, 0)
	query := "SELECT * FROM app_reviews LIMIT $1 OFFSET $2"
	err := r.db.Select(&appReviews, query)
	return monitoredApps, err
}

func (r *AppReviewsRepository) Create(appId string, title string, content string, author string, rating int, submittedAt time.Time) (*AppReview, error) {
	query := `INSERT INTO app_reviews ("app_id", "title", "content", "author", "rating", "submitted_at") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	insertedRow := r.db.QueryRow(query, appId, title, content, author, rating, submittedAt)

	var insertedRowId string
	if err := insertedRow.Scan(&insertedRowId); err != nil {
		return nil, err
	}

	return r.FindById(insertedRowId)
}