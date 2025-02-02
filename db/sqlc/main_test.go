package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jpmoraess/appointment-api/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	cfg, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	conn, err := pgx.Connect(context.Background(), cfg.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
