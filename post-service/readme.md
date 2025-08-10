### Post Service
Este microservicio es el responsable de gestionar la creación y consulta de posts (o "tweets"). Cumple con los requisitos de "publicar mensajes cortos" y proporciona los datos necesarios para que el timeline-service pueda construir el feed.

### Arquitectura Interna
El servicio sigue un patrón de Arquitectura Hexagonal para mantener una clara separación de responsabilidades:

* Domain: Contiene la lógica de negocio pura, las entidades (Post), las interfaces (Repository, Service, Notifier) y las validaciones (ej. límite de 280 caracteres).

* Delivery: La capa de API REST, encargada de manejar las peticiones y respuestas HTTP.

* Infra: Implementaciones concretas, como el repositorio en memoria y un notificador de eventos simulado (DummyNotifier).

### API Endpoints
Todas las rutas están prefijadas con /api/v1.

#### Método

* Ruta
Descripción: Crea un nuevo post.

##### POST

```
/posts
```
##### GET
Descripcion: Obtiene los posts de una lista de usuarios, filtrando por el query param user_ids.
```
/posts
```
Ejemplo de Payload para POST /posts
```json
{
    "user_id": "d9a7f3b0-5e1a-4f3c-8a2d-1b9c8e7f6a5b",
    "text": "¡Este es mi primer tweet!"
}
```

* Ejemplo de Petición para GET /posts
http://localhost:8080/api/v1/posts?user_ids=d9a7f3b0-5e1a-4f3c-8a2d-1b9c8e7f6a5b,otro-user-id

Cómo Correr el Servicio (Standalone)
Prerrequisitos
Go (versión 1.24 o superior).

##### Ejecución
Desde la raíz del repositorio (ualaChallenge/), ejecuta los siguientes comandos:

# 1. Navega al directorio del servicio
```bash
$ cd post-service
```

# 2. Instala las dependencias
```bash
$ go mod tidy
```

# 3. Corre la aplicación
```bash
$ make web
```
El servicio estará escuchando en http://localhost:8080.

Testing
Para ejecutar la suite de tests unitarios y de integración, navega al directorio del servicio y ejecuta:

```bash
$ make test ./...
```
