package main

import (
	// Needs to be the first import since it loads the env vars
	_ "github.com/joho/godotenv/autoload" 

	"fmt"
	"encoding/json"
	"net/http"
	"strconv"

	"ios-app-reviews-viewer.com/m/internal/database"
	"ios-app-reviews-viewer.com/m/internal/client"
	"ios-app-reviews-viewer.com/m/internal/repository"
)

func main() {
	db := database.Connect()

	monitoredAppsRepository := repository.NewMonitoredAppsRepository(db)
	appReviewsRepository := repository.NewAppReviewsRepository(db)

	appStoreClient := client.NewAppStoreClient()
	
	http.HandleFunc("/reviews", func(w http.ResponseWriter, r *http.Request) {
		appId := r.URL.Query().Get("appId")
		page := r.URL.Query().Get("page")
		
		if page == "" {
			page = "1"
		}

		formattedPage, err := strconv.Atoi(page)
		
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}

		reviews, err := appStoreClient.GetReviews(appId, formattedPage)

		if err != nil {
			http.Error(w, "Unable to fetch reviews from App Store", http.StatusInternalServerError)
			return
		}
		
		response := map[string]interface{}{
			"data": reviews,
			"page": formattedPage,
		}
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})
	
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
