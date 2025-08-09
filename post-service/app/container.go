package main

import (
	"context"
	"post-service/app/web"
	"post-service/config"
	postHandlers "post-service/internal/delivery/rest/post"
	"post-service/internal/domain/post"
	"post-service/internal/infra/db/post/inmemory"
	"post-service/internal/infra/queue"
	"post-service/internal/infra/queue/dummy"
)

type Container struct {
	config *config.Config

	// Dependencias cacheadas
	server            web.Server
	createPostHandler web.Handler
	getPostsHandler   web.Handler
	postService       post.Service
	postRepository    post.Repository
	notifier          queue.Notifier
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
			c.GetCreatePostHandler(ctx),
			c.GetGetPostsHandler(ctx),
		)
	}
	return c.server
}

// GetCreatePostHandler devuelve el handler para crear un post.
func (c *Container) GetCreatePostHandler(ctx context.Context) web.Handler {
	if c.createPostHandler == nil {
		service := c.GetPostService(ctx)
		c.createPostHandler = postHandlers.NewCreatePostHandler(service)
	}
	return c.createPostHandler
}

// GetGetPostsHandler devuelve el handler para obtener posts.
func (c *Container) GetGetPostsHandler(ctx context.Context) web.Handler {
	if c.getPostsHandler == nil {
		service := c.GetPostService(ctx)
		c.getPostsHandler = postHandlers.NewGetPostsHandler(service)
	}
	return c.getPostsHandler
}

// GetPostService devuelve la implementación del servicio de dominio.
func (c *Container) GetPostService(ctx context.Context) post.Service {
	if c.postService == nil {
		repo := c.GetPostRepository(ctx)
		notifier := c.GetPostNotifier(ctx)
		c.postService = post.NewService(repo, notifier)
	}
	return c.postService
}

// GetPostRepository devuelve la implementación del repositorio en memoria.
func (c *Container) GetPostRepository(ctx context.Context) post.Repository {
	if c.postRepository == nil {
		c.postRepository = inmemory.NewInMemoryPostRepository()
	}
	return c.postRepository
}

// GetPostNotifier devuelve la implementación del notificador de eventos.
func (c *Container) GetPostNotifier(ctx context.Context) queue.Notifier {
	if c.notifier == nil {
		c.notifier = dummy.NewDummyNotifier()
	}
	return c.notifier
}
