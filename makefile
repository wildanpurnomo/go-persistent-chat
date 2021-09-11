postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres chat_db

dropdb:
	docker exec -it postgres dropdb chat_db

migrateup:
	migrate -path server/db/migrations -database postgres://postgres:password@127.0.0.1:5432/chat_db?sslmode=disable&timezone=Asia/Jakarta up

migratedown:
	migrate -path server/db/migrations -database postgres://postgres:password@127.0.0.1:5432/chat_db?sslmode=disable&timezone=Asia/Jakarta down

.PHONY: postgres createdb dropdb migrateup migratedown
