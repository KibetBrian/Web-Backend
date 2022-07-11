package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

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