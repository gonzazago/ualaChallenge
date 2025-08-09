package main

import (
	"context"
	"follow-service/app/web"
	"follow-service/config"
	followHandler "follow-service/internal/delivery/rest/follow"
	"follow-service/internal/domain/follow"
	"follow-service/internal/infra/db/follow/inmemory"
)

// Container gestiona el ciclo de vida de las dependencias.
type Container struct {
	config *config.Config

	// Dependencias cacheadas
	server              web.Server
	followUserHandler   web.Handler
	getFollowingHandler web.Handler
	followService       follow.Service
	followRepository    follow.Repository
}

// NewContainer crea una nueva instancia del contenedor.
func NewContainer(cfg *config.Config) *Container {
	return &Container{
		config: cfg,
	}
}

// GetServer devuelve una instancia del servidor web.
func (c *Container) GetServer(ctx context.Context) web.Server {
	if c.server == nil {
		c.server = web.NewServer(
			c.config.ServerPort,
			c.GetFollowUserHandler(ctx),
			c.GetFollowingHandler(ctx),
		)
	}
	return c.server
}

// GetFollowUserHandler devuelve el handler para seguir a un usuario.
func (c *Container) GetFollowUserHandler(ctx context.Context) web.Handler {
	if c.followUserHandler == nil {
		service := c.GetFollowService(ctx)
		c.followUserHandler = followHandler.NewFollowUserHandler(service)
	}
	return c.followUserHandler
}

// GetFollowingHandler devuelve el handler para obtener la lista de seguidos.
func (c *Container) GetFollowingHandler(ctx context.Context) web.Handler {
	if c.getFollowingHandler == nil {
		service := c.GetFollowService(ctx)
		c.getFollowingHandler = followHandler.NewGetFollowingHandler(service)
	}
	return c.getFollowingHandler
}

// GetFollowService devuelve la implementación del servicio de dominio.
func (c *Container) GetFollowService(ctx context.Context) follow.Service {
	if c.followService == nil {
		repo := c.GetFollowRepository(ctx)
		c.followService = follow.NewService(repo)
	}
	return c.followService
}

// GetFollowRepository devuelve la implementación del repositorio en memoria.
func (c *Container) GetFollowRepository(ctx context.Context) follow.Repository {
	if c.followRepository == nil {
		c.followRepository = inmemory.NewInMemoryFollowRepository()
	}
	return c.followRepository
}
