# Gateway

API gateway for microservices usign go.

## Structure

```
.
├── book                    <-- Book microservice
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── handlers.go
│   └── main.go
├── gateway                 <-- Api Gateway (api access point)
│   ├── Dockerfile
│   ├── gateway.go
│   ├── go.mod
│   ├── go.sum
│   └── middleware.go
├── go.work
├── order                   <-- Order microservice
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── handlers.go
│   └── main.go
└── README.md
```
## Setup

To build and run source code, run below command
```
docker-compose up -d
```

To stop and clean docker resources, run below command
```
docker-compose down --rmi all
```