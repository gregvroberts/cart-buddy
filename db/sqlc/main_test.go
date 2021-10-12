package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/cart-buddy?sslmode=disable"

// testQueries Establish connection to the Queries object
var testQueries *Queries

func TestMain(m *testing.M) {

	// Establish a simple database connection
	conn, err := sql.Open(dbDriver, dbSource)

	// Check for errors in opening DB connection
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	// Define the object and assign the DB connection
	testQueries = New(conn)

	os.Exit(m.Run())
}
