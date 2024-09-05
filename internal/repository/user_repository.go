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
	AddUserCategory(sqlc.CreateUserCategoryParams) (sql.Result, error)
	UserCategoryExists(sqlc.UserCategoryExistsParams) (int64, error)
	GetOrganizer(string) (sqlc.Organizer, error)
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

func (r *userRepository) AddUserCategory(input sqlc.CreateUserCategoryParams) (sql.Result, error) {
	return r.db.CreateUserCategory(context.Background(), input)
}

func (r *userRepository) UserCategoryExists(input sqlc.UserCategoryExistsParams) (int64, error) {
	return r.db.UserCategoryExists(context.Background(), input)
}

func (r *userRepository) GetOrganizer(id string) (sqlc.Organizer, error) {
	return r.db.GetOrganizer(context.Background(), id)
}
