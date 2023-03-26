package api

import (
	"database/sql"
	"net/http"

	db "github.com/codeninjaug/authexample/db/sqlc"
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

	arg := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName: req.LastName,
	    Phone: req.Phone,
	    Email : req.Email,
	    Password: req.Password,
	} 

	user , err := server.store.CreateUser(ctx, arg)
	if err !=nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct{
	ID int32 `uri:"id" binding:"required, min=1"`
}


func (server *Server) GetUser(ctx *gin.Context){
    var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
       ctx.JSON(http.StatusBadRequest, errorResponse(err))
	   return 
	} 
	user, err := server.store.GetUser(ctx, req.ID)
	if err !=nil{
		if err ==  sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
	   ctx.JSON(http.StatusInternalServerError, errorResponse(err))	
	   return
	}
	ctx.JSON(http.StatusOK, user)
}


type listUserRequest struct{
	pageID   int32 `form:"page_id" binding:"required, min=1"`
	pageSize int32 `form:"page_size" binding:"required, min=5, max=10"`
}

func (server *Server) GetUsers(ctx *gin.Context){
    var req listUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
       ctx.JSON(http.StatusBadRequest, errorResponse(err))
	   return 
	} 
	arg := db.ListUsersParams{
		 Limit: req.pageSize,
		 Offset: (req.pageID-1) * req.pageSize,
	}
	user, err := server.store.ListUsers(ctx, arg)
	if err !=nil{
	   ctx.JSON(http.StatusInternalServerError, errorResponse(err))	
	}
	ctx.JSON(http.StatusOK, user)
}