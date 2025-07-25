# 🏦 HexaBank

## 📚 Overview

HexaBank is an example project & learning resource demonstrating hexagonal architecture (also known as Ports and Adapters) in Golang. It simulates a banking system using microservices within a monorepo. The goal is to show how to build robust, maintainable, and testable applications by separating business logic from external concerns like databases, UIs, and third-party services.

## 🛠️ Microservices

This monorepo contains multiple microservices. Microservices can be found inside the /services directory. Microservices include:

- `payment` - handling payment requets
- `fraud` - to determine if a payment is fraudulent

#### 💸 Payment Microservice

The `payment` microservice is an _HTTP API_ that accepts payment requests. Payments frequest are sent to the `fraud` microservice (via grpc) to be cleared of fraud before being stored in a PostgreSQL database. Database migrations are managed using [Goose](https://github.com/pressly/goose). This service acts as the entry point for payment operations in the system.

Below is the list of available endpoints:

- **POST** `/api/v1/payments` --> Create a payment request

- **GET** `/api/v1/payments/:id` --> Fetch a payment

#### 🧬 Fraud Microservice

The `fraud` microservice is a _GRPC API_ that accepts payments requests, and clears them of fraud.
A payment is considered **fraudulent** if the amount is a number inside the **Fibonacci Sequence** _(e.g. 1, 2, 3, 5, 8, etc.)_

## 🗂️ Microservice Hexagonal Structure

Microservices found in this project follow a hexagonal architecture, emphasizing a clear separation between core domain logic and external dependencies.

Below is an explanation of the various folders and an example of the adapter used inside the `payment` microservice:

- **`/cmd`** --> Application entry point (`main.go`).
- **`/domain`** --> Conatins core business logic.
  - **`/model`** --> Domain entities and value objects.
  - **`/port`** --> Interfaces (ports) defining contracts for the domain.
  - **`/service`** --> Business logic.
- **`/adapters`** --> Implements the `port` interfaces, connecting the domain to external tech.
  - **`/http`** --> Inbound HTTP API adapter.
  - **`/postgres`** --> Outbound PostgreSQL adapter.
    - **`/migrations`** --> SQL migration files for Postgres.

## 🚀 Running the Project

To run HexaBank, use Docker and Docker Compose. Make sure Docker Desktop is installed and running.

**Build and run the services:**

```bash
docker-compose up --build
```

The following services will be available:

- `payment` --> `localhost:8080`
- `fraud` --> `localhost:50052` (GRPC Service)
