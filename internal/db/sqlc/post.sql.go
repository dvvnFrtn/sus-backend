// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addPost = `-- name: AddPost :execresult
INSERT INTO posts (
    id, organization_id, content, image_content
) VALUES (?, ?, ?, ?)
`

type AddPostParams struct {
	ID             string
	OrganizationID string
	Content        string
	ImageContent   sql.NullString
}

func (q *Queries) AddPost(ctx context.Context, arg AddPostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addPost,
		arg.ID,
		arg.OrganizationID,
		arg.Content,
		arg.ImageContent,
	)
}

const commentPost = `-- name: CommentPost :execresult
INSERT INTO post_comments (
    id, user_id, post_id, content
) VALUES (?, ?, ?, ?)
`

type CommentPostParams struct {
	ID      string
	UserID  string
	PostID  string
	Content string
}

func (q *Queries) CommentPost(ctx context.Context, arg CommentPostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, commentPost,
		arg.ID,
		arg.UserID,
		arg.PostID,
		arg.Content,
	)
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?
`

func (q *Queries) DeletePost(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const findPostById = `-- name: FindPostById :one
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img, COUNT(pl.id) AS likes
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
LEFT JOIN post_likes pl ON p.id = pl.post_id
WHERE p.id = ?
GROUP BY p.id
`

type FindPostByIdRow struct {
	ID             string
	Content        string
	ImageContent   sql.NullString
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
	OrganizationID string
	Name           string
	ProfileImg     sql.NullString
	Likes          int64
}

func (q *Queries) FindPostById(ctx context.Context, id string) (FindPostByIdRow, error) {
	row := q.db.QueryRowContext(ctx, findPostById, id)
	var i FindPostByIdRow
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.ImageContent,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrganizationID,
		&i.Name,
		&i.ProfileImg,
		&i.Likes,
	)
	return i, err
}

const findPostByOrganization = `-- name: FindPostByOrganization :many
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img, COUNT(pl.id) AS likes
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
LEFT JOIN post_likes pl ON p.id = pl.post_id
WHERE p.organization_id = ?
GROUP BY p.id
`

type FindPostByOrganizationRow struct {
	ID             string
	Content        string
	ImageContent   sql.NullString
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
	OrganizationID string
	Name           string
	ProfileImg     sql.NullString
	Likes          int64
}

func (q *Queries) FindPostByOrganization(ctx context.Context, organizationID string) ([]FindPostByOrganizationRow, error) {
	rows, err := q.db.QueryContext(ctx, findPostByOrganization, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindPostByOrganizationRow
	for rows.Next() {
		var i FindPostByOrganizationRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.ImageContent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OrganizationID,
			&i.Name,
			&i.ProfileImg,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findPostLikes = `-- name: FindPostLikes :many
SELECT u.name, u.img, pl.liked_at, pl.post_id, pl.user_id
FROM post_likes pl
INNER JOIN users u ON pl.user_id = u.id
WHERE post_id = ?
`

type FindPostLikesRow struct {
	Name    string
	Img     sql.NullString
	LikedAt sql.NullTime
	PostID  string
	UserID  string
}

func (q *Queries) FindPostLikes(ctx context.Context, postID string) ([]FindPostLikesRow, error) {
	rows, err := q.db.QueryContext(ctx, findPostLikes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindPostLikesRow
	for rows.Next() {
		var i FindPostLikesRow
		if err := rows.Scan(
			&i.Name,
			&i.Img,
			&i.LikedAt,
			&i.PostID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isLiked = `-- name: IsLiked :one
SELECT COUNT(1) FROM post_likes WHERE user_id = ? AND post_id = ?
`

type IsLikedParams struct {
	UserID string
	PostID string
}

func (q *Queries) IsLiked(ctx context.Context, arg IsLikedParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, isLiked, arg.UserID, arg.PostID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const likedPost = `-- name: LikedPost :execresult
INSERT INTO post_likes (
    user_id, post_id
) VALUES (?, ?)
`

type LikedPostParams struct {
	UserID string
	PostID string
}

func (q *Queries) LikedPost(ctx context.Context, arg LikedPostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, likedPost, arg.UserID, arg.PostID)
}

const listPosts = `-- name: ListPosts :many
SELECT p.id, p.content, p.image_content, p.created_at, p.updated_at, p.organization_id, o.name, o.profile_img, COUNT(pl.id) AS likes
FROM posts p
INNER JOIN organizations o ON p.organization_id = o.id
LEFT JOIN post_likes pl ON p.id = pl.post_id
GROUP BY p.id
`

type ListPostsRow struct {
	ID             string
	Content        string
	ImageContent   sql.NullString
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
	OrganizationID string
	Name           string
	ProfileImg     sql.NullString
	Likes          int64
}

func (q *Queries) ListPosts(ctx context.Context) ([]ListPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPostsRow
	for rows.Next() {
		var i ListPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.ImageContent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OrganizationID,
			&i.Name,
			&i.ProfileImg,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unlikedPost = `-- name: UnlikedPost :exec
DELETE FROM post_likes WHERE user_id = ? AND post_id = ?
`

type UnlikedPostParams struct {
	UserID string
	PostID string
}

func (q *Queries) UnlikedPost(ctx context.Context, arg UnlikedPostParams) error {
	_, err := q.db.ExecContext(ctx, unlikedPost, arg.UserID, arg.PostID)
	return err
}
