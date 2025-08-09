package main

import (
	"context"
	"user-service/app/web"
	"user-service/config"
	"user-service/internal/delivery/rest"
	userHandlers "user-service/internal/delivery/rest/users"
	"user-service/internal/domain/users"
	"user-service/internal/infra/db/users/inmemory"
)

type Container struct {
	config *config.Config

	// Dependencias cacheadas
	server             web.Server
	createUserHandler  web.Handler
	getUserByIDHandler web.Handler
	healthHandler      web.Handler
	userService        users.Service
	userRepository     users.Repository
}

// NewContainer crea una nueva instancia del contenedor.
func NewContainer(cfg *config.Config) *Container {
	return &Container{
		config: cfg,
	}
}

// GetServer devuelve una instancia del servidor web, creándola si es necesario.
func (c *Container) GetServer(ctx context.Context) web.Server {
	if c.server == nil {
		c.server = web.NewServer(
			c.config.ServerPort,
			c.GetCreateUserHandler(ctx),
			c.GetHealthHandler(),
			c.GetUserByIDHandler(ctx),
		)
	}
	return c.server
}

// GetCreateUserHandler devuelve el handler para crear usuarios.
func (c *Container) GetCreateUserHandler(ctx context.Context) web.Handler {
	if c.createUserHandler == nil {
		service := c.GetUserService(ctx)
		c.createUserHandler = userHandlers.NewCreateUserHandler(service)
	}
	return c.createUserHandler
}

func (c *Container) GetUserByIDHandler(ctx context.Context) web.Handler {
	if c.getUserByIDHandler == nil {
		service := c.GetUserService(ctx)
		c.getUserByIDHandler = userHandlers.NewGetUserHandler(service)
	}
	return c.getUserByIDHandler
}

func (c *Container) GetHealthHandler() web.Handler {
	if c.healthHandler == nil {
		c.healthHandler = rest.NewHealthHandler()
	}
	return c.healthHandler
}

// GetUserService devuelve la implementación del servicio de dominio de usuario.
func (c *Container) GetUserService(ctx context.Context) users.Service {
	if c.userService == nil {
		repo := c.GetUserRepository(ctx)
		c.userService = users.NewService(repo)
	}
	return c.userService
}

// GetUserRepository devuelve la implementación del repositorio de usuario para MySQL.
func (c *Container) GetUserRepository(ctx context.Context) users.Repository {
	if c.userRepository == nil {
		c.userRepository = inmemory.NewInMemoryUserRepository()
	}
	return c.userRepository
}
