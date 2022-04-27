package service

import (
	"context"
	"fmt"

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

func (s *AuthGrpcService) Login(_ context.Context, loginData *auth.AuthData) (*auth.LoginResult, error) {
	user, err := s.UserRepo.GetUser(loginData.Username)
	if err != nil {
		return nil, err
	}
	if password.Compare(user.Password, loginData.Password) != nil {
		return nil, fmt.Errorf("wrong password")
	}

	token, err := s.SessionRepo.Create(user.Id)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResult{Token: token}, err
}

func (s *AuthGrpcService) CheckAuth(_ context.Context, authData *auth.LoginResult) (*auth.CheckAuthResult, error) {
	user, err := s.SessionRepo.GetUser(authData.Token)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user is not authed")
	}

	return &auth.CheckAuthResult{
		Username:      user.Username,
		TelegramToken: user.TelegramAccessKey,
	}, nil
}
