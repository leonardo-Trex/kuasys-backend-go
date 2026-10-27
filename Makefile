default: run
	
sqlc-gen:
	@sqlc generate

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

run: sqlc-gen
	@go run ./cmd/api/main.go

build: sqlc-gen vet
	@go build -o api ./cmd/api/main.go


