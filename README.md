# 🏦 HexaBank

## 📚 Overview

HexaBank is an example project demonstrating hexagonal architecture (also known as Ports and Adapters) in Golang. It simulates a banking system using microservices within a monorepo. The goal is to show how to build robust, maintainable, and testable applications by separating business logic from external concerns like databases, UIs, and third-party services.

## 🛠️ Services

This monorepo contains multiple microservices. Currently, the main service is:

### 💸 Payment Service

The `payment` service is an HTTP API that accepts payment requests, validates them, and stores them in a PostgreSQL database. Database migrations are managed using [Goose](https://github.com/pressly/goose). This service acts as the entry point for payment operations in the system.

## 🗂️ Project Structure

HexaBank follows a hexagonal architecture, emphasizing a clear separation between core domain logic and external dependencies. The folder structure is:

```
services/
└── payment/
    ├── cmd/          # Application entry points (main.go, Dockerfile)
    ├── domain/       # Core business logic (the "hexagon")
    │   ├── model/    # Domain entities and value objects
    │   ├── port/     # Interfaces defining what the domain needs (ports)
    │   └── service/  # Domain services implementing business rules
    ├── adapters/     # Implementations of ports (adapters)
    │   ├── http/     # HTTP API adapter (inbound)
    │   └── postgres/ # PostgreSQL database adapter (outbound)
    ├── errors/       # Custom error definitions
    └── migrations/   # Database migration scripts
```

- **`cmd/`**: Application entry point (`main.go`, Dockerfile). Wires dependencies and starts the service.
- **`domain/`**: The "hexagon" with core business logic.
  - **`model/`**: Domain entities and value objects.
  - **`port/`**: Interfaces (ports) defining contracts for the domain.
  - **`service/`**: Business rules and orchestration.
- **`adapters/`**: Implements the `port` interfaces, connecting the domain to external tech.
  - **`http/`**: Inbound HTTP API adapter.
  - **`postgres/`**: Outbound PostgreSQL adapter.
- **`errors/`**: Custom error types.
- **`migrations/`**: SQL scripts for database migrations.

## Running the Project 🚀

To run HexaBank, use Docker and Docker Compose. Make sure Docker Desktop is installed and running.

1. **Build and run the services:**

   ```bash
   docker-compose up --build
   ```

   This will build images and start containers, including the `payment` service and its dependencies (like PostgreSQL).

2. **Access the services:**
   The `payment` service API is usually available at `http://localhost:8080` (or as configured in `docker-compose.yml`).

3. **Stop the services:**
   To stop and clean up:
   ```bash
   docker-compose down
   ```


<img width="1024" height="1536" alt="ChatGPT Image Jul 23, 2025, 08_13_51 PM" src="https://github.com/user-attachments/assets/aa224d0e-34d6-416a-bd68-752646af83a0" />
