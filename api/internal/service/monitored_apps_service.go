package service

import (
	"time"

	"ios-app-reviews-viewer.com/m/internal/repository"
)

type MonitoredAppsService struct {
	repository *repository.MonitoredAppsRepository
}

func NewMonitoredAppsService(repository *repository.MonitoredAppsRepository) *MonitoredAppsService {
	return &MonitoredAppsService{repository: repository}
}

func (s *MonitoredAppsService) FindAll() ([]repository.MonitoredApp, error) {
	return s.repository.FindAll()
}

func (s *MonitoredAppsService) FindById(appId string) (*repository.MonitoredApp, error) {
	return s.repository.FindById(appId)
}

func (s *MonitoredAppsService) Create(appId, appName, logoUrl string, nickname *string) (*repository.MonitoredApp, error) {
	return s.repository.Create(appId, appName, logoUrl, nickname)
}

func (s *MonitoredAppsService) UpdateLastSyncedAt(appId string, lastSyncedAt time.Time) (*repository.MonitoredApp, error) {
	return s.repository.UpdateLastSyncedAt(appId, lastSyncedAt)
}

func (s *MonitoredAppsService) Delete(appId string) error {
	return s.repository.Delete(appId)
}
