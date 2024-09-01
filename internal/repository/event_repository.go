package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type EventRepository interface {
	GetEvents() ([]sqlc.Event, error)
	GetEventByID(string) (sqlc.Event, error)
	CreateEvent(sqlc.CreateEventParams) (sql.Result, error)
}

type eventRepository struct {
	db *sqlc.Queries
}

func NewEventRepository(db *sqlc.Queries) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetEvents() ([]sqlc.Event, error) {
	return r.db.GetEvents(context.Background())
}

func (r *eventRepository) GetEventByID(id string) (sqlc.Event, error) {
	return r.db.GetEventByID(context.Background(), id)
}

func (r *eventRepository) CreateEvent(input sqlc.CreateEventParams) (sql.Result, error) {
	return r.db.CreateEvent(context.Background(), input)
}
