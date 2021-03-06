package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tcms-auth/internal/database/model"
	"tcms-auth/internal/errors"
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

	resErr := fmt.Errorf("some error")

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

	resErr := fmt.Errorf("some error")

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

func TestAuthGrpcService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const userId = 1
	const username = "name"
	const password = "password"
	const passwordHash = "passwordHash"
	const token = "token"

	user := &model.User{
		Id:       userId,
		Username: username,
		Password: passwordHash,
	}

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetUser(gomock.Eq(username)).Return(user, nil)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)
	sessionRepo.EXPECT().Create(gomock.Eq(userId)).Return(token, nil)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)
	passwordGenerator.EXPECT().Compare(gomock.Eq(passwordHash), gomock.Eq(password)).Return(nil)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Login(context.Background(), authData)
	assert.Nil(t, err)
	assert.Equal(t, token, res.Token)
}

func TestAuthGrpcService_Login_userRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const username = "name"
	const password = "password"

	resErr := fmt.Errorf("some error")

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetUser(gomock.Eq(username)).Return(nil, resErr)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Login(context.Background(), authData)
	assert.Nil(t, res)
	assert.Equal(t, resErr, err)
}

func TestAuthGrpcService_Login_noUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const username = "name"
	const password = "password"

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetUser(gomock.Eq(username)).Return(nil, nil)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Login(context.Background(), authData)
	assert.Nil(t, res)
	assert.Equal(t, errors.ErrNoUser, err)
}

func TestAuthGrpcService_Login_wrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const userId = 1
	const username = "name"
	const password = "password"
	const passwordHash = "passwordHash"

	resErr := fmt.Errorf("some error")

	user := &model.User{
		Id:       userId,
		Username: username,
		Password: passwordHash,
	}

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetUser(gomock.Eq(username)).Return(user, nil)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)
	passwordGenerator.EXPECT().Compare(gomock.Eq(passwordHash), gomock.Eq(password)).Return(resErr)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Login(context.Background(), authData)
	assert.Nil(t, res)
	assert.Equal(t, errors.ErrWrongPassword, err)
}

func TestAuthGrpcService_Login_sessionRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const userId = 1
	const username = "name"
	const password = "password"
	const passwordHash = "passwordHash"

	resErr := fmt.Errorf("some error")

	user := &model.User{
		Id:       userId,
		Username: username,
		Password: passwordHash,
	}

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetUser(gomock.Eq(username)).Return(user, nil)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)
	sessionRepo.EXPECT().Create(gomock.Eq(userId)).Return("", resErr)

	passwordGenerator := mock_password.NewMockGenerator(ctrl)
	passwordGenerator.EXPECT().Compare(gomock.Eq(passwordHash), gomock.Eq(password)).Return(nil)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	authData := &auth.AuthData{
		Username: username,
		Password: password,
	}

	res, err := service.Login(context.Background(), authData)
	assert.Nil(t, res)
	assert.Equal(t, err, resErr)
}

func TestAuthGrpcRepo_CheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const username = "name"
	const token = "token"
	const telegramAccessKey = "tg_access_key"

	user := &model.AuthUser{
		Username:          username,
		TelegramAccessKey: telegramAccessKey,
	}

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	passwordGenerator := mock_password.NewMockGenerator(ctrl)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)
	sessionRepo.EXPECT().GetUser(gomock.Eq(token)).Return(user, nil)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	resData := &auth.CheckAuthResult{
		Username:      username,
		TelegramToken: telegramAccessKey,
	}

	income := &auth.LoginResult{Token: token}

	res, err := service.CheckAuth(context.Background(), income)
	assert.Nil(t, err)
	assert.Equal(t, resData, res)
}

func TestAuthGrpcRepo_CheckAuth_sessionRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const token = "token"

	resErr := fmt.Errorf("some error")

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	passwordGenerator := mock_password.NewMockGenerator(ctrl)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)
	sessionRepo.EXPECT().GetUser(gomock.Eq(token)).Return(nil, resErr)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	income := &auth.LoginResult{Token: token}

	res, err := service.CheckAuth(context.Background(), income)
	assert.Nil(t, res)
	assert.Equal(t, resErr, err)
}

func TestAuthGrpcRepo_CheckAuth_noAuthError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const token = "token"

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	passwordGenerator := mock_password.NewMockGenerator(ctrl)

	sessionRepo := mock_repository.NewMockSessionRepository(ctrl)
	sessionRepo.EXPECT().GetUser(gomock.Eq(token)).Return(nil, nil)

	service := AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: passwordGenerator,
	}

	income := &auth.LoginResult{Token: token}

	res, err := service.CheckAuth(context.Background(), income)
	assert.Nil(t, res)
	assert.Equal(t, errors.ErrNoAuth, err)
}
