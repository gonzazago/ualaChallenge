package web

import (
	"context"
	"github.com/gin-gonic/gin"
)

// Server es una interfaz que define nuestro servidor web.
type Server interface {
	Start(ctx context.Context) error
}

type ginServer struct {
	engine            *gin.Engine
	port              string
	createPostHandler Handler
	getPostsHandler   Handler
}

// NewServer crea y configura una nueva instancia del servidor.
func NewServer(port string, createPostHandler Handler, getPostsHandler Handler) Server {
	s := &ginServer{
		engine:            gin.Default(),
		port:              port,
		createPostHandler: createPostHandler,
		getPostsHandler:   getPostsHandler,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configura las rutas de la aplicación.
func (s *ginServer) setupRoutes() {
	v1 := s.engine.Group("/api/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.POST("/", AdaptHandler(s.createPostHandler))
			posts.GET("/", AdaptHandler(s.getPostsHandler))
		}
	}
}

// Start inicia el servidor HTTP y lo pone a escuchar en el puerto configurado.
func (s *ginServer) Start(ctx context.Context) error {
	return s.engine.Run(s.port)
}
