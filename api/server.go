package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/petrostrak/simple-bank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(s *db.Store) *Server {
	server := &Server{store: s}
	router := gin.Default()

	router.POST("/accounts", server.CreateAccount)

	server.router = router
	return server
}
