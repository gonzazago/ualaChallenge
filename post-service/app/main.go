package main

import (
	"context"
	"log"
	"post-service/config"
)

func main() {
	log.Println("Iniciando Post Service (modo in-memory)...")
	ctx := context.Background()

	// 1. Cargar configuración
	cfg := config.LoadConfig()

	// 2. Crear el contenedor de dependencias
	container := NewContainer(cfg)

	// 3. Obtener el servidor desde el contenedor
	server := container.GetServer(ctx)

	// 4. Iniciar el servidor web
	log.Printf("Servidor escuchando en el puerto %s", cfg.ServerPort)
	if err := server.Start(ctx); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
