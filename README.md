# Go RabbitMQ Example

A simple RabbitMQ example in Go

## Running locally

- Run RabbitMQ with Docker

```bash
docker compose up -d
```

1. Install dependencies

```bash
go mod tidy
```

2. Run consumer

```bash
go run consumer/main.go
```

3. Publish messages

```bash
go run publisher/main.go <message> <routing_key(optional)>
```
