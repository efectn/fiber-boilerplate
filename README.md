# Fiber Boilerplate
Simple and scalable boilerplate to build powerful and organized REST projects with [Fiber](https://github.com/gofiber/fiber). 

Structure inspired by [project-layout](https://github.com/golang-standards/project-layout).

## Directory Structure:

```
├── build
│   ├── Dockerfile
│   └── DockerfileAir
├── cmd
│   └── example
│       ├── generate.go
│       └── main.go
├── config
│   └── example.toml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── LICENSE
├── pkg
│   ├── controllers
│   │   ├── article_controller.go
│   │   └── contorller.go
│   ├── database
│   │   ├── database.go
│   │   ├── schemas
│   │   │   └── article.go
│   │   └── seeds
│   │       └── article_seeder.go
│   ├── helpers
│   │   ├── logger.go
│   │   └── webserver.go
│   ├── middlewares
│   │   ├── register.go
│   │   └── token
│   │       └── token.go
│   ├── requests
│   │   └── article_request.go
│   ├── router
│   │   └── api.go
│   ├── services
│   │   ├── article_service.go
│   │   └── services.go
│   └── utils
│       ├── config
│       │   └── config.go
│       ├── response
│       │   ├── response.go
│       │   └── validator.go
│       └── utils.go
├── README.md
└── storage
    ├── ascii_art.txt
    ├── private
    │   └── example.html
    ├── private.go
    └── public
        └── example.txt
```

## Usage:
You can run that commands to run project:

```go mod download```

```go run cmd/example/main.go``` or ```air -c .air.toml``` if you want to use air

### Docker:
```shell
docker-compose build
docker-compose up

CUSTOM="Air" docker-compose up # Use with Air
```

## Tech Stack:
- [Go](https://go.dev)
- [PostgreSQL](https://www.postgresql.org)
- [Docker](https://www.docker.com/)
- [Fiber](https://github.com/gofiber/fiber)
- [Fx](https://github.com/uber-go/fx)
- [Zerolog](https://github.com/rs/zerolog)