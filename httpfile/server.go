package httpfile

import (
	"demo/dbfile"
	"demo/httpfile/handlers"

	"github.com/gin-gonic/gin"
)

// it is a Server
type Server struct {
	// db client
	dc *dbfile.Client
	// router
	router *gin.Engine
}

// NewServer returns a server instance.
func NewServer(dc *dbfile.Client) *Server {
	return &Server{
		dc:     dc,
		router: gin.Default(),
	}
}

func (s *Server) registerRoutes() {
	s.router.GET("/user/:id", handlers.ReadUser(s.dc)) // id must be a int
	s.router.POST("/user", handlers.CreateUser(s.dc))  //
	s.router.PUT("/user", handlers.UpdateUser(s.dc))
	s.router.DELETE("/user/:id", handlers.DeleteUser(s.dc))
	// s.router.DELETE("/user/:name", handlers.DeleteUserByName(s.dc))
}

// Run starts server.
func (s *Server) Run(port string) {
	// register routes
	s.registerRoutes()
	s.router.Run(port)
}
