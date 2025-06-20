package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"ios-app-reviews-viewer.com/m/internal/service"
	"ios-app-reviews-viewer.com/m/internal/helpers"
)

type AppReviewsController struct {
	appReviewsService *service.AppReviewsService
}

func NewAppReviewsController(appReviewsService *service.AppReviewsService) *AppReviewsController {
	return &AppReviewsController{appReviewsService: appReviewsService}
}

func (c *AppReviewsController) GetAppReviews(w http.ResponseWriter, r *http.Request) {
	appId := mux.Vars(r)["appId"]
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	startsAt := r.URL.Query().Get("startsAt")
	endsAt := r.URL.Query().Get("endsAt")

	if appId == "" {
		http.Error(w, "App ID is required", http.StatusBadRequest)
		return
	}

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	if startsAt == "" {
		startsAt = time.Now().Add(-48 * time.Hour).Format(time.RFC3339)
	} 

	if endsAt == "" {
		endsAt = time.Now().Format(time.RFC3339)
	}
	
	parsedStartsAt, err := helpers.ParseDateTime(startsAt)

	if err != nil {
		http.Error(w, "Invalid startsAt date", http.StatusBadRequest)
		return
	}
	
	parsedEndsAt, err := helpers.ParseDateTime(endsAt)

	if err != nil {
		http.Error(w, "Invalid endsAt date", http.StatusBadRequest)
		return
	}

	parsedLimit, err := strconv.Atoi(limit)

	if err != nil {
		http.Error(w, "Invalid limit number", http.StatusBadRequest)
		return
	}

	parsedPage, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	offset := (parsedPage - 1) * parsedLimit
	appReviews, count, err := c.appReviewsService.FindAllWithCount(appId, parsedLimit, offset, parsedStartsAt, parsedEndsAt)

	response := map[string]interface{}{
		"data":       appReviews,
		"page":       parsedPage,
		"limit":      parsedLimit,
		"total":      count,
		"totalPages": math.Ceil(float64(count) / float64(parsedLimit)),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
