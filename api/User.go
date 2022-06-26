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

type GetUserRequest struct {
	Id uuid.UUID `json:"id" uri:"email" binding:"required"`
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