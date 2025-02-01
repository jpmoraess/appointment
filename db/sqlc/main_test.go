package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"
)

const (
	connString = "postgresql://postgres:postgres@localhost:5432/appointment_system?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
