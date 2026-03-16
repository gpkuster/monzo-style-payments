# banking-style-payments

A simplified payments backend written in Go, inspired by the architecture used by fintech companies like Monzo and Stripe.

## Goal

The goal of this project is to explore how modern payment systems are built using Go and a microservice-based architecture.

## Current status

- Basic HTTP server in Go
- `/payments` endpoint to create a payment + Idempotent payment API (header)
- `/transactions` endpoint
- Project structure for services and handlers
- Ledger service with double-entry accounting
- Docker setup
- Service-to-service using HTTP API call (ledger_client.go)
- 
## Planned features
- Service-to-service communication using [Typhon](https://github.com/monzo/typhon)
- PostgreSQL persistence

## Run the service
### Using docker
```bash
docker compose up --build
```

### Using Go in the terminal
Run the payments service
```bash
go run ./cmd/payments-service
```
Run the ledger service
```bash
go run ./cmd/ledger-service
```
