FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git build-base # gcc gettext musl-dev

# dependencies
COPY go.mod go.sum ./
RUN go mod download

# build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd

# run
FROM alpine:3.20 AS runner
WORKDIR /app
RUN adduser -D -u 10001 app
USER app
COPY --from=builder /app/bin/app /app/app

COPY docker.env .env
ENV ADDR=:8080 \
    STORAGE_MODE=postgres \
    POSTGRES_PG_DSN=""
EXPOSE 8080
ENTRYPOINT ["/app/app"]