package api

import (
	"context"
	"net/http"

	db "github.com/KibetBrian/backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetPendingVoters(c *gin.Context){
	ctx := context.Background();
	voters, err := s.db.PendingVoters(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(voters) == 0 {
		c.JSON(200,gin.H{"pendingVoters": []db.Voter{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pendingVoters": voters})
}

func (s *Server) GetTotalRegisteredVoters(c *gin.Context){
	ctx := context.Background();
	voters, err := s.db.VerifiedVoters(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(voters) == 0 {
		c.JSON(200, []db.Voter{})
		return
	}
	c.JSON(http.StatusOK, voters)
}

func (s *Server) GetTotalVotersNumb(c *gin.Context){
	ctx := context.Background();
	num, err := s.db.TotalVotersNum(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, num)
}

func (s *Server) GetTotalVotedVoters(c *gin.Context){
	ctx := context.Background();
	num, err := s.db.TotalVotedVoters(ctx);
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(200, num)
}
