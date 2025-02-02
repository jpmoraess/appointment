**Install migrate:** $ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.0/migrate.windows-amd64.zip

Ex: migrate create -ext sql -seq init_schema

**Install sqlc:** go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

Run: sqlc init

docker run --name appointment-system --network appointment-system-network -p 8080:8080 -e GIN_MODE=release 
-e DB_SOURCE=postgresql://postgres:postgres@postgres-database:5432/appointment_system?sslmode=disable appointment-api:latest