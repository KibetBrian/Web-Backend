package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)
var testQueries *Queries

const (
	dbDriver="postgres"
	dbSource="postgresql://brian:brian@127.0.0.1:5432/election?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver,dbSource);
	if err != nil{
		log.Fatalf("\n Database connection failed: %v", err)
	}
	testQueries=New(conn)
	os.Exit(m.Run())
}