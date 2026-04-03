# RabbitMQ Microservices Dummy

A demo app that shows how microservices can talk using RabbitMQ.

## What this project does
This repo simulates a basic order pipeline:

1. **Server service** gets an order request.
2. It publishes order events to RabbitMQ.
3. **Inventory worker** checks stock.
4. **Payment worker** processes payment.
5. **Notification worker** sends a final message.

## Why this exists
This is a **learning project**.
It is made to practice architecture ideas, not to be production-ready.

## Tech used
- Go
- RabbitMQ
- HTTP handlers
- Pub/Sub messaging

## Project structure
- `cmd/server` – API/server entrypoint
- `cmd/inventory` – inventory worker service
- `cmd/payment` – payment worker service
- `cmd/notification` – notification worker service
- `internal/handlers` – HTTP endpoints and request flow
- `internal/workers` – worker logic
- `internal/rabbitmq` – RabbitMQ publish/subscribe helper

## Quick start
1. Start RabbitMQ.
2. Run the services (server + workers).
3. Send an order request to the server.
4. Watch the workers process events.