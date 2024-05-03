FROM golang:1.19 AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /api

# Deploy the application binary into a lean image
FROM nginx

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /api /api
COPY --from=builder /app/docs /usr/share/nginx/html
COPY --from=builder /app/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 8081 8082

RUN echo "#!/bin/sh" >> /bin/start.sh \
 && echo "service nginx restart &" >> /bin/start.sh \
 && echo "/api" >> /bin/start.sh \
 && chmod 755 /bin/start.sh

ENTRYPOINT ["/bin/start.sh"]