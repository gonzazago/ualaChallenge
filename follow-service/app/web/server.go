package web

import (
	"context"
	"github.com/gin-gonic/gin"
)

// Server interface that define web server.
type Server interface {
	Start(ctx context.Context) error
}

type ginServer struct {
	engine              *gin.Engine
	port                string
	followUserHandler   Handler
	getFollowingHandler Handler
}

// NewServer create and config a new server
func NewServer(port string, followUserHandler Handler, getFollowingHandler Handler) Server {
	s := &ginServer{
		engine:              gin.Default(),
		port:                port,
		followUserHandler:   followUserHandler,
		getFollowingHandler: getFollowingHandler,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configure routes for our app.
func (s *ginServer) setupRoutes() {
	v1 := s.engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/:followerID/follow", AdaptHandler(s.followUserHandler))
			users.GET("/:userID/following", AdaptHandler(s.getFollowingHandler))
		}
	}
}

// Start server and listen port
func (s *ginServer) Start(ctx context.Context) error {
	return s.engine.Run(s.port)
}
