-- name: GetActivitiesByOrganizationID :many
SELECT * FROM activities WHERE organization_id = ?;

-- name: GetActivityByID :one
SELECT * FROM activities WHERE id = ?;

-- name: CreateActivity :execresult
INSERT INTO activities (id, organization_id, title, note, start_time, end_time)
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteActivity :exec
DELETE FROM activities WHERE id = ?;