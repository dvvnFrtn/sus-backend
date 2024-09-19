-- name: CreateEvent :execresult
INSERT INTO events (
    id, organization_id, title,
    description, max_registrant, date
) VALUES (?, ?, ?, ?, ?, ?);

-- name: GetEventByID :one
SELECT * FROM events WHERE id = ?;

-- name: GetEvents :many
SELECT * FROM events;

-- name: GetEventsByCategory :many
SELECT events.* FROM events
INNER JOIN user_categories ON user_id = organization_id
WHERE FIND_IN_SET(category_id, ?);

-- name: GetEventsByOrganizationID :many
SELECT * FROM events WHERE organization_id = ?;

-- name: GetEventsOfFollowedOrganizations :many
SELECT events.* FROM events
INNER JOIN followers ON organization_id = followers.organization_id
WHERE follower_id = ?;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?;

-- name: CreateEventPricing :execresult
INSERT INTO event_pricings (event_id, event_type, price)
VALUES (?, ?, ?);

-- name: GetEventPricings :many
SELECT * FROM event_pricings WHERE event_id = ?;

-- name: CreateEventAgenda :execresult
INSERT INTO event_agendas (
    id, event_id, title, description,
    start_time, end_time, location
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetAgendasByEventID :many
SELECT * FROM event_agendas WHERE event_id = ?;

-- name: CreateSpeaker :execresult
INSERT INTO speakers (id, agenda_id, name, title, description)
VALUES (?, ?, ?, ?, ?);

-- name: GetSpeakersByAgendaID :many
SELECT * FROM speakers WHERE agenda_id = ?;