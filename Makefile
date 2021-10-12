postgres:
	docker run --name cart-buddy-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it cart-buddy-postgres createdb --username=root --owner=root cart-buddy

dropdb:
	docker exec -it cart-buddy-postgres dropdb cart-buddy

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/cart-buddy?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/cart-buddy?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc

