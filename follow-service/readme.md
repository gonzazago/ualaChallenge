# Follow Service
Este microservicio es el responsable de gestionar el grafo social de la plataforma. Se encarga de crear y consultar las relaciones de seguimiento entre los usuarios, cumpliendo con el requisito de "poder seguir a otros usuarios".

## Arquitectura Interna
El servicio sigue un patrón de Arquitectura Hexagonal para mantener una clara separación de responsabilidades:

* Domain: Contiene la lógica de negocio pura, las interfaces (Repository, Service) y las validaciones (ej. un usuario no puede seguirse a sí mismo).

* Delivery: La capa de API REST, encargada de manejar las peticiones y respuestas HTTP.

* Infra: Implementaciones concretas, como el repositorio en memoria.

### API Endpoints
Todas las rutas están prefijadas con /api/v1.

* Método

* Ruta

* Descripción: Permite que un usuario siga a otro.

POST

```
/users/{followerID}/follow
```

* GET
Descripción: Obtiene la lista de usuarios que sigue un usuario.
```
/users/{userID}/following
```


Ejemplo de Payload para POST /users/{followerID}/follow
```json
{
    "user_id_to_follow": "uuid-del-usuario-a-seguir"
}
```

### Cómo Correr el Servicio (Standalone)
##### Prerrequisitos
* Go (versión 1.24 o superior).

#### Ejecución
Desde la raíz del repositorio (ualaChallenge/), ejecuta los siguientes comandos:

##### 1. Navega al directorio del servicio
```bash
$ cd follow-service
```

##### 2. Instala las dependencias
```bash
$ go mod tidy
```

##### 3. Corre la aplicación
```bash
$ make web
```
El servicio estará escuchando en http://localhost:8080.

##### Testing
Para ejecutar la suite de tests unitarios y de integración, navega al directorio del servicio y ejecuta:

```bash
$ make test
```
Para regenerar los mocks si cambias alguna interfaz, usa el Makefile del servicio:

$ make mocks
