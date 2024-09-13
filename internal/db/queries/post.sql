-- name: AddPost :execresult
INSERT INTO posts (
    id, organization_id, content, image_content
) VALUES (?, ?, ?, ?);

-- name: FindPostById :one
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img, COUNT(pl.id) AS likes, COUNT(pc.id) AS comments
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
LEFT JOIN post_likes pl ON p.id = pl.post_id
LEFT JOIN post_comments pc ON p.id = pc.post_id
WHERE p.id = ?
GROUP BY p.id;

-- name: ListPosts :many
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img, COUNT(pl.id) AS likes, COUNT(pc.id) AS comments
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
LEFT JOIN post_likes pl ON p.id = pl.post_id
LEFT JOIN post_comments pc ON p.id = pc.post_id
GROUP BY p.id;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?;

-- name: FindPostByOrganization :many
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img, COUNT(pl.id) AS likes, COUNT(pc.id) AS comments
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
LEFT JOIN post_likes pl ON p.id = pl.post_id
LEFT JOIN post_comments pc ON p.id = pc.post_id
WHERE p.organization_id = ?
GROUP BY p.id;

-- name: LikedPost :execresult
INSERT INTO post_likes (
    user_id, post_id
) VALUES (?, ?);

-- name: IsLiked :one
SELECT COUNT(1) FROM post_likes WHERE user_id = ? AND post_id = ?;

-- name: UnlikedPost :exec
DELETE FROM post_likes WHERE user_id = ? AND post_id = ?;

-- name: FindPostLikes :many
SELECT u.name, u.img, pl.liked_at, pl.post_id, pl.user_id
FROM post_likes pl
INNER JOIN users u ON pl.user_id = u.id
WHERE post_id = ?;

-- name: CommentPost :execresult
INSERT INTO post_comments (
    id, user_id, post_id, content
) VALUES (?, ?, ?, ?);

-- name: FindPostComments :many
SELECT pc.id, pc.post_id, pc.user_id, pc.content, pc.created_at, u.name, u.img
FROM post_comments pc
INNER JOIN users u ON pc.user_id = u.id
WHERE pc.post_id = ?;

-- name: DeleteComment :exec
DELETE FROM post_comments WHERE id = ?;

-- name: FindCommentById :one
SELECT pc.id, pc.post_id, pc.user_id, pc.content, pc.created_at, u.name, u.img
FROM post_comments pc
INNER JOIN users u ON pc.user_id = u.id
WHERE pc.id = ?;
