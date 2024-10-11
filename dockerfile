# Etapa de build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Etapa final
FROM alpine:latest

COPY --from=builder /app/main /main
COPY .env .env

EXPOSE 8081

ENTRYPOINT ["/main"]
