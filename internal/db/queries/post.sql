-- name: AddPost :execresult
INSERT INTO posts (
    id, organization_id, content, image_content, created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?);

-- name: FindPostById :one
SELECT * FROM posts WHERE id = ?;

-- name: ListPosts :many
SELECT * from posts;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?;

-- name: FindPostByOrganization :many
SELECT * FROM posts WHERE organization_id = ?;
