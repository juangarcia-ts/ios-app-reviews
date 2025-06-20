package controller

import (
	"net/http"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"
	"ios-app-reviews-viewer.com/m/internal/service"
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
	appReviews, err := c.appReviewsService.FindAll(appId, parsedLimit, offset)
	
	response := map[string]interface{}{
		"data": appReviews,
		"page": parsedPage,
		"limit": parsedLimit,
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}