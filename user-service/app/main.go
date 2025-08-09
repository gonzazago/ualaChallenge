package main

import (
	"context"
	"log"
	"user-service/config"
)

func main() {
	// Creamos el contexto principal de la aplicación.
	log.Println("Iniciando User Service...")
	ctx := context.Background()

	// 1. Cargar configuración
	cfg := config.LoadConfig()

	// 2. Crear el contenedor de dependencias
	container := NewContainer(cfg)

	// 3. Obtener el servidor desde el contenedor (con sus dependencias inyectadas)
	server := container.GetServer(ctx)

	// 4. Iniciar el servidor web
	log.Printf("Servidor escuchando en el puerto %s", cfg.ServerPort)
	if err := server.Start(ctx); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
