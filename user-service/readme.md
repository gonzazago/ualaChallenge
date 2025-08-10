## User Service
Este microservicio es el responsable de gestionar toda la información relacionada con los perfiles de usuario. Se encarga de la creación, consulta y validación de los datos de los usuarios.

## Arquitectura Interna
El servicio sigue un patrón de Arquitectura Hexagonal (Puertos y Adaptadores) para mantener una clara separación de responsabilidades:

* Domain: Contiene la lógica de negocio pura, las entidades (User) y las interfaces (UserRepository, Service).

* Delivery: La capa de API REST, encargada de manejar las peticiones y respuestas HTTP.

* Infra: Implementaciones concretas, como el repositorio en memoria.

### API Endpoints
Todas las rutas están prefijadas con /api/v1.

#### Método

* Ruta

#### Descripción

##### POST
```
/users
```

Crea un nuevo usuario.

GET

```
/users/{userID}
```

Obtiene la información de un usuario por su ID.

GET

```
/health
```

Endpoint de health check del servicio.

Ejemplo de Payload para POST /users
```json
{
"username": "gonzazago",
"email": "gonzalo.zago@example.com"
}
```

Cómo Correr el Servicio (Standalone)
Prerrequisitos
Go (versión 1.24 o superior).

Ejecución
Desde la raíz del repositorio (ualaChallenge/), ejecuta los siguientes comandos:

# 1. Navega al directorio del servicio
$ cd user-service

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
