Assumptions del Proyecto
Este documento describe los componentes de la arquitectura que, para los fines de este desafío, se asume que están completamente desarrollados y funcionales. El desarrollo se ha centrado en los servicios principales que cumplen con los requisitos explícitos del ejercicio (user-service, follow-service, post-service, timeline-service).

search-service
Se asume la existencia de un servicio de búsqueda completamente funcional.

Responsabilidad: Indexar el contenido de los posts y permitir búsquedas de texto completo sobre ellos.

Flujo Asumido: El post-ingest-service envía un evento a una cola de mensajes (post-ingest-queue) cada vez que se crea un post. Un search-consumer escucha esta cola, procesa el evento y actualiza el índice en un motor de búsqueda como Elasticsearch. El search-service expone una API para consultar este índice.

multimedia-service
Se asume la existencia de un servicio dedicado a la gestión de contenido multimedia.

Responsabilidad: Manejar la subida, el procesamiento (ej. redimensionamiento de imágenes) y la entrega de archivos multimedia asociados a los posts.

Flujo Asumido: Este servicio se integraría con el post-ingest-service para recibir archivos, los almacenaría en un sistema de almacenamiento de objetos como Amazon S3 y devolvería una URL para ser asociada al post.

post-processor y Actualización de Caché del Timeline
Se asume que el proceso de actualización de la caché del timeline (Redis) se realiza de forma asíncrona y eficiente.

Responsabilidad: Distribuir un nuevo post a los timelines de todos los seguidores del autor (estrategia fan-out-on-write).

Flujo Asumido: Tras la creación de un post, el post-processor (o un notificador) enviaría un evento. Un consumidor de este evento se encargaría de obtener la lista de seguidores del autor (desde el follow-service) y añadir el ID del nuevo post a la caché de timeline de cada uno de esos seguidores en el feed-service (Redis). Esto garantiza que la lectura del timeline sea extremadamente rápida.
