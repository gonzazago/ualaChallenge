Desafío Técnico Ualá - Plataforma de Microblogging
Este repositorio contiene la implementación de una versión simplificada de una plataforma de microblogging, desarrollada
como parte del desafío técnico de Ualá. La solución está diseñada siguiendo un enfoque de microservicios para garantizar
la escalabilidad y el mantenimiento.

## 1. Arquitectura de la Solución

La plataforma se basa en una arquitectura de microservicios desacoplados, donde cada servicio tiene una única
responsabilidad bien definida. Se ha seguido un patrón de Arquitectura Hexagonal (Puertos y Adaptadores) inspirado en
los principios de Domain-Driven Design (DDD) para lograr una clara separación entre la lógica de negocio, la capa de
aplicación y la infraestructura.

El lenguaje elegido para todos los servicios es Go, por su alto rendimiento, su excelente manejo de la concurrencia y su
robusto ecosistema para el desarrollo de APIs.

## Diagrama de Arquitectura

<img width="5276" height="3224" alt="image" src="https://github.com/user-attachments/assets/6ac1f8c9-4794-424d-8e8b-703b316929bf" />


### Descripción de los Servicios Implementados

* API Gateway: Actúa como el único punto de entrada para todas las peticiones externas. Se encarga de enrutar el tráfico
  al microservicio correspondiente.

* User Service: Gestiona toda la información relacionada con los perfiles de usuario, como la creación y consulta de
  datos.
  * Utiliza como persistencia una MySql, debido a que son datos que no debieran variar con el tiempo y mantienen una estructura solida
   a su vez una cache para optimizar la lectura de datos de usuario

* Follow Service: Administra el grafo social, es decir, las relaciones de seguimiento entre los usuarios.
  * Utiliza Aws Neptune, la eleccion de dicha base se debe a que la misma es optimizado para grafos, lo que facilitaria la busqueda de los followers de los usuarios 

* Post Service (Ingest): Se encarga de la ingesta de nuevos posts ("tweets"), validando el contenido y notificando a
  otros sistemas de forma asíncrona (simulado).
  * Post-service utiliza casandra como motor de persistencia,  habiendoe valuado dynamo y reddis, opte por cassandra, dado que la performance sobre read throughput asi como tambien es optima de cara a la escritura, como asi tambien a nivel de costes ya que con reddis o dynamo a una gran cantidad de tweets el coste de las misma se dispararia

* Timeline Service: Orquesta la comunicación entre los otros servicios para construir y devolver el timeline
  personalizado de un usuario.
  * Utiliza un reddis, para obtener los feed de manera eficiente,  en este caso lo que se busca en este reddis es obtener de manera rapidad un timeline por usuario, el mismo seria actualizado como consecuencia de los post de los seguidores
# 2. Setup y Ejecución del Proyecto

La solución está completamente "dockerizada", lo que permite levantar todo el stack de microservicios con un único
comando.

### Prerrequisitos

Tener instalado Docker y Docker Compose.

Tener instalado Make (generalmente viene preinstalado en macOS y Linux).

### Pasos para la Ejecución

Clona el repositorio y asegúrate de estar en el directorio raíz (donde se encuentra el archivo docker-compose.yml).

Abre una terminal en el directorio raíz.

#### Ejecuta el siguiente comando:

```
make run-all
```

Este comando construirá las imágenes de Docker para cada uno de los cinco servicios, los iniciará en segundo plano y los
conectará a una red virtual compartida.

Una vez que los contenedores estén corriendo, la plataforma estará completamente funcional y accesible a través del API
Gateway en el puerto 8000.

#### Detener la Aplicación

Para detener y eliminar todos los contenedores, ejecuta:

```
 make stop-all
```

# 3. Testing

El proyecto incluye una suite de tests unitarios

Ejecutar los Tests
Para correr los tests de un servicio específico, navega a su directorio y ejecuta:

```
go test ./...
```

# 4. Posmtman collection

