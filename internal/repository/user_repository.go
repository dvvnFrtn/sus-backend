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
	FindUserByID(string) (sqlc.User, error)
	UpdateUser(sqlc.UpdateUserByIDParams) (sql.Result, error)
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
	return r.db.FindUserByEmail(context.Background(), email)
}

func (r *userRepository) CreateUser(input sqlc.AddUserParams) (sql.Result, error) {
	return r.db.AddUser(context.Background(), input)
}

func (r *userRepository) FindUserByID(id string) (sqlc.User, error) {
	return r.db.FindUserByID(context.Background(), id)
}

func (r *userRepository) UpdateUser(input sqlc.UpdateUserByIDParams) (sql.Result, error) {
	return r.db.UpdateUserByID(context.Background(), input)
}
