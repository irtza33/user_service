# user-service

This project is a gRPC-based user service that interacts with a PostgreSQL database. It provides methods for managing users, including retrieving user information, creating new users, and deleting existing users. The service is designed following clean architecture principles and best practices.

## Project Structure

- **cmd/server/main.go**: Entry point of the application, initializes the gRPC server.
- **internal/config/config.go**: Configuration settings for database connections and logging.
- **internal/domain/user.go**: Defines the User struct and related methods.
- **internal/repository/postgres/user_store.go**: Implements the UserStore interface for PostgreSQL.
- **internal/repository/interfaces.go**: Defines the UserStore interface with methods for user operations.
- **internal/service/user_service.go**: Implements the user service business logic.
- **internal/metrics/prometheus.go**: Sets up Prometheus metrics for monitoring.
- **pkg/logger/logger.go**: Provides logging functionality for requests, responses, and errors.
- **pkg/database/postgres.go**: Handles database connections and queries.
- **proto/user/user.proto**: Defines the gRPC service and message types.
- **proto/user/user.pb.go**: Generated Go code from the user.proto file.
- **migrations/001_create_users_table.sql**: SQL script to create the users table.
- **go.mod**: Go module definition with dependencies.
- **go.sum**: Checksums for dependencies.
- **Makefile**: Build instructions and commands.
- **Dockerfile**: Docker image definition for the application.

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd user-service
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run migrations**:
   Ensure your PostgreSQL database is running and execute the migration script:
   ```bash
   psql -U <username> -d <database> -f migrations/001_create_users_table.sql
   ```

4. **Start the server**:
   ```bash
   go run cmd/server/main.go
   ```

## Usage

Once the server is running, you can interact with the user service using gRPC clients. The service supports the following methods:

- `GetUser(user_id)`: Retrieves the user's name by user_id.
- `CreateUser(name)`: Creates a new user and returns the user_id.
- `DeleteUser(user_id)`: Deletes a user and returns a confirmation of deletion.

## Logging and Monitoring

All gRPC requests and responses are logged. Errors are logged with appropriate security levels. Prometheus metrics are available for monitoring system stats, including request counts and database connections.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.