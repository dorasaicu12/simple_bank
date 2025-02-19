.PHONY: createdb dropdb migrateup migratedown postgres sqlc test

DB_CONTAINER=d3d89de55d86
DB_NAME=simple_bank
DB_URL=postgres://root:admin@localhost:5432/$(DB_NAME)?sslmode=disable

createdb:
	docker exec -it $(DB_CONTAINER) createdb --username=root --owner=root $(DB_NAME)

dropdb:
	docker exec -it $(DB_CONTAINER) dropdb $(DB_NAME)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

postgres:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

test:
	go test -v -cover ./...

sqlc:
	sqlc generate