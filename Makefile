build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

gen:
	docker compose exec app go generate ./...

test:
	cd app && go test -v ./...

lint:
	docker compose exec app go vet ./...

tidy:
	docker compose exec app go mod tidy

sqlc-gen:
	docker compose exec app sqlc generate

swag:
	docker compose exec app swag init --output ./docs/swagger
