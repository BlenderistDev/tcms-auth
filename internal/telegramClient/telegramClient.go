package telegramClient

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"tcms-auth/internal/database/repository"
	"tcms-auth/pkg/telegram"
)

type TelegramClient interface {
	Authorization(ctx context.Context, phone string) error
	AuthSignIn(ctx context.Context, code string) error
}

type telegramClient struct {
	telegram telegram.TelegramClient
	userRepo repository.UserRepository
}

func newTelegram(tg telegram.TelegramClient, userRepo repository.UserRepository) TelegramClient {
	return &telegramClient{
		telegram: tg,
		userRepo: userRepo,
	}
}

// GetTelegram create new telegram client
func GetTelegram(conn *grpc.ClientConn, userRepo repository.UserRepository) TelegramClient {
	tg := telegram.NewTelegramClient(conn)
	return newTelegram(tg, userRepo)
}

// Authorization request for authorization in telegram client
func (t *telegramClient) Authorization(ctx context.Context, phone string) error {
	_, err := t.telegram.Login(ctx, &telegram.LoginMessage{Phone: phone})
	return err
}

// AuthSignIn request for sign in telegram client with auth code
func (t *telegramClient) AuthSignIn(ctx context.Context, code string) error {
	res, err := t.telegram.Sign(ctx, &telegram.SignMessage{Code: code})
	if err != nil {
		return err
	}
	md, _ := metadata.FromIncomingContext(ctx)
	userId := md.Get("userId")[0]
	if len(userId) == 0 {
		return errors.New("no user id")
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}

	err = t.userRepo.UpdateTelegramAccessKey(id, res.Session)
	if err != nil {
		return err
	}
	res.GetSession()
	return err
}
