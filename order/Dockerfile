FROM golang:1.19 AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /main

# # Run the tests in the container
# FROM builder AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:3.14

WORKDIR /

COPY --from=builder /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]