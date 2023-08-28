package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type Server struct {
	store  *db.Store   // store the db connection
	router *gin.Engine // gin engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// API routers
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	// (server *Server) 是一个方法接收器，即 Server拥有方法 Start
	return server.router.Run(address)
}
