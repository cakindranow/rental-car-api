DB_URL=postgresql://root:secret@localhost:5432/food_order?sslmode=disable

postgres:
	docker run --name postgres_car_order -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres_car_order createdb --username=root --owner=root cars

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down