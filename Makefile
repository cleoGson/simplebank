postgres:
	docker run --name postgres  -p 5452:5432 -ePOSTGRES_USER=root  -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_banks
dropdb:
	docker exec -it postgres dropdb simple_banks
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5452/simple_banks?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5452/simple_banks?sslmode=disable" -verbose down
migrateup1:
	migrate -path db/migrations -database db/migrations -database "postgresql://root:secret@localhost:5452/simple_banks?sslmode=disable" -verbose up 1 force 1
sqlc:
	sqlc generate
server:
	go run main.go

.PHONY:postgres createdb dropdb migratedown migrateup sqlc server migrateup1
