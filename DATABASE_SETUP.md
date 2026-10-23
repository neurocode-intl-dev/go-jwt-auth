# Database Connection & User Table Setup - Complete ✓

## Summary

The project has been successfully configured to connect to PostgreSQL and automatically create the users table on startup.

## Changes Made

### 1. **Created Configuration Module** - [internal/config/config.go](internal/config/config.go)

- Loads environment variables from `.env` file
- Builds PostgreSQL connection URL from individual components
- Manages all application settings (DB, JWT, tokens)
- Provides defaults for all configuration values

### 2. **Updated Database Connection** - [internal/database/postgres.go](internal/database/postgres.go)

- Proper error handling for connection failures
- Added connection ping test
- New `CreateUserTable()` function that:
  - Creates `users` table with proper schema
  - Adds unique constraint on email
  - Adds index on email for faster queries
  - Includes `created_at` and `updated_at` timestamps
  - Uses `IF NOT EXISTS` to prevent errors on repeated runs

### 3. **Updated Server Initialization** - [cmd/server/main.go](cmd/server/main.go)

- Loads `.env` file at startup
- Loads configuration with proper error handling
- Initializes database connection
- Automatically creates users table
- Shows success messages for each step

### 4. **Enhanced User Model** - [internal/models/user.go](internal/models/user.go)

- Added `created_at` and `updated_at` timestamp fields
- Proper time.Time type for database mapping

## Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
```

## Environment Configuration (from .env)

```
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5431
DB_USER=postgres
DB_PASSWORD=admin@26
DB_NAME=go_db
JWT_SECRET=supersecretkey
REFRESH_SECRET=your-other-secret-here
ACCESS_TOKEN_EXPIRE=15
REFRESH_TOKEN_EXPIRE=43200
```

## Testing Results ✓

### Build Status

```
✓ Project builds successfully with: go build -o bin/server cmd/server/main.go
```

### Runtime Output

```
✓ Successfully connected to PostgreSQL database
✓ Users table created/verified successfully
✓ Server started on port 8080
✓ Echo framework running
```

### Database Verification

```
✓ Table exists: SELECT * FROM users;
✓ Columns verified: id, email, password, created_at, updated_at
✓ Constraints verified: UNIQUE email, PRIMARY KEY id
✓ Indexes verified: idx_users_email for fast email lookups
```

## How to Use

### Start the Server

```bash
make run
# or
go run cmd/server/main.go
```

The server will automatically:

1. Load environment variables
2. Connect to PostgreSQL
3. Create the users table (if it doesn't exist)
4. Start the Echo framework on port 8080

### API Endpoints Available

- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `GET /user/profile` - Get user profile (requires JWT token)

### Database Operations

The project now supports:

- User registration with password hashing
- User authentication
- JWT token generation
- Email-based user lookup with indexed queries

## Technical Details

### Packages Used

- `github.com/jackc/pgx/v5` - PostgreSQL driver
- `github.com/jackc/pgx/v5/pgxpool` - Connection pooling
- `github.com/joho/godotenv` - Environment variable loading
- `golang.org/x/crypto/bcrypt` - Password hashing
- `github.com/golang-jwt/jwt/v5` - JWT authentication
- `github.com/labstack/echo/v4` - Web framework

### Connection Details

- Uses pgxpool for connection pooling
- Connection string format: `postgres://username:password@host:port/database?sslmode=disable`
- Email field indexed for O(log n) lookups
- UNIQUE constraint prevents duplicate emails

## Next Steps (Optional)

1. **Add migrations**: Use a tool like `pressly/goose` for version control of schema changes
2. **Add validation**: Add email validation in the Register handler
3. **Add user profile endpoints**: Extend `/user/profile` endpoint with more data
4. **Add password reset**: Implement password reset functionality
5. **Add refresh token endpoint**: Implement token refresh endpoint
