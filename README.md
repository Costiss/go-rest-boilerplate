# Golang REST API Boilerplate

This repository is a boilerplate app for a Golang REST API with modular architecture.
It has been created with the following features:

- [x] Atlasgo (Database Migrations)
- [x] Gin (Web Framework)
- [x] GORM (Database ORM)
- [x] Fx (Dependency Injection)
- [ ] Structured Logging
- [ ] Unit Testing
- [x] Swagger
- [ ] Telemetry

## Project Structure

```
├── internal
│   ├── app.go                      # Application Module
│   ├── domain                      # Domain layer
│   │   └── user
│   │       ├── module.go           # Module definition
│   │       ├── router.go
│   │       └── service.go
│   ├── infra                       # Infrastructure layer
│   │   ├── database.go
│   │   └── server.go
│   └── model                       # GORM model definitions
│       └── user.go
├── migrations                      # Database migration files
│   ├── 20240710214933.sql
│   └── atlas.sum
├── main.go                         # App entry point
├── atlas.hcl                       # Atlasgo config
├── docker-compose.yml
├── go.mod
├── go.sum
├── .air.toml                       # Air config
├── Makefile
└── README.md
```
