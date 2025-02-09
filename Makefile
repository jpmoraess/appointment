postgres:
	docker run --name postgres-database --network appointment-system-network -p 5432:5432 -e POSTGRES_USERNAME=postgres -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres-database createdb --username=postgres --owner=postgres appointment_system

dropdb:
	docker exec -it postgres-database dropdb --username=postgres --owner=postgres appointment_system

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/appointment_system?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/appointment_system?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/jpmoraess/appointment-api/db/sqlc Store

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock proto