// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: category.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addCategory = `-- name: AddCategory :execresult
INSERT INTO categories (id, category_name)
VALUES (?, ?)
`

type AddCategoryParams struct {
	ID           string
	CategoryName string
}

func (q *Queries) AddCategory(ctx context.Context, arg AddCategoryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addCategory, arg.ID, arg.CategoryName)
}

const createUserCategory = `-- name: CreateUserCategory :execresult
INSERT INTO user_categories (category_id, user_id)
VALUES (?, ?)
`

type CreateUserCategoryParams struct {
	CategoryID string
	UserID     string
}

func (q *Queries) CreateUserCategory(ctx context.Context, arg CreateUserCategoryParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUserCategory, arg.CategoryID, arg.UserID)
}

const userCategoryExists = `-- name: UserCategoryExists :one
SELECT COUNT(1) from user_categories
WHERE category_id = ? AND user_id = ?
`

type UserCategoryExistsParams struct {
	CategoryID string
	UserID     string
}

func (q *Queries) UserCategoryExists(ctx context.Context, arg UserCategoryExistsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, userCategoryExists, arg.CategoryID, arg.UserID)
	var count int64
	err := row.Scan(&count)
	return count, err
}