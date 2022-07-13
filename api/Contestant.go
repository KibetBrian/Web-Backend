package api

import (
	"context"

	db "github.com/KibetBrian/backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AllContestantRequest struct {
	ID uuid.UUID    `json:"id"`
	FullName string `json:"fullName"`
	Position string `json:"position"`
	Address string `json:"address"`
}



func (s *Server) GetAllContestants(c *gin.Context){
	ctx := context.Background();
	contestants, err := s.db.GetAllCandidates(ctx)
	if err != nil{
		c.JSON(500, gin.H{"error": err})
		return
	}
	allContestants := []AllContestantRequest{}
	for _,v := range contestants{
		arg := &AllContestantRequest{
			ID: v.ID,
			FullName:v.FullName,
			Position:v.Position,
			Address:v.EthereumAddress,
		}
		allContestants = append(allContestants, *arg);
	}
	c.JSON(200, allContestants)
}

func (s *Server) GetPresidentContestants(c *gin.Context){
	ctx := context.Background();
	contestants, err := s.db.GetPresidentialCandidates(ctx)
	if err != nil{
		c.JSON(500, gin.H{"error": err})
		return
	}
	if len(contestants) == 0{
		c.JSON(200, []db.Contestant{})
		return
	}
	c.JSON(200, contestants)
}
func (s *Server) GetGorvenorContestants(c *gin.Context){
	ctx := context.Background();
	contestants, err := s.db.GetGubernatorialCandidates(ctx)
	if err != nil{
		c.JSON(500, gin.H{"error": err})
		return
	}
	if len(contestants) == 0{
		c.JSON(200, []db.Contestant{})
		return
	}
	c.JSON(200, contestants)
}