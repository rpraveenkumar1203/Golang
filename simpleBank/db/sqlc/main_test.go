package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Initialize the connection to the database
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Unable to connect to the database:", err)
	}

	// Check if the connection is alive
	err = testDB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	// Initialize testQueries
	testQueries = New(testDB)

	// Run the tests
	os.Exit(m.Run())
}
