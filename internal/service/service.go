package service

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
	"tcms-auth/internal/database/model"
	"tcms-auth/internal/database/repository"
	"tcms-auth/internal/errors"
	"tcms-auth/internal/password"
	"tcms-auth/internal/telegramClient"
	"tcms-auth/pkg/auth"
)

type AuthGrpcService struct {
	UserRepo          repository.UserRepository
	SessionRepo       repository.SessionRepository
	PasswordGenerator password.Generator
	TelegramClient    telegramClient.TelegramClient
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
		return nil, errors.ErrNoUser
	}
	if s.PasswordGenerator.Compare(user.Password, loginData.Password) != nil {
		return nil, errors.ErrWrongPassword
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
		return nil, errors.ErrNoAuth
	}

	return &auth.CheckAuthResult{
		Username:      user.Username,
		TelegramToken: user.TelegramAccessKey,
	}, nil
}

// TelegramAuth sends code for telegram auth
func (s *AuthGrpcService) TelegramAuth(ctx context.Context, authData *auth.TelegramAuthRequest) (*auth.TelegramAuthResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "userId", strconv.Itoa(int(authData.UserId)))
	err := s.TelegramClient.Authorization(ctx, authData.Phone)
	if err != nil {
		return nil, err
	}

	return &auth.TelegramAuthResponse{Success: true}, nil
}

// TelegramSign sign in telegram
func (s *AuthGrpcService) TelegramSign(ctx context.Context, signData *auth.TelegramSignRequest) (*auth.TelegramAuthResponse, error) {
	err := s.TelegramClient.AuthSignIn(ctx, int(signData.UserId), signData.Code)
	if err != nil {
		return nil, err
	}

	return &auth.TelegramAuthResponse{Success: true}, nil
}
