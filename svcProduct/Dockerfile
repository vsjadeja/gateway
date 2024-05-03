FROM golang:1.20 AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /svcProduct

# # Run the tests in the container
# FROM builder AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest

WORKDIR /

COPY --from=builder /svcProduct /svcProduct

EXPOSE 8080

ENTRYPOINT ["/svcProduct"]