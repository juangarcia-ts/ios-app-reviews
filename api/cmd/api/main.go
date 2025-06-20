package main

import (
	// Needs to be the first import since it loads the env vars
	_ "github.com/joho/godotenv/autoload"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"

	"ios-app-reviews-viewer.com/m/internal/client"
	"ios-app-reviews-viewer.com/m/internal/controller"
	"ios-app-reviews-viewer.com/m/internal/database"
	"ios-app-reviews-viewer.com/m/internal/repository"
	"ios-app-reviews-viewer.com/m/internal/service"
)

func main() {
	// Set up dependencies
	db := database.Connect()
	appStoreClient := client.NewAppStoreClient()
	appReviewsRepository := repository.NewAppReviewsRepository(db)
	monitoredAppsRepository := repository.NewMonitoredAppsRepository(db)
	appReviewsService := service.NewAppReviewsService(appReviewsRepository, appStoreClient)
	monitoredAppsService := service.NewMonitoredAppsService(monitoredAppsRepository)
	appReviewsController := controller.NewAppReviewsController(appReviewsService)
	monitoredAppsController := controller.NewMonitoredAppsController(appReviewsService, monitoredAppsService, appStoreClient)

	// Create Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/v1/apps", monitoredAppsController.GetMonitoredApps).Methods("GET")
	router.HandleFunc("/api/v1/apps", monitoredAppsController.CreateMonitoredApp).Methods("POST")
	router.HandleFunc("/api/v1/apps/{appId}", monitoredAppsController.GetMonitoredApp).Methods("GET")
	router.HandleFunc("/api/v1/apps/{appId}", monitoredAppsController.DeleteMonitoredApp).Methods("DELETE")
	router.HandleFunc("/api/v1/apps/{appId}/reviews", appReviewsController.GetAppReviews).Methods("GET")
	router.HandleFunc("/api/v1/apps/{appId}/lookup", monitoredAppsController.GetAppInfoFromAppStore).Methods("GET")
	router.HandleFunc("/api/v1/apps/{appId}/sync", monitoredAppsController.SyncReviews).Methods("POST")

	// Special route for health checks
	router.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("HEAD")

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Apply CORS middleware to router
	handler := c.Handler(router)

	// Start server
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", handler)
}
