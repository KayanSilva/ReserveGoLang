# Etapa 1: build da aplicação
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o /app/api

# Etapa de execução
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api .
COPY --from=builder /app/dados ./dados
EXPOSE 8080
CMD ["./api"]