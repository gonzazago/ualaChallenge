## Timeline Service
Este microservicio es el responsable de orquestar la comunicación con otros servicios para construir y devolver el timeline personalizado de un usuario. Cumple con el requisito principal del desafío: "ver una línea de tiempo que muestre los tweets de los usuarios a los que siguen".

## Arquitectura Interna
El servicio sigue un patrón de Arquitectura Hexagonal para mantener una clara separación de responsabilidades:

* Domain: Contiene la lógica de orquestación y las interfaces (Service, FollowServiceClient, PostServiceClient) que definen los contratos con los servicios externos.

* Delivery: La capa de API REST, encargada de manejar las peticiones y respuestas HTTP.

* Infra: Implementaciones concretas de los clientes HTTP que se comunican con el follow-service y el post-service.

## API Endpoints
Todas las rutas están prefijadas con /api/v1.

### Método

* Ruta

Descripción

GET
```
/users/{userID}/timeline
```

Obtiene el timeline de un usuario específico.

Cómo Correr el Servicio (Standalone)
Prerrequisitos
Go (versión 1.21 o superior).

Los servicios user-service, follow-service y post-service deben estar corriendo.

Ejecución
Desde la raíz del repositorio (ualaChallenge/), ejecuta los siguientes comandos:

# 1. Navega al directorio del servicio
```bash
cd timeline-service
```

# 2. Instala las dependencias
```bash
go mod tidy
```
# 3. Corre la aplicación
```bash
make web
```
El servicio estará escuchando en http://localhost:8080.

Testing
Para ejecutar la suite de tests unitarios y de integración, navega al directorio del servicio y ejecuta:
```bash
make test
```
