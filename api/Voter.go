package api

import (
	"context"
	"net/http"

	db "github.com/KibetBrian/backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
type PendingVotersResponse struct {
	Id uuid.UUID `json:"id"`
	FullName string `json:"fullName"`
	Email string `json:"email"`
	NationalIdNumber int64     `json:"nationalIdNumber"`
	Voted bool `json:"voted"`
	Verified bool `json:"verified"`
	EthereumAddress  string `json:"ethereumAddress"`
}

func (s *Server) GetPendingVoters(c *gin.Context){
	ctx := context.Background();

	res := []PendingVotersResponse{};

	voters, err := s.db.PendingVoters(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for _, v := range voters {
		voter := PendingVotersResponse{
			Id: v.ID,
			FullName: v.FullName,
			Email: v.Email,
			NationalIdNumber: v.NationalIDNumber,
			Voted: v.Voted.Bool,
			Verified: v.Verified.Bool,
			EthereumAddress:  v.EthereumAddress,
		}
		res = append(res, voter)
	}

	if len(voters) == 0 {
		c.JSON(200,gin.H{"pendingVoters": []db.Voter{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"pendingVoters": res})
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
