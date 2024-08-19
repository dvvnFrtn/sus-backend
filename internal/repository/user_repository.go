package repository

import "sus-backend/internal/db/sqlc"

type UserRepository interface {
}

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{db}
}
