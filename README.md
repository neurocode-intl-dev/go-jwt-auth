## What Is `bin/`?

`bin/` is used for compiled binary files.

When you run:

```bash
make build
```

Go builds the application and writes the executable here:

```text
bin/server
```

You can then start the compiled app with:

```bash
./bin/server
```

In many Go projects, `bin/` is not hand-written source code. It is generated
output from the build step.

## What Is `go.mod`?

`go.mod` defines this project as a Go module.

It contains:

- The module path: `github.com/rahulreghupdm/tashub-project`
- The Go version used by the project
- Direct dependencies such as Gin, JWT, pgx, godotenv, and bcrypt
- Indirect dependencies required by those packages

Common command:

```bash
go mod tidy
```

This cleans the dependency list by adding missing dependencies and removing
unused ones.

## What Is `go.sum`?

`go.sum` stores cryptographic checksums for downloaded Go dependencies.

It helps Go verify that dependency versions have not changed unexpectedly.

You normally do not edit `go.sum` manually. Go updates it automatically when
you run commands such as:

```bash
go mod tidy
go test ./...
go build ./...
```

## What Is `Makefile`?

The `Makefile` contains shortcut commands so you do not need to type long Go
commands every time.

Available commands:

```bash
make run
```

Runs the server directly:

```bash
go run cmd/server/main.go
```

```bash
make build
```

Builds the server binary:

```bash
go build -o bin/server cmd/server/main.go
```

```bash
make test
```

Runs all tests:

```bash
go test ./...
```

```bash
make tidy
```

Cleans Go dependencies:

```bash
go mod tidy
```

```bash
make fmt
```

Formats Go code:

```bash
go fmt ./...
```

```bash
make start
```

Starts the server. In this project it does the same thing as `make run`.

# Go JWT Auth API


This project is a small JWT authentication API written in Go. It uses the
Gin web framework for HTTP routing, PostgreSQL for storing users, bcrypt for
password hashing, and JSON Web Tokens for authentication.

The API currently supports:

- User registration
- User login
- Access token and refresh token generation
- Protected profile route using JWT middleware
- Automatic PostgreSQL connection setup
- Automatic `users` table creation on server startup

## Tech Stack

- Go: backend programming language
- Gin: HTTP web framework
- PostgreSQL: database
- pgx/pgxpool: PostgreSQL driver and connection pool
- bcrypt: password hashing
- golang-jwt/jwt: JWT creation and validation
- godotenv: loads environment variables from `.env`

## Folder Structure

```text
go-jwt-auth/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Loads app, database, and token config
│   ├── database/
│   │   └── postgres.go          # PostgreSQL connection and table creation
│   ├── handlers/
│   │   └── auth_handler.go      # HTTP handlers for register and login
│   ├── middleware/
│   │   └── jwt_middleware.go    # JWT protection for private routes
│   ├── models/
│   │   └── user.go              # User struct/model
│   ├── repository/
│   │   └── user_repository.go   # Database queries for users
│   ├── routes/
│   │   └── auth_routes.go       # Auth route registration
│   ├── services/
│   │   └── auth_service.go      # Business logic for auth
│   └── utils/
│       └── jwt.go               # Access and refresh token generation
├── bin/
│   └── server                   # Compiled server binary after `make build`
├── DATABASE_SETUP.md            # Extra notes about database setup
├── GIN_MIGRATION.md             # Notes from migration to Gin
├── go.mod                       # Go module name and dependency list
├── go.sum                       # Dependency checksum lock file
├── Makefile                     # Shortcut commands for running/building
└── README.md                    # Project documentation
```

## What Each Main File Does

### `cmd/server/main.go`

This is where the application starts.

It does these things in order:

1. Loads environment variables from `.env` using `godotenv`.
2. Loads configuration from `internal/config`.
3. Connects to PostgreSQL.
4. Creates the `users` table if it does not already exist.
5. Creates a Gin router.
6. Registers public auth routes.
7. Registers protected user routes.
8. Starts the HTTP server.

### `internal/config/config.go`

This file reads configuration values from environment variables.

Examples:

- `APP_PORT`
- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `JWT_SECRET`
- `REFRESH_SECRET`
- `ACCESS_TOKEN_EXPIRE`
- `REFRESH_TOKEN_EXPIRE`

It also provides default values when environment variables are missing.

### `internal/database/postgres.go`

This file manages PostgreSQL.

It creates a global database connection pool using `pgxpool`, checks that the
database is reachable with `Ping`, and creates the `users` table.

The table contains:

- `id`
- `email`
- `password`
- `created_at`
- `updated_at`

The email column is unique, so the same email cannot be registered twice.

### `internal/models/user.go`

This defines the Go `User` struct used by the application.

It represents a user row from the database:

