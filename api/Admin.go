package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	db "github.com/KibetBrian/backend/db/sqlc"
	"github.com/KibetBrian/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VoterRegistrationRequest struct {
	UserId uuid.UUID `json:"userId"`
	VotersName string `json:"votersName"`
	VotersEmail string `json:"votersEmail"`
	VotersAddress string `json:"votersAddress"`
}

type ContestantRegistrationRequest struct {
	FullName string `json:"fullName" binding:"required"`
	Email string `json:"email" binding:"required"`
	Position string `json:"position" binding:"required"`
	Region string `json:"region" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageAddress string `json:"imageAddress" binding:"required"`
	EthereumAddress string `json:"ethereumAddress" binding:"required"`
	NationalIdNumber int64 `json:"nationalIdNumber" binding:"required"`	
}

type VoterChainParams struct {
	VotersAddress string `json:"votersAddress"`
}

func (s *Server) RegisterVoter(c  *gin.Context) {
	var req VoterRegistrationRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResponse("Bad request, try api with correct credentials", err))
		return
	}

	ctx := context.Background()
	dbUser, err := s.db.GetUser(ctx, req.UserId)
	if err != nil{
		c.JSON(http.StatusNotFound, "Seems you have not signed up")
		return
	}
	if dbUser.Email != req.VotersEmail{
		c.JSON(http.StatusForbidden, "UserId doesn't match with email provided")
		return
	}	
	arg := db.RegisterVoterParams{
		 FullName: req.VotersName,
		 Email: req.VotersEmail,
		 VotersPublicAddress: req.VotersAddress,
	}


	//Make an post  api request to blockchain client
	const url = "http://127.0.0.1:8000/voter/add"
	const method = "POST"
	const contentType = "application/json"

	values := map[string]string{
		"admin": "brian",
		"address": arg.VotersPublicAddress,
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		c.JSON(http.StatusInternalServerError,"Failed to marshal values")
		return
	}
	res, err := http.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError,"Failed to make a post request to blockchain client")
		return
	}
	var jsonRes map[string]interface{}
	json.NewDecoder(res.Body).Decode(&jsonRes)

	if jsonRes["success"] == false {
		c.JSON(http.StatusConflict, jsonRes)
		return
	}

	voterId, err := s.db.RegisterVoter(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError,utils.ErrResponse("Voter registration failed", err))
		return
	}

	c.JSON(http.StatusOK,gin.H{"Message": "Voter Registered", "VoterId: ": voterId, "Blockchain ":jsonRes})
}


// This function makes an api request to the blockchain client and returs list of registered voters
func(s *Server) GetChainVoters(c *gin.Context){
	const  url = "http://127.0.0.1:8000/voter/all"
	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrResponse("Failed to make http get request", err))
		return
	}
	defer res.Body.Close()
	data , err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrResponse("Failed to read the response body", err))
		return
	}

	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil{
		c.JSON(http.StatusInternalServerError, utils.ErrResponse("Failed to unmarshal the response body", err))
		return
	}
	c.JSON(http.StatusOK, jsonData)
}

func (s *Server) RegisterContestant(c *gin.Context){
	var req ContestantRegistrationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrResponse("Invalid parameters. Error: ", err))
		return
	}

	//Make a post request to the blockchain client
	const url = "http://127.0.0.1:8000/candidate/add"
	const method = "POST"
	const contentType = "application/json"

	values := map[string]string{
		"admin": "brian",
		"candidatesName": req.FullName,
		"candidatesAddress": req.EthereumAddress,
		"position": req.Position,
	}
	jsonData, err := json.Marshal(values)
	if err != nil {
		c.JSON(http.StatusInternalServerError,"Failed to marshal values")
		return
	}
	res, err := http.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError,"Failed to make a post request to blockchain client")
		return
	}
	var jsonRes map[string]interface{}
	json.NewDecoder(res.Body).Decode(&jsonRes)

	if jsonRes["success"] == false {
		c.JSON(http.StatusConflict, jsonRes)
		return
	}
	//Save data to the sql database
	arg := db.RegisterContestantParams{
		FullName: req.FullName,
		Email: req.Email,
		Position: req.Position,
		Description: req.Description,
		Region: req.Region,
		EthereumAddress: req.EthereumAddress,
		NationalIDNumber: req.NationalIdNumber,
		ImageAddress: req.ImageAddress,
	}
	ctx := context.Background();
	contestant, err := s.db.RegisterContestant(ctx, arg)
	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message": "Failed to register contestant", "error": err})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message:":"Contestant registered","contestant":contestant})
}