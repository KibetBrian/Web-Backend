package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	db "github.com/KibetBrian/backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserRequest struct {
	Id uuid.UUID `json:"id" uri:"email" binding:"required"`
}

type LoginResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
}


func (s *Server) RegisterUser(ctx *gin.Context){
	var user User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse("Binding failed", err))
		return
	}
	
	row, err:= s.db.CheckEmail(context.Background(), user.Email)
	if err != nil && err != sql.ErrNoRows{
		ctx.JSON(http.StatusInternalServerError, ErrResponse("Error occured", err))
		return
	}
	if int(row.Count) > 0{
		ctx.JSON(http.StatusConflict, ErrResponse("Seems you are already registered",errors.New("Email already exist")))
		return
	}

	arg := db.RegisterUserParams{
		FullName: user.FullName,
		Email: user.Email,
		Password: user.Password,
	}
	
	registered, err := s.db.RegisterUser(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse("Failed to register user", err))
		return
	}
	ctx.JSON(http.StatusOK,registered)
}

func ErrResponse(message string, err error) gin.H{
	res := gin.H{
		"Message": message,
		"Error: ": err,
	}
	return res
}

func (s *Server) GetUser(ctx *gin.Context){
	var req GetUserRequest;

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse("Error occured", err))
		return
	}
	
	user, err := s.db.GetUser(context.Background(), req.Id)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(404, ErrResponse("No user with such id", err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,ErrResponse("Error occured while getting user", err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *Server)Login (c *gin.Context){
	var req LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrResponse("Invalid credentials format", err))
		return
	}
	ctx := context.Background()
	user, err := s.db.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrResponse("Error occured while getting email", err))
		return
	}
	if (user.Password != req.Password){
		c.JSON(http.StatusUnauthorized, ErrResponse("Invalid email or password", nil))
		return
	}
	res := &LoginResponse{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
	}
	c.JSON(http.StatusOK,res)
}