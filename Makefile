up:
	docker compose up -d

down:
	docker compose down

gen:
	docker compose exec app go generate ./...

test:
	docker compose exec app go test ./...

lint:
	docker compose exec app go vet ./...

tidy:
	docker compose exec app go mod tidy

sqlc-gen:
	docker compose exec app sqlc generate
