// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: event.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createEvent = `-- name: CreateEvent :execresult
INSERT INTO events (
    id, organization_id, title, description,
    max_registrant, date, start_time, end_time
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateEventParams struct {
	ID             string
	OrganizationID string
	Title          string
	Description    sql.NullString
	MaxRegistrant  sql.NullInt32
	Date           time.Time
	StartTime      sql.NullTime
	EndTime        sql.NullTime
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEvent,
		arg.ID,
		arg.OrganizationID,
		arg.Title,
		arg.Description,
		arg.MaxRegistrant,
		arg.Date,
		arg.StartTime,
		arg.EndTime,
	)
}

const createEventPricing = `-- name: CreateEventPricing :execresult
INSERT INTO event_pricings (event_id, event_type, price)
VALUES (?, ?, ?)
`

type CreateEventPricingParams struct {
	EventID   string
	EventType sql.NullString
	Price     sql.NullInt32
}

func (q *Queries) CreateEventPricing(ctx context.Context, arg CreateEventPricingParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEventPricing, arg.EventID, arg.EventType, arg.Price)
}

const createSpeaker = `-- name: CreateSpeaker :execresult
INSERT INTO speakers (id, event_id, name, title, description)
VALUES (?, ?, ?, ?, ?)
`

type CreateSpeakerParams struct {
	ID          string
	EventID     sql.NullString
	Name        string
	Title       sql.NullString
	Description sql.NullString
}

func (q *Queries) CreateSpeaker(ctx context.Context, arg CreateSpeakerParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createSpeaker,
		arg.ID,
		arg.EventID,
		arg.Name,
		arg.Title,
		arg.Description,
	)
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?
`

func (q *Queries) DeleteEvent(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const getEventByID = `-- name: GetEventByID :one
SELECT id, organization_id, title, img, description, registrant, max_registrant, date, start_time, end_time, created_at, updated_at FROM events WHERE id = ?
`

func (q *Queries) GetEventByID(ctx context.Context, id string) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEventByID, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Title,
		&i.Img,
		&i.Description,
		&i.Registrant,
		&i.MaxRegistrant,
		&i.Date,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEventPricings = `-- name: GetEventPricings :many
SELECT id, event_id, event_type, price, created_at, updated_at FROM event_pricings WHERE event_id = ?
`

func (q *Queries) GetEventPricings(ctx context.Context, eventID string) ([]EventPricing, error) {
	rows, err := q.db.QueryContext(ctx, getEventPricings, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EventPricing
	for rows.Next() {
		var i EventPricing
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.EventType,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEvents = `-- name: GetEvents :many
SELECT id, organization_id, title, img, description, registrant, max_registrant, date, start_time, end_time, created_at, updated_at FROM events
`

func (q *Queries) GetEvents(ctx context.Context) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Title,
			&i.Img,
			&i.Description,
			&i.Registrant,
			&i.MaxRegistrant,
			&i.Date,
			&i.StartTime,
			&i.EndTime,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventsByCategory = `-- name: GetEventsByCategory :many
SELECT events.id, events.organization_id, events.title, events.img, events.description, events.registrant, events.max_registrant, events.date, events.start_time, events.end_time, events.created_at, events.updated_at FROM events
INNER JOIN user_categories ON user_id = organization_id
WHERE FIND_IN_SET(category_id, ?)
`

func (q *Queries) GetEventsByCategory(ctx context.Context, findINSET string) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getEventsByCategory, findINSET)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Title,
			&i.Img,
			&i.Description,
			&i.Registrant,
			&i.MaxRegistrant,
			&i.Date,
			&i.StartTime,
			&i.EndTime,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSpeakersByEventID = `-- name: GetSpeakersByEventID :many
SELECT id, name, title, img, description, event_id, created_at, updated_at FROM speakers WHERE event_id = ?
`

func (q *Queries) GetSpeakersByEventID(ctx context.Context, eventID sql.NullString) ([]Speaker, error) {
	rows, err := q.db.QueryContext(ctx, getSpeakersByEventID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Speaker
	for rows.Next() {
		var i Speaker
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Title,
			&i.Img,
			&i.Description,
			&i.EventID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
