package service

import (
	"context"
	"errors"

	"tcms-auth/internal/database/model"
	"tcms-auth/internal/database/repository"
	"tcms-auth/internal/password"
	"tcms-auth/pkg/auth"
)

// ErrNoUser user not found
var ErrNoUser = errors.New("no user")

// ErrWrongPassword password is incorrect
var ErrWrongPassword = errors.New("wrong password")

// ErrNoAuth user is not authed
var ErrNoAuth = errors.New("user is not authed")

type AuthGrpcService struct {
	UserRepo          repository.UserRepository
	SessionRepo       repository.SessionRepository
	PasswordGenerator password.Generator
	auth.UnimplementedTcmsAuthServer
}

// Register grpc endpoint for user registration
func (s *AuthGrpcService) Register(_ context.Context, user *auth.AuthData) (*auth.Result, error) {
	pass, err := s.PasswordGenerator.Generate(user.GetPassword())
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

// Login grpc endpoint for user login
func (s *AuthGrpcService) Login(_ context.Context, loginData *auth.AuthData) (*auth.LoginResult, error) {
	user, err := s.UserRepo.GetUser(loginData.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrNoUser
	}
	if s.PasswordGenerator.Compare(user.Password, loginData.Password) != nil {
		return nil, ErrWrongPassword
	}

	token, err := s.SessionRepo.Create(user.Id)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResult{Token: token}, err
}

// CheckAuth grpc endpoint for authorization check
func (s *AuthGrpcService) CheckAuth(_ context.Context, authData *auth.LoginResult) (*auth.CheckAuthResult, error) {
	user, err := s.SessionRepo.GetUser(authData.Token)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrNoAuth
	}

	return &auth.CheckAuthResult{
		Username:      user.Username,
		TelegramToken: user.TelegramAccessKey,
	}, nil
}
