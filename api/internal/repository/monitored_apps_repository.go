package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type MonitoredApp struct {
	Id             				string    	 `json:"id" db:"id"`
	AppId            			string    	 `json:"app_id" db:"app_id"`
	Nickname        			string    	 `json:"nickname" db:"nickname"`
	LastSyncedAt 				time.Time 	 `json:"last_synced_at" db:"last_synced_at"`
	CreatedAt      				time.Time 	 `json:"created_at" db:"created_at"`
	UpdatedAt      				time.Time 	 `json:"updated_at" db:"updated_at"`
}

type MonitoredAppsRepository struct {
	db *sqlx.DB
}

func NewMonitoredAppsRepository(db *sqlx.DB) *MonitoredAppsRepository {
	return &MonitoredAppsRepository{db: db}
}

func (r *MonitoredAppsRepository) FindAll() ([]MonitoredApp, error) {
	monitoredApps := make([]MonitoredApp, 0)
	query := "SELECT * FROM monitored_apps"
	err := r.db.Select(&monitoredApps, query)
	return monitoredApps, err
}

func (r *MonitoredAppsRepository) FindById(id string) (*MonitoredApp, error) {
	monitoredApp := MonitoredApp{}
	query := "SELECT * FROM monitored_apps WHERE id = $1"
	err := r.db.Get(&monitoredApp, query, id)
	return &monitoredApp, err
}

func (r *MonitoredAppsRepository) Create(appId string, nickname string) (*MonitoredApp, error) {
	query := `INSERT INTO monitored_apps ("app_id", "nickname") VALUES ($1, $2) RETURNING id`
	insertedRow := r.db.QueryRow(query, appId, nickname)

	var insertedRowId string
	if err := insertedRow.Scan(&insertedRowId); err != nil {
		return nil, err
	}

	return r.FindById(insertedRowId)
}

func (r *MonitoredAppsRepository) UpdateLastSyncedAt(id string, lastSyncedAt time.Time) (*MonitoredApp, error) {
	query := "UPDATE monitored_apps SET last_synced_at = $1 WHERE id = $2 RETURNING id"
	updatedRow := r.db.QueryRow(query, lastSyncedAt, id)

	var updatedRowId string
	if err := updatedRow.Scan(&updatedRowId); err != nil {
		return nil, err
	}

	return r.FindById(updatedRowId)
}

func (r *MonitoredAppsRepository) Delete(id string) (string, error) {
	query := "DELETE FROM monitored_apps WHERE id = $1"
	result := r.db.MustExec(query, id)

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return "", err
	} else if rowsAffected == 0 {
		return "", nil
	} else {
		return id, nil
	}
}