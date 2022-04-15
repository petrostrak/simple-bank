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

// NewServer creates a new HTTP server and setup routing.
func NewServer(s *db.Store) *Server {
	server := &Server{store: s}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
