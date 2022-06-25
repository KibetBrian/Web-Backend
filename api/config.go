package api

import (
	"database/sql"
	"log"

	db "github.com/KibetBrian/backend/db/sqlc"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

const (
	dbDriver="postgres"
	dbSource="postgresql://brian:brian@127.0.0.1:5432/election?sslmode=disable"
)

type Server struct{
	router *gin.Engine
	db *db.Queries
}

func Conn() *sql.DB{
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("\n Database Connection Failed. Error : %v\n", err)
	}
	return conn
}

func DBQueries() *db.Queries{
	queries := db.New(Conn())
	return queries
}

func (s *Server) Start(address string){
	err := s.router.Run()
	if err != nil {
		log.Fatalf("\nFailed to start the server. Error: %v", err)
	}
}


