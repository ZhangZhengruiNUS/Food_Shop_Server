postgres:
	docker run --name food_shop_product -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it food_shop_product createdb --username=root --owner=root food_shop_product

dropdb:
	docker exec -it food_shop_product dropdb --username=root --owner=root food_shop_product

migrateup:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/food_shop_product?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/food_shop_product?sslmode=disable" -verbose down

migratecreate:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateforce:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/food_shop_product?sslmode=disable" force $(version)

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -cover ./db/sqlc/ ./handlers

mock:
	mockgen -package mockdb -destination db/mock/store.go Food_Shop_Server/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migratecreate migrateforce sqlc server mock test

