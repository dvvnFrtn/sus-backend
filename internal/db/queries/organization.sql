-- name: AddOrganization :execresult
INSERT INTO organizations (
    id, user_id, name, description, header_img, profile_img, created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: FindOrganizationById :one
SELECT * FROM organizations WHERE id = ?;

-- name: FindOrganizationByUserId :one
SELECT * FROM organizations WHERE user_id = ?;

-- name: ListOrganization :many
SELECT * FROM organizations;

-- name: UpdateOrganization :execresult
UPDATE organizations
SET name = ?, description = ?, header_img = ?, profile_img = ?
WHERE id = ?;

-- name: DeleteOrganization :exec
DELETE FROM organizations
WHERE id = ?;
