default: run
	
sqlc-gen:
	@sqlc generate

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

run: sqlc-gen
	@go run ./cmd/api/main.go

build: vet
	@go build -o api ./cmd/api/main.go


