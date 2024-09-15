package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type EventRepository interface {
	GetEvents() ([]sqlc.Event, error)
	GetEventsByCategory(string) ([]sqlc.Event, error)
	GetEventByID(string) (sqlc.Event, error)
	CreateEvent(sqlc.CreateEventParams) (sql.Result, error)
	DeleteEvent(string) error
	GetPricingsForEvent(string) ([]sqlc.EventPricing, error)
	CreateEventPricing(sqlc.CreateEventPricingParams) (sql.Result, error)
	GetSpeakersForEvent(string) ([]sqlc.Speaker, error)
	CreateSpeaker(sqlc.CreateSpeakerParams) (sql.Result, error)
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

func (r *eventRepository) GetEventsByCategory(ctg_ids string) ([]sqlc.Event, error) {
	return r.db.GetEventsByCategory(context.Background(), ctg_ids)
}

func (r *eventRepository) DeleteEvent(id string) error {
	return r.db.DeleteEvent(context.Background(), id)
}

func (r *eventRepository) GetPricingsForEvent(event_id string) ([]sqlc.EventPricing, error) {
	return r.db.GetEventPricings(context.Background(), event_id)
}

func (r *eventRepository) CreateEventPricing(input sqlc.CreateEventPricingParams) (sql.Result, error) {
	return r.db.CreateEventPricing(context.Background(), input)
}

func (r *eventRepository) GetSpeakersForEvent(agenda_id string) ([]sqlc.Speaker, error) {
	return r.db.GetSpeakersByEventID(context.Background(), agenda_id)
}

func (r *eventRepository) CreateSpeaker(input sqlc.CreateSpeakerParams) (sql.Result, error) {
	return r.db.CreateSpeaker(context.Background(), input)
}
