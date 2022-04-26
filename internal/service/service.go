package service

import (
	"context"

	"tcms-auth/internal/database/model"
	"tcms-auth/internal/database/repository"
	"tcms-auth/pkg/auth"
)

type AuthGrpcService struct {
	UserRepo repository.UserRepository
	auth.UnimplementedTcmsAuthServer
}

func (s AuthGrpcService) Register(_ context.Context, user *auth.RegisterMessage) (*auth.Result, error) {
	err := s.UserRepo.Create(&model.User{
		Username:          user.GetUsername(),
		Password:          user.GetPassword(),
		TelegramAccessKey: "",
	})
	if err != nil {
		return &auth.Result{}, err
	}
	return &auth.Result{Success: true}, nil
}
