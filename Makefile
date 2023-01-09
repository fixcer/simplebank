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

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/fixcer/simplebank/db/sqlc Store

.PHONY: createdb dropdb mi grateup migratedown sqlc test server mock
