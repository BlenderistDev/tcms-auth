package service

import (
	"context"

	"tcms-auth/internal/database/model"
	"tcms-auth/internal/database/repository"
	"tcms-auth/internal/password"
	"tcms-auth/pkg/auth"
)

type AuthGrpcService struct {
	UserRepo    repository.UserRepository
	SessionRepo repository.SessionRepository
	auth.UnimplementedTcmsAuthServer
}

// Register grpc endpoint for user registration
func (s *AuthGrpcService) Register(_ context.Context, user *auth.AuthData) (*auth.Result, error) {
	pass, err := password.Generate(user.GetPassword())
	if err != nil {
		return nil, err
	}
	if err := s.UserRepo.Create(&model.User{
		Username:          user.GetUsername(),
		Password:          pass,
		TelegramAccessKey: "",
	}); err != nil {
		return &auth.Result{}, err
	}
	return &auth.Result{Success: true}, nil
}
