# API Gateway Service

The **API Gateway** serves as the entry point for all client-facing HTTP/REST requests in the OmniPay microservices ecosystem. It is responsible for routing RESTful requests to the appropriate downstream gRPC services.

## Responsibility
-   Expose REST API endpoints for clients.
-   Perform request validation and protocol translation (HTTP to gRPC).
-   Centralize authentication (to be implemented) and request logging.

## Tech Stack
-   **Language**: Go 1.24+
-   **Framework**: [go-kit](https://github.com/go-kit/kit)
-   **Router**: [gorilla/mux](https://github.com/gorilla/mux)
-   **Communication**: gRPC (Client to downstream services)

## Configuration
The service is configured via environment variables (see `configmap.yaml` in the deployment config):
-   `SERVICE_PORT`: TCP port for the HTTP server (default: `8080`).
-   `ACCOUNT_SERVICE_URL`: gRPC address of the Account Service.
-   `TRANSACTION_SERVICE_URL`: gRPC address of the Transaction Service.

## API Endpoints
| Method | Path                        | Description                 |
| ------ | --------------------------- | --------------------------- |
| `GET`  | `/v1/accounts/{id}/balance` | Retrieve account balance     |
| `POST` | `/v1/payments`              | Initiate a new payment       |
| `GET`  | `/v1/audits`                | Retrieve filtered audit logs |

## Local Development
To run the gateway locally:
```bash
go run main.go
```
Ensure that the downstream services (Account and Transaction) are running and accessible via the configured URLs.
