package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rpraveenkumar/Golang/db/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := utils.LoadConfig("../..")

	if err != nil {
		log.Fatal("unable to load config folder", err)
	}

	// Initialize the connection to the database
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
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
