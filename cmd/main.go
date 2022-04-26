package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"tcms-auth/internal/database"
	"tcms-auth/internal/database/repository"
	"tcms-auth/internal/service"
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

	if err := startGrpcServer(userRepo, sessionRepo); err != nil {
		log.Fatal(err)
	}

}

func startGrpcServer(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) error {
	addr := os.Getenv("AUTH_GRPC_HOST")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	auth.RegisterTcmsAuthServer(s, &service.AuthGrpcService{
		UserRepo:    userRepo,
		SessionRepo: sessionRepo,
	})

	return s.Serve(lis)
}
