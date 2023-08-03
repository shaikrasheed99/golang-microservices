up:
	@echo "Starting all Docker images..."
	docker compose up -d
	@echo "Docker images started!!"

down:
	@echo "Stopping all Docker images..."
	docker compose down
	@echo "Done!!"