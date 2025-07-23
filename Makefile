.PHONY: createdb dropdb migrateup migratedown postgres sqlc test mock server proto

DB_CONTAINER=d3d89de55d86
DB_NAME=simple_bank
DB_URL=postgres://root:ZsbUtgSTLI68JdgOR4jq@simple-bank.chwcewcwcq07.ap-southeast-2.rds.amazonaws.com:5432/$(DB_NAME)

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

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/dorasaicu12/simplebank/db/sqlc Store

server:
	go run main.go
	
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto