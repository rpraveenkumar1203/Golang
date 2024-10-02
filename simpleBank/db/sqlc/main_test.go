package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	drivername     = "postgres"
	databasesource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(drivername, databasesource)

	if err != nil {
		log.Fatal("unable to connect to database due to this ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())

}
