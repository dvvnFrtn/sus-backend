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
SELECT events.* FROM events
INNER JOIN user_categories ON user_id = organization_id
WHERE FIND_IN_SET(category_id, ?);

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?;
