# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/server ./cmd/app && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/seed ./cmd/seed

RUN GOBIN=/app/bin go install github.com/pressly/goose/v3/cmd/goose@v3.24.0

FROM alpine:3.21

RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup && \
    apk add --no-cache curl

WORKDIR /app

COPY --from=builder /app/bin/server /app/bin/server
COPY --from=builder /app/bin/seed /app/bin/seed
COPY --from=builder /app/bin/goose /app/bin/goose

COPY migrations ./migrations
COPY seed ./seed/

RUN chmod +x /app/bin/server /app/bin/seed /app/bin/goose && \
    chown -R appuser:appgroup /app

COPY deployments/docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh && \
    chown appuser:appgroup /entrypoint.sh

USER appuser

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]