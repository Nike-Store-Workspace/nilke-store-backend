
# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -mod=readonly -o /app/main ./cmd/api

# Stage 2
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN adduser -D appuser

COPY --from=builder /app/main .

USER appuser

EXPOSE 8080

CMD ["./main"]