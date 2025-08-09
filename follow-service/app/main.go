package main

import (
	"context"
	"follow-service/config"
	"log"
)

func main() {
	log.Println("Iniciando Follow Service (modo in-memory)...")
	ctx := context.Background()

	cfg := config.LoadConfig()
	container := NewContainer(cfg)
	server := container.GetServer(ctx)
	log.Printf("Servidor escuchando en el puerto %s", cfg.ServerPort)
	if err := server.Start(ctx); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
