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
	engine             *gin.Engine
	port               string
	getTimelineHandler Handler
}

// NewServer crea y configura una nueva instancia del servidor.
func NewServer(port string, getTimelineHandler Handler) Server {
	s := &ginServer{
		engine:             gin.Default(),
		port:               port,
		getTimelineHandler: getTimelineHandler,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configura las rutas de la aplicación.
func (s *ginServer) setupRoutes() {
	v1 := s.engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/:userID/timeline", AdaptHandler(s.getTimelineHandler))
		}
	}
}

// Start inicia el servidor HTTP y lo pone a escuchar en el puerto configurado.
func (s *ginServer) Start(ctx context.Context) error {
	return s.engine.Run(s.port)
}
