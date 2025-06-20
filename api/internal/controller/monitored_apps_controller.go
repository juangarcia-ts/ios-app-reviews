package controller

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"

	"github.com/gorilla/mux"
	"ios-app-reviews-viewer.com/m/internal/service"
	"ios-app-reviews-viewer.com/m/internal/client"
)

type MonitoredAppsController struct {
	appReviewsService *service.AppReviewsService
	monitoredAppsService *service.MonitoredAppsService
	appStoreClient *client.AppStoreClient
}

func NewMonitoredAppsController(appReviewsService *service.AppReviewsService, monitoredAppsService *service.MonitoredAppsService, appStoreClient *client.AppStoreClient) *MonitoredAppsController {
	return &MonitoredAppsController{appReviewsService: appReviewsService, monitoredAppsService: monitoredAppsService, appStoreClient: appStoreClient}
}

func (c *MonitoredAppsController) GetMonitoredApps(w http.ResponseWriter, r *http.Request) {
	monitoredApps, err := c.monitoredAppsService.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(monitoredApps)
}

func (c *MonitoredAppsController) GetMonitoredApp(w http.ResponseWriter, r *http.Request) {
	appId := mux.Vars(r)["appId"]
	monitoredApp, err := c.monitoredAppsService.FindById(appId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(monitoredApp)
}

func (c *MonitoredAppsController) CreateMonitoredApp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AppId    string `json:"app_id"`
		AppName  string `json:"app_name"`
		LogoUrl  string `json:"logo_url"`
		Nickname *string `json:"nickname"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	appId := req.AppId
	appName := req.AppName
	logoUrl := req.LogoUrl
	nickname := req.Nickname

	if appId == "" {
		http.Error(w, "Missing required fields: app_id", http.StatusBadRequest)
		return
	}

	if appName == "" {
		http.Error(w, "Missing required fields: app_name", http.StatusBadRequest)
		return
	}

	if logoUrl == "" {
		http.Error(w, "Missing required fields: logo_url", http.StatusBadRequest)
		return
	}

	monitoredApp, err := c.monitoredAppsService.Create(appId, appName, logoUrl, nickname)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(monitoredApp)
}

func (c *MonitoredAppsController) DeleteMonitoredApp(w http.ResponseWriter, r *http.Request) {
	appId := mux.Vars(r)["appId"]

	if err := c.monitoredAppsService.Delete(appId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *MonitoredAppsController) GetAppInfoFromAppStore(w http.ResponseWriter, r *http.Request) {
	appId := mux.Vars(r)["appId"]
	appInfo, err := c.appStoreClient.GetAppInfoFromAppStore(appId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appInfo)
}

func (c *MonitoredAppsController) SyncReviews(w http.ResponseWriter, r *http.Request) {
	appId := mux.Vars(r)["appId"]

	monitoredApp, err := c.monitoredAppsService.FindById(appId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = c.appReviewsService.SyncAppReviews(monitoredApp.AppId, monitoredApp.LastSyncedAt)
		
	if err == nil {
		c.monitoredAppsService.UpdateLastSyncedAt(monitoredApp.AppId, time.Now())
		fmt.Printf("[App ID: %s] Last synced at successfully updated\n", monitoredApp.AppId)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}