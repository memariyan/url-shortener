
# URL Shortener Service

A high-performance URL shortener built with modern Go technologies, providing URL shortening, redirection, and tracking features. This project uses Go's powerful libraries.

## Features

- Generate short URLs for long links.
- Redirect users to the original URL using the shortened link.
- Persistent storage for URL mappings.
- Built-in observability for tracing and monitoring.
- Unit testing for robust code quality.

## Tech Stack

- **[Echo](https://echo.labstack.com/)**: Web framework for building RESTful APIs.
- **[Cobra](https://github.com/spf13/cobra)**: CMD runner
- **[GORM](https://gorm.io/)**: ORM library for database management.
- **[Viper](https://github.com/spf13/viper)**: Configuration management.
- **[Testify](https://github.com/stretchr/testify)**: Unit testing and mocking.
- **[Jaeger](https://www.jaegertracing.io/)**: Distributed tracing.
- **[OpenTelemetry](https://opentelemetry.io/)**: Observability and monitoring.

---

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/) (v1.20+ recommended)
- [Docker](https://www.docker.com/) (optional for external dependencies)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/memariyan/url-shortener.git
   cd url-shortener
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   cd ./docker
   sh start-services.sh
   ```

---

## Configuration

The project uses **Viper** for managing configurations. Configuration values can be set in the `config.yaml` file or via environment variables.

Example `config.yaml`:

```yaml
server:
   address: "http://localhost"
   port:  8001

mysql:
   host: "mysql"
   port: 3306
   username: "root"
   password: "root"
   db: "url_shortener"

redis:
   host: "redis"
   port: 6379

jaeger:
   host: "jaeger"
   port: 4318

worker:
   size: 10
```

---

## API Endpoints

| Method | Endpoint     | Description                  |
|--------|--------------|------------------------------|
| POST   | `/convert`   | Create a new short URL.      |
| GET    | `/:shortURL` | Redirect to the original URL |
| GET    | `/metrics`   | Check service metrics.       |

---

## Testing

The project uses **Testify** for unit testing. Run the tests with:

```bash
go test ./...
```

---

## Observability

### Tracing with Jaeger

1. Access the Jaeger UI at [http://localhost:16686](http://localhost:16686).

2. The service automatically sends traces to the Jaeger endpoint specified in `config.yaml`.

### Prometheus metrics

1. Access the Grafana UI at [http://localhost:3003](http://localhost:3003).

2. You can see all metrics has been exposed by application in /metrics endpoint

---
