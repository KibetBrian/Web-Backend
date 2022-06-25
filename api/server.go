package api

import (
	"github.com/gin-gonic/gin"
)


func NewServer() *Server{
	server := &Server{db: DBQueries(), router: gin.Default()}	

	server.router.POST("/user/create", server.RegisterUser)
	return server;
}
