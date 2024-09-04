package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type PostRepository interface {
	Create(sqlc.AddPostParams) (sql.Result, error)
	FindById(string) (sqlc.Post, error)
	ListAll() ([]sqlc.Post, error)
	Delete(string) error
	FindByOrganization(orgId string) ([]sqlc.Post, error)
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

func (r *postRepository) FindById(id string) (sqlc.Post, error) {
	org, err := r.db.FindPostById(context.Background(), id)
	return org, err
}

func (r *postRepository) ListAll() ([]sqlc.Post, error) {
	return r.db.ListPosts(context.Background())
}

func (r *postRepository) Delete(id string) error {
	return r.db.DeletePost(context.Background(), id)
}

func (r *postRepository) FindByOrganization(orgId string) ([]sqlc.Post, error) {
	return r.db.FindPostByOrganization(context.Background(), orgId)
}
