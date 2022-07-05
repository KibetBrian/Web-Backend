package api

import (
	"database/sql"
	"log"

	db "github.com/KibetBrian/backend/db/sqlc"
	"github.com/KibetBrian/backend/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	configs, err :=utils.LoadConfig("./")
	if err != nil{
		log.Fatalf("\nFailed to load enviroment variables. Err %v", err)
	}
	conn, err := sql.Open(configs.DbDriver, configs.DbSource)
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
	err := s.router.Run(address)
	if err != nil {
		log.Fatalf("\nFailed to start the server. Error: %v", err)
	}
}


