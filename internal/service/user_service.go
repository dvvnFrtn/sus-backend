package service

import "sus-backend/internal/repository"

type UserService interface {
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}
