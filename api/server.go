package api

import (
	db "github.com/codeninjaug/authexample/db/sqlc"
	"github.com/gin-gonic/gin"
)

//Server serves HTTP requests for our auth service
type Server struct{
	 store db.Store
	 router *gin.Engine
}

func NewServer(store db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()
	server.router = router
	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.GetUser)
	router.GET("/users", server.GetUsers)
	return server
}

//start runs the http server on a specific address
func (server *Server) Start(address string) error{
     return server.router.Run(address)
}

//errror response for errors
func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}