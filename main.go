package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/appointment-api/api"
	db "github.com/jpmoraess/appointment-api/db/sqlc"
	"log"
)

const (
	connString    = "postgresql://postgres:postgres@localhost:5432/appointment_system?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal("cannot parse connection:", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal("cannot create connection to database: ", err)
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer pool.Close()

	store := db.NewStore(pool)

	server := api.NewServer(store)

	err = server.StartServer(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
