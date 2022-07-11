package api

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)


func NewServer() *Server{
	server := &Server{db: DBQueries(), router: gin.Default()}
	server.router.Use(cors.Default())

	server.router.POST("/user/create", server.RegisterUser)
	server.router.GET("/user/get", server.GetUser)
	server.router.POST("/user/login", server.Login)
	server.router.GET("/voter/chain", server.GetChainVoters)
	server.router.POST("/voter/register", server.RegisterVoter)
	server.router.POST("/candidate/register", server.RegisterContestant)
	server.router.GET("/voter/total", server.GetTotalVotersNumb)
	server.router.GET("/user/total", server.GetTotalUsersNum)
	server.router.GET("/voter/voted", server.GetTotalVotedVoters)
	server.router.GET("/candidates", server.GetAllContestants)

	return server;
}
