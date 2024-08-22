-- name: AddOrganization :execresult
INSERT INTO organizations (
    id, name, description, header_img, profile_img, created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?);
