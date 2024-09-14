-- name: AddOrganization :execresult
INSERT INTO organizations (
    id, user_id, name, description, header_img, profile_img
) VALUES (?, ?, ?, ?, ?, ?);

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

-- name: IsOrganizationExist :one
SELECT COUNT(1) FROM organizations INNER JOIN users ON organizations.user_id = users.id WHERE user_id = ?;

-- name: FollowOrganizaiton :execresult
INSERT INTO followers (
    organization_id, follower_id
) VALUES (
    ?, ?
);

-- name: UnfollowOrganization :exec
DELETE FROM followers WHERE organization_id = ? AND follower_id = ?;

-- name: IsFollowed :one
SELECT COUNT(1) FROM followers WHERE organization_id = ? AND follower_id = ?;
