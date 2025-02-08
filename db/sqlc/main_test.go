package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jpmoraess/appointment-api/pkg/utils"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	cfg, err := utils.LoadConfig("../..")
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
