package telegramClient

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	mock_repository "tcms-auth/internal/testing/database/repository"
	mock_telegram "tcms-auth/internal/testing/telegram"
	"tcms-auth/pkg/telegram"
)

func TestTelegramClient_Authorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const phone = "999999"
	request := &telegram.LoginMessage{Phone: phone}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Login(gomock.Eq(ctx), gomock.Eq(request))

	userRepo := mock_repository.NewMockUserRepository(ctrl)

	tg := newTelegram(client, userRepo)
	err := tg.Authorization(ctx, phone)
	assert.Nil(t, err)
}

func TestTelegramClient_AuthSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const code = "999999"
	const userId = 1
	const session = "session"

	resp := &telegram.SignResult{Session: session}

	request := &telegram.SignMessage{Code: code}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Sign(gomock.Eq(ctx), gomock.Eq(request)).Return(resp, nil)

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().UpdateTelegramAccessKey(gomock.Eq(userId), gomock.Eq(session)).Return(nil)

	tg := newTelegram(client, userRepo)
	err := tg.AuthSignIn(ctx, userId, code)
	assert.Nil(t, err)
}

func TestTelegramClient_AuthSignIn_telegramError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const code = "999999"
	const userId = 1
	resErr := fmt.Errorf("some error")

	request := &telegram.SignMessage{Code: code}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Sign(gomock.Eq(ctx), gomock.Eq(request)).Return(nil, resErr)

	userRepo := mock_repository.NewMockUserRepository(ctrl)

	tg := newTelegram(client, userRepo)
	err := tg.AuthSignIn(ctx, userId, code)
	assert.Equal(t, resErr, err)
}

func TestTelegramClient_AuthSignIn_repoErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const code = "999999"
	const userId = 1
	const session = "session"
	resErr := fmt.Errorf("some error")

	resp := &telegram.SignResult{Session: session}

	request := &telegram.SignMessage{Code: code}
	ctx := context.Background()

	client := mock_telegram.NewMockTelegramClient(ctrl)
	client.EXPECT().Sign(gomock.Eq(ctx), gomock.Eq(request)).Return(resp, nil)

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	userRepo.EXPECT().UpdateTelegramAccessKey(gomock.Eq(userId), gomock.Eq(session)).Return(resErr)

	tg := newTelegram(client, userRepo)
	err := tg.AuthSignIn(ctx, userId, code)
	assert.Equal(t, resErr, err)
}

func TestGetTelegram(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conn, err := grpc.Dial("host", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err)

	userRepo := mock_repository.NewMockUserRepository(ctrl)

	_ = GetTelegram(conn, userRepo)
}
