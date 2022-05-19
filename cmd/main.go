package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tcms-auth/internal/database"
	"tcms-auth/internal/database/repository"
	"tcms-auth/internal/password"
	"tcms-auth/internal/service"
	"tcms-auth/internal/telegramClient"
	"tcms-auth/pkg/auth"
)

func main() {

	// Load values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.GetConnection(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	conn, err := grpc.Dial(os.Getenv("TELEGRAM_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	tg := telegramClient.GetTelegram(conn, userRepo)

	if err := startGrpcServer(userRepo, sessionRepo, tg); err != nil {
		log.Fatal(err)
	}

}

func startGrpcServer(userRepo repository.UserRepository, sessionRepo repository.SessionRepository, telegramClient telegramClient.TelegramClient) error {
	addr := os.Getenv("AUTH_GRPC_HOST")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	auth.RegisterTcmsAuthServer(s, &service.AuthGrpcService{
		UserRepo:          userRepo,
		SessionRepo:       sessionRepo,
		PasswordGenerator: password.NewGenerator(),
		TelegramClient: telegramClient,
	})

	return s.Serve(lis)
}
