package web

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(ctx context.Context) error
}

// ginServer envuelve el motor de Gin y gestiona la lógica del servidor HTTP.
type ginServer struct {
	engine             *gin.Engine
	port               string
	createUserHandler  Handler
	getUserByIDHandler Handler
	healthHandler      Handler
}

// NewServer crea y configura una nueva instancia del servidor.
func NewServer(port string,
	createUserHandler,
	healthHandler,
	getUserByIDHandler Handler) Server {
	s := &ginServer{
		engine:             gin.Default(),
		port:               port,
		healthHandler:      healthHandler,
		createUserHandler:  createUserHandler,
		getUserByIDHandler: getUserByIDHandler,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configura las rutas de la aplicación.
func (s *ginServer) setupRoutes() {
	s.engine.GET("/health", AdaptHandler(s.healthHandler))
	v1 := s.engine.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", AdaptHandler(s.createUserHandler))
			users.GET("/:userID", AdaptHandler(s.getUserByIDHandler))
		}
	}
}

// Start inicia el servidor HTTP y lo pone a escuchar en el puerto configurado.
func (s *ginServer) Start(ctx context.Context) error {
	return s.engine.Run(s.port)
}
