package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type SeederRepository interface {
	AddCategory(sqlc.AddCategoryParams) (sql.Result, error)
	AddCategoryGroup(string) (sql.Result, error)
	AddUser(sqlc.AddUserParams) (sql.Result, error)
	AddPost(sqlc.AddPostParams) (sql.Result, error)
	AddOrganization(sqlc.AddOrganizationParams) (sql.Result, error)
	CategoryExists(sqlc.CategoryExistsParams) (int64, error)
	CategoryGroupExists(string) (int64, error)
	GetGroupIDByName(string) (int32, error)
}

type seederRepository struct {
	db *sqlc.Queries
}

func NewSeederRepository(db *sqlc.Queries) SeederRepository {
	return &seederRepository{db}
}

func (r *seederRepository) CategoryExists(arg sqlc.CategoryExistsParams) (int64, error) {
	return r.db.CategoryExists(context.Background(), arg)
}

func (r *seederRepository) CategoryGroupExists(name string) (int64, error) {
	return r.db.CategoryGroupExists(context.Background(), name)
}

func (r *seederRepository) AddCategory(input sqlc.AddCategoryParams) (sql.Result, error) {
	return r.db.AddCategory(context.Background(), input)
}

func (r *seederRepository) AddCategoryGroup(name string) (sql.Result, error) {
	return r.db.AddCategoryGroup(context.Background(), name)
}

func (r *seederRepository) GetGroupIDByName(name string) (int32, error) {
	return r.db.GetCategoryGroupIDByName(context.Background(), name)
}

func (r *seederRepository) AddUser(arg sqlc.AddUserParams) (sql.Result, error) {
	return r.db.AddUser(context.Background(), arg)
}

func (r *seederRepository) AddPost(arg sqlc.AddPostParams) (sql.Result, error) {
	return r.db.AddPost(context.Background(), arg)
}

func (r *seederRepository) AddOrganization(arg sqlc.AddOrganizationParams) (sql.Result, error) {
	return r.db.AddOrganization(context.Background(), arg)
}
