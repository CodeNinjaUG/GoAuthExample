package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct{
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}


func(server *Server) createUser(ctx *gin.Context){
    var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
       ctx.JSON(http.StatusBadRequest, errorResponse(err))
	   return 
	} 
}