createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple-bank

dropdb:
	docker exec -it postgres dropdb --username=postgres simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb mi grateup migratedown sqlc
