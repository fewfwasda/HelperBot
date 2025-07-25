package services

import (
	services "HelperBot/internal/services/users"
	"HelperBot/internal/user"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user user.User) (string, error)
}

type Service struct {
	UserService UserService
}

func NewService() *Service {
	return &Service{
		UserService: services.NewUserService(),
	}
}
