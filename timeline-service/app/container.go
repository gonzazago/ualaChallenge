package main

import (
	"context"
	"net/http"
	"timeline-service/app/web"
	"timeline-service/config"
	timelineHandler "timeline-service/internal/delivery/rest/timeline"
	"timeline-service/internal/domain/timeline"
	"timeline-service/internal/infra/client/follow"
	"timeline-service/internal/infra/client/post"
)

type Container struct {
	config *config.Config

	// Dependencias cacheadas
	server             web.Server
	getTimelineHandler web.Handler
	timelineService    timeline.Service
	followClient       timeline.FollowClient
	postClient         timeline.PostClient
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
			c.GetTimelineHandler(ctx),
		)
	}
	return c.server
}

// GetTimelineHandler devuelve el handler para obtener el timeline.
func (c *Container) GetTimelineHandler(ctx context.Context) web.Handler {
	if c.getTimelineHandler == nil {
		service := c.GetTimelineService(ctx)
		c.getTimelineHandler = timelineHandler.NewGetTimelineHandler(service)
	}
	return c.getTimelineHandler
}

// GetTimelineService devuelve la implementación del servicio de dominio.
func (c *Container) GetTimelineService(ctx context.Context) timeline.Service {
	if c.timelineService == nil {
		followClient := c.GetFollowClient(ctx)
		postClient := c.GetPostClient(ctx)
		c.timelineService = timeline.NewService(followClient, postClient)
	}
	return c.timelineService
}

// GetFollowClient devuelve la implementación del cliente para el follow-service.
func (c *Container) GetFollowClient(ctx context.Context) timeline.FollowClient {
	if c.followClient == nil {
		c.followClient = follow.NewFollowClient(http.DefaultClient, c.config.FollowServiceURL)
	}
	return c.followClient
}

// GetPostClient devuelve la implementación del cliente para el post-service.
func (c *Container) GetPostClient(ctx context.Context) timeline.PostClient {
	if c.postClient == nil {
		c.postClient = post.NewPostClient(http.DefaultClient, c.config.PostServiceURL)
	}
	return c.postClient
}
