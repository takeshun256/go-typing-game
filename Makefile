.PHONY: build run clean docker-build docker-run docker-clean

# Compile main.go
build:
	go build main.go

# Run the game
run: 
	go run main.go

# Clean up compiled files
clean:
	rm -f main

# Build the Docker image and start the container
docker-build:
	docker-compose -f docker-compose.prod.yml up -d --build

# Enter the container and start the game
docker-run:
	docker-compose -f docker-compose.prod.yml exec app go run main.go

# Stop and remove the Docker container
docker-clean:
	docker-compose -f docker-compose.prod.yml down

