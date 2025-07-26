# 🏦 HexaBank

## 📚 Overview

HexaBank is an example project & learning resource demonstrating hexagonal architecture (also known as Ports and Adapters) in Golang. It simulates a banking system using microservices within a monorepo. The goal is to show how to build robust, maintainable, and testable applications by separating business logic from external concerns like databases, UIs, and third-party services.

## 🛠️ Microservices

This monorepo contains multiple microservices. Microservices can be found inside the /services directory. Microservices include:

- `payment` - handling payment requests
- `fraud` - to determine if a payment is fraudulent

#### 💸 Payment Microservice

The `payment` microservice is an _HTTP API_ that accepts payment requests. Payments requests are sent to the `fraud` microservice (via grpc) to be cleared of fraud before being stored in a PostgreSQL database. Database migrations are managed using [Goose](https://github.com/pressly/goose). This service acts as the entry point for payment operations in the system.

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
- **`/domain`** --> Contains core business logic.
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

## 🔍 Observability Services

As part of this project, we aim to incorporate observability to enable quick debugging in the event of a system failure. Traces help us understand how requests flow through the microservices, metrics monitor resource usage for each service, and logs, collected using Loki, provide detailed insights into application behavior and errors.

#### 🔥 Prometheus _(Metrics)_

We use **Prometheus** in this project to gather metrics about our services.

**Custom Metrics**

- `gateway_payment_requests_total` - tracks the number of requests sent to the **gateway** service in total overtime.

**Config**

Location: `internal/observability/metrics/prometheus-config.yaml`

**Docker**

Below is how we define `prometheus` in the `docker-compose.yml`:

```bash
prometheus:
    image: prom/prometheus
    volumes:
      - ./internal/observability/metrics/prometheus-config.yaml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"

```

#### 🧬 Opentelemetry & Tempo _(Traces)_

We use **OpenTelemetry** and **Tempo** for **distributed tracing.**
This setup allows you to trace the full lifecycle of a request across services.

**Config**

Location: `internal/observability/tracing/tempo-config.yaml`

**Docker**

Below is how we define `tempo` in the `docker-compose.yml`:

```bash
tempo:
    image: grafana/tempo
    command: ["-config.file=/etc/tempo/tempo.yaml"]
    volumes:
        - ./internal/observability/tracing/tempo-config.yaml:/etc/tempo/tempo.yaml
    ports:
        - "4318:4318" # OPENTELEMETRY PORT
        - "3200:3200" # TEMPO PORT

```

#### 🪓 Loki & Promtail _(Logs)_

We use **Loki** and **Promtail** for **centralized logs.**
This is so we can store and visualize logs from all different services in one place.

_Loki_ - stores the logs & in used by Grafana for visualization

_Promtail_ - scrapes logs from the Docker container and pushes data to Loki

**Config(s)**

_Loki_ - `internal/observability/logging/loki-config.yaml`

_Promtail_ - `internal/observability/logging/promtrail-config.yaml`

**Docker**

Below is how we define `loki` and `promtail` in the `docker-compose.yml`:

```bash
loki:
    image: grafana/loki:2.9.5
    ports:
      - "3100:3100"
    volumes:
      - ./internal/observability/logging/loki-config.yaml:/etc/loki/config.yaml
    command: -config.file=/etc/loki/config.yaml

promtail:
    image: grafana/promtail:2.9.5
    volumes:
        - ./internal/observability/logging/promtrail-config.yaml:/etc/promtail/config.yml
        - /var/run/docker.sock:/var/run/docker.sock:ro
    command: -config.file=/etc/promtail/config.yml
    depends_on:
        - loki
```
