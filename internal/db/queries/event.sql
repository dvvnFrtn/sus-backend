-- name: CreateEvent :execresult
INSERT INTO events (
    id, organization_id, title, description,
    max_registrant, date
) VALUES (?, ?, ?, ?, ?, ?);

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

-- name: CreateEventPricing :execresult
INSERT INTO event_pricings (event_id, event_type, price)
VALUES (?, ?, ?);

-- name: GetEventPricings :many
SELECT * FROM event_pricings WHERE event_id = ?;

-- name: CreateSpeaker :execresult
INSERT INTO speakers (id, agenda_id, name, title, description)
VALUES (?, ?, ?, ?, ?);

-- name: GetSpeakersByEventID :many
SELECT * FROM speakers WHERE agenda_id = ?;