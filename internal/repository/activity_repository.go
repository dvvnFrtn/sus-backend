package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type ActivityRepository interface {
	GetActivityByID(string) (sqlc.Activity, error)
	GetActivitiesByOrganizationID(string) ([]sqlc.Activity, error)
	CreateActivity(sqlc.CreateActivityParams) (sql.Result, error)
	DeleteActivity(string) error
}

type activityRepository struct {
	db *sqlc.Queries
}

func NewActivityRepository(db *sqlc.Queries) ActivityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) GetActivityByID(id string) (sqlc.Activity, error) {
	return r.db.GetActivityByID(context.Background(), id)
}

func (r *activityRepository) GetActivitiesByOrganizationID(org_id string) ([]sqlc.Activity, error) {
	return r.db.GetActivitiesByOrganizationID(context.Background(), org_id)
}

func (r *activityRepository) CreateActivity(input sqlc.CreateActivityParams) (sql.Result, error) {
	return r.db.CreateActivity(context.Background(), input)
}

func (r *activityRepository) DeleteActivity(id string) error {
	return r.db.DeleteActivity(context.Background(), id)
}
