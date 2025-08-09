# Construye y levanta todos los servicios
run-all:
	@echo "Construyendo y levantando todos los servicios..."
	docker-compose up --build -d

# Detiene y elimina todos los contenedores
stop-all:
	@echo "Deteniendo todos los servicios..."
	docker-compose down --remove-orphans

# Muestra los logs de todos los servicios en modo 'follow'
logs:
	@echo "Mostrando logs de todos los servicios..."
	docker-compose logs -f
