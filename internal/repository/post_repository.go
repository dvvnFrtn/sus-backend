package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type PostRepository interface {
	Create(sqlc.AddPostParams) (sql.Result, error)
	FindById(string) (sqlc.FindPostByIdRow, error)
	ListAll() ([]sqlc.FindPostByIdRow, error)
	Delete(string) error
	FindByOrganization(string) ([]sqlc.FindPostByIdRow, error)
	LikedPost(sqlc.LikedPostParams) (sql.Result, error)
	UnlikedPost(sqlc.UnlikedPostParams) error
	IsLiked(sqlc.IsLikedParams) (int64, error)
}

type postRepository struct {
	db *sqlc.Queries
}

func NewPostRepository(db *sqlc.Queries) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(in sqlc.AddPostParams) (sql.Result, error) {
	return r.db.AddPost(context.Background(), in)
}

func (r *postRepository) FindById(id string) (sqlc.FindPostByIdRow, error) {
	return r.db.FindPostById(context.Background(), id)
}

func (r *postRepository) ListAll() ([]sqlc.FindPostByIdRow, error) {
	rows, err := r.db.ListPosts(context.Background())
	if err != nil {
		return nil, err
	}

	var posts []sqlc.FindPostByIdRow
	for _, p := range rows {
		temp := sqlc.FindPostByIdRow(p)
		posts = append(posts, temp)
	}

	return posts, nil
}

func (r *postRepository) Delete(id string) error {
	return r.db.DeletePost(context.Background(), id)
}

func (r *postRepository) FindByOrganization(orgId string) ([]sqlc.FindPostByIdRow, error) {
	rows, err := r.db.FindPostByOrganization(context.Background(), orgId)
	if err != nil {
		return nil, err
	}

	var posts []sqlc.FindPostByIdRow
	for _, p := range rows {
		temp := sqlc.FindPostByIdRow(p)
		posts = append(posts, temp)
	}

	return posts, nil
}

func (r *postRepository) LikedPost(in sqlc.LikedPostParams) (sql.Result, error) {
	return r.db.LikedPost(context.Background(), in)
}

func (r *postRepository) IsLiked(in sqlc.IsLikedParams) (int64, error) {
	return r.db.IsLiked(context.Background(), in)
}

func (r *postRepository) UnlikedPost(in sqlc.UnlikedPostParams) error {
	return r.db.UnlikedPost(context.Background(), in)
}
