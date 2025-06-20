package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type MonitoredApp struct {
	AppId        string     `json:"app_id" db:"app_id"`
	AppName      string     `json:"app_name" db:"app_name"`
	Nickname     string     `json:"nickname" db:"nickname"`
	LogoUrl      string     `json:"logo_url" db:"logo_url"`
	LastSyncedAt *time.Time `json:"last_synced_at" db:"last_synced_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

type MonitoredAppsRepository struct {
	db *sqlx.DB
}

func NewMonitoredAppsRepository(db *sqlx.DB) *MonitoredAppsRepository {
	return &MonitoredAppsRepository{db: db}
}

func (r *MonitoredAppsRepository) FindAll() ([]MonitoredApp, error) {
	monitoredApps := make([]MonitoredApp, 0)
	query := "SELECT * FROM monitored_apps ORDER BY created_at DESC" // Guaranteed to be in the correct order
	err := r.db.Select(&monitoredApps, query)
	return monitoredApps, err
}

func (r *MonitoredAppsRepository) FindById(id string) (*MonitoredApp, error) {
	monitoredApp := MonitoredApp{}
	query := "SELECT * FROM monitored_apps WHERE app_id = $1"
	err := r.db.Get(&monitoredApp, query, id)
	return &monitoredApp, err
}

func (r *MonitoredAppsRepository) Create(appId string, appName string, logoUrl string, nickname *string) (*MonitoredApp, error) {
	query := `INSERT INTO monitored_apps ("app_id", "app_name", "logo_url", "nickname") VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, appId, appName, logoUrl, nickname)

	if err != nil {
		return nil, err
	}

	return r.FindById(appId)
}

func (r *MonitoredAppsRepository) UpdateLastSyncedAt(appId string, lastSyncedAt time.Time) (*MonitoredApp, error) {
	query := "UPDATE monitored_apps SET last_synced_at = $1 WHERE app_id = $2"
	_, err := r.db.Exec(query, lastSyncedAt, appId)

	if err != nil {
		return nil, err
	}

	return r.FindById(appId)
}

func (r *MonitoredAppsRepository) Delete(appId string) error {
	query := "DELETE FROM monitored_apps WHERE app_id = $1"
	result := r.db.MustExec(query, appId)
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("app not found")
	}

	return nil
}
