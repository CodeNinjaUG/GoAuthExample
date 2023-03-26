package api

import (
	db "github.com/codeninjaug/authexample/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for our auth service
type Server struct{
	 store *db.Store
	 router *gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()
	server.router = router
	router.POST("/create/user", server.CreateUser())
	return server
}


func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}