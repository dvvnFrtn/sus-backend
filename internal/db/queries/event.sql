-- name: CreateEvent :execresult
INSERT INTO events (
    id, organization_id, title, description,
    max_registrant, date, start_time, end_time
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetEventByID :one
SELECT * FROM events WHERE id = ?;

-- name: GetEvents :many
SELECT * FROM events;

-- name: GetEventsByCategory :many
