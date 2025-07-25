package services

import (
	"HelperBot/internal/user"
	"context"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(ctx context.Context, user user.User) (string, error) {
	panic("asda")
}
