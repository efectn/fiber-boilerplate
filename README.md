# Fiber Boilerplate
[![Go Reference](https://pkg.go.dev/badge/github.com/efectn/fiber-boilerplate.svg)](https://pkg.go.dev/github.com/efectn/fiber-boilerplate)

Simple and scalable boilerplate to build powerful and organized REST projects with [Fiber](https://github.com/gofiber/fiber). 

## Directory Structure

```
├── app
│   ├── database
│   │   ├── schema
│   │   │   └── article.go
│   │   └── seeder
│   │       └── article_seeder.go
│   ├── middleware
│   │   ├── register.go
│   │   └── token
│   │       └── token.go
│   ├── module
│   │   └── article
│   │       ├── article_module.go
│   │       ├── controller
│   │       │   ├── article_controller.go
│   │       │   ├── article_controller_mock.go
│   │       │   └── controller.go
│   │       ├── repository
│   │       │   ├── article_repository.go
│   │       │   └── article_repository_mock.go
│   │       ├── request
│   │       │   └── article_request.go
│   │       └── service
│   │           ├── article_service.go
│   │           └── article_service_mock.go
│   └── router
│       └── api.go
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
├── internal
│   └── bootstrap
│       ├── database
│       │   └── database.go
│       ├── logger.go
│       └── webserver.go
├── LICENSE
├── Makefile
├── README.md
├── storage
│   ├── ascii_art.txt
│   ├── private
│   │   └── example.html
│   ├── private.go
│   └── public
│       └── example.txt
└── utils
    ├── config
    │   └── config.go
    ├── response
    │   ├── response.go
    │   └── validator.go
    └── utils.go
```

## Usage
You can run that commands to run project:

```go mod download```

```go run cmd/example/main.go``` or ```air -c .air.toml``` if you want to use air

### Docker
```shell
docker-compose build
docker-compose up

CUSTOM="Air" docker-compose up # Use with Air
```

## Tech Stack
- [Go](https://go.dev)
- [PostgreSQL](https://www.postgresql.org)
- [Docker](https://www.docker.com/)
- [Fiber](https://github.com/gofiber/fiber)
- [Ent](https://github.com/ent/ent)
- [Fx](https://github.com/uber-go/fx)
- [Zerolog](https://github.com/rs/zerolog)
- [GoMock](https://github.com/golang/mock)

## To-Do List
- [x] More error-free logging.
- [x] Add makefile to make something shorter.
- [x] Introduce repository pattern.
- [ ] Add unit tests.
- [x] Add mocking with GoMock.

## License
fiber-boilerplate is licensed under the terms of the **MIT License** (see [LICENSE](LICENSE)).
