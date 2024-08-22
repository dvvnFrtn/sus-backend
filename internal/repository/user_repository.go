package repository

import (
	"context"
	"database/sql"
	"sus-backend/internal/db/sqlc"
)

type UserRepository interface {
	EmailExists(string) (int64, error)
	FindByEmail(string) (sqlc.User, error)
	CreateUser(sqlc.AddUserParams) (sql.Result, error)
}

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) EmailExists(email string) (int64, error) {
	return r.db.EmailExists(context.Background(), email)
}

func (r *userRepository) FindByEmail(email string) (sqlc.User, error) {
	return r.db.FindByEmail(context.Background(), email)
}

func (r *userRepository) CreateUser(input sqlc.AddUserParams) (sql.Result, error) {
	return r.db.AddUser(context.Background(), input)
}
