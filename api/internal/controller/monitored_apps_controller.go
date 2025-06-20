package controller

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"ios-app-reviews-viewer.com/m/internal/service"
)

type MonitoredAppsController struct {
	monitoredAppsService *service.MonitoredAppsService
}

func NewMonitoredAppsController(monitoredAppsService *service.MonitoredAppsService) *MonitoredAppsController {
	return &MonitoredAppsController{monitoredAppsService: monitoredAppsService}
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