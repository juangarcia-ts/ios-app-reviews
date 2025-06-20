package main

import (
	// Needs to be the first import since it loads the env vars
	_ "github.com/joho/godotenv/autoload" 

	"fmt"
	"net/http"

	"ios-app-reviews-viewer.com/m/internal/database"
	"ios-app-reviews-viewer.com/m/internal/client"
	"ios-app-reviews-viewer.com/m/internal/service"
	"ios-app-reviews-viewer.com/m/internal/repository"
	"ios-app-reviews-viewer.com/m/internal/controller"
)

func main() {
	// Set up dependencies
	db := database.Connect()
	appStoreClient := client.NewAppStoreClient()
	appReviewsRepository := repository.NewAppReviewsRepository(db)
	monitoredAppsRepository := repository.NewMonitoredAppsRepository(db)
	appReviewsService := service.NewAppReviewsService(appReviewsRepository, appStoreClient)
	service.NewMonitoredAppsService(monitoredAppsRepository)
	appReviewsController := controller.NewAppReviewsController(appReviewsService)
	http.HandleFunc("/reviews", appReviewsController.GetAppReviews)	

	// Start server
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
