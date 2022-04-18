# Makefile for some commands.

# Generate Ent boilerplate
generate:
	cd cmd/example
	go generate .

# Run application outside Docker
run:
	go run cmd/example/main.go

# Run application inside Docker
docker:
	docker-compose up

# Run application inside Docker with Air
air:
	CUSTOM="Air" docker-compose up --build example

# Tidy
tidy:
	go mod tidy

# Stop application
stop:
	docker-compose down