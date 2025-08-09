FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./app

FROM scratch

COPY --from=builder /app/main /app

EXPOSE 8080

# El comando que se ejecutará cuando el contenedor inicie.
ENTRYPOINT ["/app"]
