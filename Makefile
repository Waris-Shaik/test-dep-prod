build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))
	
migrate-up:
	@echo "Applying database migrations..."
	@go run cmd/migrate/main.go up

migrate-down:
	@echo "Rolling back database migrations..."
	@go run cmd/migrate/main.go down

migrate-force:
	@echo "Forcing migration version..."
	@go run cmd/migrate/main.go force $(version)


# /*CREATE TABLE IF NOT EXISTS users (
#     `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
#     `first_name` VARCHAR(100) NOT NULL,
#     `last_name` VARCHAR(100) NOT NULL,
#     `username` VARCHAR(100) NOT NULL,
#     `email` TEXT UNIQUE NOT NULL,
#     `password` TEXT NOT NULL,
#     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
# );
# */