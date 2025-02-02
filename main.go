package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jpmoraess/appointment-api/api"
	db "github.com/jpmoraess/appointment-api/db/sqlc"
	"github.com/jpmoraess/appointment-api/util"
	"log"
)

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConfig, err := pgxpool.ParseConfig(cfg.DbSource)
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

	err = server.StartServer(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
