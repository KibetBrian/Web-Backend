package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/KibetBrian/backend/utils"
	_ "github.com/lib/pq"
)
var testQueries *Queries

func TestMain(m *testing.M) {
	configs, err := utils.LoadConfig("../../")
	if err != nil{
		log.Fatalf("Failed to load config. Error: %v", err)
	}
	conn, err := sql.Open(configs.DbDriver,configs.DbSource);
	if err != nil{
		log.Fatalf("\n Database connection failed: %v", err)
	}
	testQueries=New(conn)
	os.Exit(m.Run())
}