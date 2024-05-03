# Gateway

API gateway for microservices usign go.

## Structure

```
.
├── go.work
├── go.work.sum
├── README.md
├── svcApi                  <--- Api proxy service
│   ├── default.conf
│   ├── Dockerfile
│   ├── docs                <--- swagger documentation
│   ├── go.mod
│   ├── go.sum
│   ├── middleware.go
│   └── proxy.go
├── svcOrder                <--- Order service
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── handlers.go
│   └── order.go
├── svcProduct              <--- Productt service
│   ├── application
│   │   └── application.go
│   ├── database
│   │   └── client.go
│   ├── Dockerfile
│   ├── entities
│   │   └── product.go
│   ├── go.mod
│   ├── go.sum
│   ├── handlers.go
│   └── main.go
└── svcUser                 <--- User service
    ├── application
    │   └── application.go
    ├── database
    │   └── client.go
    ├── Dockerfile
    ├── entities
    │   └── user.go
    ├── go.mod
    ├── go.sum
    ├── handlers.go
    └── main.go
```
## Setup

To build and run source code, run below command

```
docker compose up -d
```
Swagger documentation is available at `http://localhost:8082/`

To stop and clean docker resources, run below command

```
docker compose down --rmi local
```