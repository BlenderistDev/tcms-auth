test:
	@go test ./... -race -count=1 -cover -coverprofile=coverage.txt && go tool cover -func=coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm coverage.txt

format:
	@go fmt ./...

run:
	@go run cmd/main.go

build:
	@mkdir -p ./bin
	@go build -o bin ./cmd

gen-auth:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/auth.proto
	@mv proto/auth*.go pkg/auth/

gen-session-repository-mock:
	@mockgen -source=internal/database/repository/sessionRepository.go -destination=internal/testing/database/repository/sessionRepository.go

gen-user-repository-mock:
	@mockgen -source=internal/database/repository/userRepository.go -destination=internal/testing/database/repository/userRepository.go

gen-password-generator-mock:
	@mockgen -source=internal/password/hash.go -destination=internal/testing/password/hash.go

gen-telegram:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/telegram.proto
	@mv proto/telegram*.go pkg/telegram/

gen-telegram-mock:
	@mockgen -source=pkg/telegram/telegram_grpc.pb.go -destination=internal/testing/telegram/telegram_grpc.pb.go