```go
type User struct {
    ID        int64
    Email     string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### `internal/repository/user_repository.go`

The repository layer talks directly to the database.

It currently has:

- `CreateUser(email, hashedPassword string)`
- `GetUserByEmail(email string)`

Keeping SQL queries here makes the rest of the app cleaner.

### `internal/services/auth_service.go`

The service layer contains business logic.

For registration, it hashes the password with bcrypt and stores the user.

For login, it:

1. Finds the user by email.
2. Compares the plain password with the hashed password.
3. Generates an access token.
4. Generates a refresh token.

### `internal/handlers/auth_handler.go`

Handlers receive HTTP requests and send HTTP responses.

This file contains:

- `Register(c *gin.Context)`
- `Login(c *gin.Context)`

The handlers read JSON from the request body, call the service layer, and
return JSON responses.

### `internal/routes/auth_routes.go`

This file connects URL paths to handler functions.

Current auth routes:

```text
POST /auth/register
POST /auth/login
```

### `internal/utils/jwt.go`

This file generates JWT tokens.

- Access tokens expire after 15 minutes.
- Refresh tokens expire after 30 days.
- Access tokens are signed with `JWT_SECRET`.
- Refresh tokens are signed with `REFRESH_SECRET`.

### `internal/middleware/jwt_middleware.go`

Middleware runs before protected route handlers.

The auth middleware:

1. Reads the `Authorization` header.
2. Removes the `Bearer ` prefix if present.
3. Parses and validates the JWT.
4. Extracts `user_id` from the token claims.
5. Stores `user_id` in the Gin context.
6. Allows the request to continue with `c.Next()`.

If the token is missing or invalid, the request returns `401 Unauthorized`.


## About Gin

Gin is a fast HTTP web framework for Go. It is used in this project to define
routes, read request JSON, send JSON responses, group routes, and attach
middleware.

Gin is similar in style to Express.js.

Express example:

```js
app.get("/profile", (req, res) => {
  res.json({ message: "ok" });
});
```

Gin example:

```go
r.GET("/profile", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "ok"})
})
```

Important Gin concepts used in this project:

- `gin.New()` creates a new router.
- `e.Group("/auth")` creates a route group.
- `rg.POST("/login", handlers.Login)` registers a POST route.
- `c.ShouldBindJSON(&req)` reads JSON request body into a Go struct.
- `c.JSON(statusCode, gin.H{...})` sends a JSON response.
- `gin.HandlerFunc` is used for middleware.
- `c.Next()` continues to the next handler after middleware logic.
- `c.Set("user_id", value)` stores data in the request context.
- `c.Get("user_id")` reads data from the request context.

## Request Flow

### Register

```text
POST /auth/register
        |
        v
routes/auth_routes.go
        |
        v
handlers/auth_handler.go
        |
        v
services/auth_service.go
        |
        v
repository/user_repository.go
        |
        v
PostgreSQL users table
```

### Login

```text
POST /auth/login
        |
        v
handler reads email/password
        |
        v
service checks user and password
        |
        v
utils/jwt.go creates tokens
        |
        v
JSON response with access_token and refresh_token
```

### Protected Profile

```text
GET /user/profile
Authorization: Bearer <access_token>
        |
        v
middleware/jwt_middleware.go validates token
        |
        v
main.go profile handler returns user_id
```

## Environment Variables

Create a `.env` file in the project root.

Example:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
JWT_SECRET=change-this-access-secret
REFRESH_SECRET=change-this-refresh-secret
ACCESS_TOKEN_EXPIRE=15
REFRESH_TOKEN_EXPIRE=43200
```

For local learning this is enough. For production, use strong secrets and do
not commit real `.env` values.

## How to Run

Install dependencies:

```bash
go mod tidy
```

Start PostgreSQL and make sure the database values in `.env` are correct.

Run the server:

```bash
make run
```

Or:

```bash
go run cmd/server/main.go
```

The server starts on:

```text
http://localhost:8080
```

If `APP_PORT` is different, use that port instead.

## Build and Run Binary

Build:

```bash
make build
```

Run:

```bash
./bin/server
```

## API Endpoints

### Register

```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

Success response:

```json
{
  "message": "User registered"
}
```

### Login

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

Success response:

```json
{
  "access_token": "...",
  "refresh_token": "..."
}
```

### Profile

```bash
curl -X GET http://localhost:8080/user/profile \
  -H "Authorization: Bearer <access_token>"
```

Success response:

```json
{
  "user_id": 1
}
```

## Database Schema

The app creates this table automatically:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
```

## Useful Commands

```bash
make run      # Run the API
make build    # Build binary into bin/server
make test     # Run tests
make tidy     # Clean dependencies
make fmt      # Format Go code
make start    # Run the API
```

## Notes

- `internal/` is a Go convention for private application code. Code outside
  this module cannot import packages from this folder.
- Passwords are stored as bcrypt hashes, not plain text.
- Access tokens are used for protected API routes.
- Refresh tokens are generated, but there is not yet a refresh-token endpoint.
- The profile route currently returns only the authenticated `user_id`.
