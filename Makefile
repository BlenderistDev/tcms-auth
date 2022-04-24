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
