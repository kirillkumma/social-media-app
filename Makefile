postgres:
	docker run -d --name social-media-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass postgres:15-alpine

createdb:
	docker exec -it social-media-db createdb --username=root --owner=root social-media

dropdb:
	docker exec -it social-media-db dropdb --username=root social-media

migrateup:
	migrate -path migration -database "postgresql://root:pass@localhost:5432/social-media?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:pass@localhost:5432/social-media?sslmode=disable" -verbose down

run:
	go run cmd/app/main.go

.PHONY: postgres createdb dropdb migrateup migratedown run