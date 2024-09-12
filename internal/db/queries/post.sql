-- name: AddPost :execresult
INSERT INTO posts (
    id, organization_id, content, image_content
) VALUES (?, ?, ?, ?);

-- name: FindPostById :one
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
WHERE p.id = ?;

-- name: ListPosts :many
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?;

-- name: FindPostByOrganization :many
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
WHERE p.organization_id = ?;

-- name: LikedPost :execresult
INSERT INTO post_likes (
    user_id, post_id
) VALUES (?, ?);

-- name: IsLiked :one
SELECT COUNT(1) FROM post_likes WHERE user_id = ? AND post_id = ?;

-- name: UnlikedPost :exec
DELETE FROM post_likes WHERE user_id = ? AND post_id = ?;