```
{
    "info": {
        "_postman_id": "da86b199-00a4-47b3-9664-8ac780fbcc78",
        "name": "uala",
        "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
        "_exporter_id": "3310404"
    },
    "item": [
        {
            "name": "user-service",
            "item": [
                {
                    "name": "health",
                    "request": {
                        "method": "GET",
                        "header": [],
                        "url": {
                            "raw": "http://localhost:8000/health",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "health"
                            ]
                        }
                    },
                    "response": []
                },
                {
                    "name": "create-user",
                    "request": {
                        "method": "POST",
                        "header": [],
                        "body": {
                            "mode": "raw",
                            "raw": "{\n    \"username\":\"user-2\",\n    \"email\": \"user@email.com\"\n}",
                            "options": {
                                "raw": {
                                    "language": "json"
                                }
                            }
                        },
                        "url": {
                            "raw": "http://localhost:8000/api/v1/users",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "api",
                                "v1",
                                "users"
                            ]
                        }
                    },
                    "response": []
                },
                {
                    "name": "get-user-by-id",
                    "request": {
                        "method": "GET",
                        "header": [],
                        "url": {
                            "raw": "http://localhost:8080/api/v1/users/::userID",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8080",
                            "path": [
                                "api",
                                "v1",
                                "users",
                                "::userID"
                            ],
                            "variable": [
                                {
                                    "key": ":userID",
                                    "value": "a54cc0a5-2e06-4610-a6e5-d554a61f0dd7"
                                }
                            ]
                        }
                    },
                    "response": []
                }
            ]
        },
        {
            "name": "post-service",
            "item": [
                {
                    "name": "create-post",
                    "request": {
                        "method": "POST",
                        "header": [],
                        "body": {
                            "mode": "raw",
                            "raw": "{\n    \"user_id\": \"6afefafd-5aeb-4905-8f4c-d88bd03bc131\",\n    \"text\": \"¡Este es mi tercero tweet desde el post-service! #UalaChallenge\"\n}",
                            "options": {
                                "raw": {
                                    "language": "json"
                                }
                            }
                        },
                        "url": {
                            "raw": "http://localhost:8000/api/v1/posts/",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "api",
                                "v1",
                                "posts",
                                ""
                            ]
                        }
                    },
                    "response": []
                },
                {
                    "name": "get-post-by-user",
                    "request": {
                        "method": "GET",
                        "header": [],
                        "url": {
                            "raw": "http://localhost:8000/api/v1/posts?user_ids=a54cc0a5-2e06-4610-a6e5-d554a61f0dd7",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "api",
                                "v1",
                                "posts"
                            ],
                            "query": [
                                {
                                    "key": "user_ids",
                                    "value": "a54cc0a5-2e06-4610-a6e5-d554a61f0dd7"
                                }
                            ]
                        }
                    },
                    "response": []
                }
            ]
        },
        {
            "name": "follow-service",
            "item": [
                {
                    "name": "follow-user",
                    "request": {
                        "method": "POST",
                        "header": [],
                        "body": {
                            "mode": "raw",
                            "raw": "{\n    \"user_id_to_follow\": \"879e4a88-c4f6-437e-9384-e8fbaefcdb3e\"\n}",
                            "options": {
                                "raw": {
                                    "language": "json"
                                }
                            }
                        },
                        "url": {
                            "raw": "http://localhost:8000/api/v1/users/::followerID/follow",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "api",
                                "v1",
                                "users",
                                "::followerID",
                                "follow"
                            ],
                            "variable": [
                                {
                                    "key": ":followerID",
                                    "value": "879e4a88-c4f6-437e-9384-e8fbaefcdb3e"
                                }
                            ]
                        }
                    },
                    "response": []
                },
                {
                    "name": "user-following",
                    "protocolProfileBehavior": {
                        "disableBodyPruning": true
                    },
                    "request": {
                        "method": "GET",
                        "header": [],
                        "body": {
                            "mode": "raw",
                            "raw": "{\n    \"user_id_to_follow\": \"879e4a88-c4f6-437e-9384-e8fbaefcdb3e\"\n}",
                            "options": {
                                "raw": {
                                    "language": "json"
                                }
                            }
                        },
                        "url": {
                            "raw": "http://localhost:8000/api/v1/users/::userID/following",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "api",
                                "v1",
                                "users",
                                "::userID",
                                "following"
                            ],
                            "variable": [
                                {
                                    "key": ":userID",
                                    "value": "20f029db-7c59-48bb-8afb-10a64274a591"
                                }
                            ]
                        }
                    },
                    "response": []
                }
            ]
        },
        {
            "name": "timeline-service",
            "item": [
                {
                    "name": "get-timeline",
                    "request": {
                        "method": "GET",
                        "header": [],
                        "url": {
                            "raw": "http://localhost:8000/api/v1/users/::userID/timeline",
                            "protocol": "http",
                            "host": [
                                "localhost"
                            ],
                            "port": "8000",
                            "path": [
                                "api",
                                "v1",
                                "users",
                                "::userID",
                                "timeline"
                            ],
                            "variable": [
                                {
                                    "key": ":userID",
                                    "value": "a54cc0a5-2e06-4610-a6e5-d554a61f0dd7"
                                }
                            ]
                        }
                    },
                    "response": []
                }
            ]
        }
    ]
}

```
