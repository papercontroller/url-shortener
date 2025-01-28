FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /url-shortener .
COPY config /app/config
COPY storage /app/storage

EXPOSE 8000

CMD ["./url-shortener"]



