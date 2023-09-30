DB_URL=postgresql://root:secret@localhost:5432/db_auth?sslmode=disable

postgres:
	docker run --name pg-local -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg-local createdb --username=root --owner=root db_auth

dropdb:
	docker exec -it pg-local dropdb db_auth

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

re-db: dropdb createdb migrateup

sqlc-win:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

run:
	go run ./cmd/

build-image:
	docker build -t efner/auth-microservice:1.0 .

.PHONY: postgres createdb migrateup migrateup1 migratedown migratedown1 new_migration re-db run build-image