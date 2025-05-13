.PHONY: build run test migrate-up migrate-down swagger

# Build aplikasi
build:
	go build -o bin/api cmd/api/main.go

# Jalankan aplikasi
run:
	go run cmd/api/main.go

# Jalankan test
test:
	go test -v ./...

# Jalankan migrasi database
migrate-up:
	migrate -path migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" up

# Rollback migrasi database
migrate-down:
	migrate -path migrations -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" down

# Generate Swagger documentation
swagger:
	swag init -g cmd/api/main.go -o docs

# Install dependencies
deps:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Clean build files
clean:
	rm -rf bin/
	rm -rf docs/ 