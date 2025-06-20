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

func (s *MonitoredAppsService) FindById(id string) (*repository.MonitoredApp, error) {
	return s.repository.FindById(id)
}

func (s *MonitoredAppsService) UpdateLastSyncedAt(id string, lastSyncedAt time.Time) (*repository.MonitoredApp, error) {
	return s.repository.UpdateLastSyncedAt(id, lastSyncedAt)
}