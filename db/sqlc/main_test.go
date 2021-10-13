package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/cart-buddy?sslmode=disable"

// testQueries Establish connection to the Queries object
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Establish a simple database connection
	testDB, err = sql.Open(dbDriver, dbSource)

	// Check for errors in opening DB connection
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	// Define the object and assign the DB connection
	testQueries = New(testDB)

	os.Exit(m.Run())
}
