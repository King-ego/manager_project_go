## Technical Highlights

- **Clean Architecture** with strict dependency rule
- **DDD-inspired** design with rich domain models
- **CQRS pattern** for read/write operations
- **Unit of Work** pattern for complex transactions
- **OpenTelemetry** instrumentation
- **Property-based testing** for core domain logic

## Getting Started

### Prerequisites

- Go 1.21+
- Docker (for optional PostgreSQL container)
- Make (for build automation)

### Installation

```bash
# Clone repository
git clone https://github.com/yourrepo/project-mgmt-api.git

# Start dependencies (PostgreSQL, Prometheus)
make infra-up

# Run application
make run