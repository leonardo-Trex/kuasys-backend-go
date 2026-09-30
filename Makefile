default: run
	
fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

run:
	@go run ./cmd/api/main.go

build: vet
	@go build -o api ./cmd/api/main.go


