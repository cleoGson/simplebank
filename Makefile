postgres:
	docker run --name postgres -ePOSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_banks
dropdb:
	docker exec -it postgres dropdb simple_banks
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5453/simple_banks?ssl_mode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5453/simple_banks?ssl_mode=disable" -verbose down
sqlc:
	sqlc generate
.PHONY:postgres createdb dropdb migratedown migrateup sqlc
