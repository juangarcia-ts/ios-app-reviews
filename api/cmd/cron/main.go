package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ios-app-reviews-viewer.com/m/internal/client"
	"ios-app-reviews-viewer.com/m/internal/database"
	"ios-app-reviews-viewer.com/m/internal/repository"
	"ios-app-reviews-viewer.com/m/internal/service"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	scheduler, err := gocron.NewScheduler()
	job, err := scheduler.NewJob(
		gocron.CronJob("*/5 * * * *", false), // Every 5 minutes
		gocron.NewTask(syncAllMonitoredApps),
		gocron.WithStartAt(
			gocron.WithStartImmediately(),
		),
	)

	if err != nil {
		log.Fatal("Unable to create cron job: ", err)
	}

	fmt.Printf("Starting cron job %s", job.ID())
	scheduler.Start()

	select {}
}

func syncAllMonitoredApps() {
	// Set up dependencies
	db := database.Connect()
	appStoreClient := client.NewAppStoreClient()
	monitoredAppsRepository := repository.NewMonitoredAppsRepository(db)
	appReviewsRepository := repository.NewAppReviewsRepository(db)
	appReviewsService := service.NewAppReviewsService(appReviewsRepository, appStoreClient)
	monitoredAppsService := service.NewMonitoredAppsService(monitoredAppsRepository)

	// Sync all monitored apps
	fmt.Println("Syncing all monitored apps")
	monitoredApps, err := monitoredAppsService.FindAll()

	if err != nil {
		log.Fatal("Unable to find all monitored apps: ", err)
	}

	for _, monitoredApp := range monitoredApps {
		err := appReviewsService.SyncAppReviews(monitoredApp.AppId, monitoredApp.LastSyncedAt)

		if err == nil {
			monitoredAppsService.UpdateLastSyncedAt(monitoredApp.AppId, time.Now())
			fmt.Printf("[App ID: %s] Last synced at successfully updated\n", monitoredApp.AppId)
		}
	}

	// Start a simple HTTP server for health checks
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8000", nil)
}
