package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tcms-auth/internal/database/model"
	mock_repository "tcms-auth/internal/testing/database/repository"
	mock_password "tcms-auth/internal/testing/password"
	"tcms-auth/pkg/auth"
)

func TestAuthGrpcService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const username = "name"
	const password = "password"
	const passwordHash = "passwordHash"

	user := &model.User{
		Username: username,
		Password: passwordHash,
	}

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().Create(gomock.Eq(user)).Return(nil)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)
	passwordGenerator.EXPECT().Generate(gomock.Eq(password)).Return(passwordHash, nil)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Register(context.Background(), authData)
	assert.Nil(t, err)
	assert.True(t, res.Success)
}

func TestAuthGrpcService_Register_passwordGenerateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const username = "name"
	const password = "password"

	resErr := errors.New("some error")

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)
	passwordGenerator := mock_password.NewMockGenerator(ctrl)
	passwordGenerator.EXPECT().Generate(gomock.Eq(password)).Return("", resErr)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Register(context.Background(), authData)
	assert.Nil(t, res)
	assert.Equal(t, resErr, err)
}

func TestAuthGrpcService_Register_repositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const username = "name"
	const password = "password"
	const passwordHash = "passwordHash"

	resErr := errors.New("some error")

	user := &model.User{
		Username: username,
		Password: passwordHash,
	}

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().Create(gomock.Eq(user)).Return(resErr)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)
	passwordGenerator.EXPECT().Generate(gomock.Eq(password)).Return(passwordHash, nil)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Register(context.Background(), authData)
	assert.False(t, res.GetSuccess())
	assert.Equal(t, resErr, err)
}
