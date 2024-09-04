package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type SeederRepository interface {
	AddCategory(sqlc.AddCategoryParams) (sql.Result, error)
}

type seederRepository struct {
	db *sqlc.Queries
}

func NewSeederRepository(db *sqlc.Queries) SeederRepository {
	return &seederRepository{db}
}

func (r *seederRepository) AddCategory(input sqlc.AddCategoryParams) (sql.Result, error) {
	return r.db.AddCategory(context.Background(), input)
}
