up:
	@echo "Starting all Docker images..."
	docker compose up -d
	@echo "Docker images started!!"

down:
	@echo "Stopping all Docker images..."
	docker compose down
	@echo "Done!!"

up_build:
	@echo "Stopping all Docker images if they are running..."
	docker compose down
	@echo "Building and starting Docker images"
	docker compose up --build -d
	@echo "Docker images are built and started!!"
	