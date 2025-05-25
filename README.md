# Go E-commerce Backend

A modern e-commerce backend service built with Go, featuring user management, product catalog, and order processing.

## Prerequisites

- Go 1.23.5 or later
- MySQL database (for production) or any SQL database for development
- Basic knowledge of Go programming language
- Familiarity with web development concepts and RESTful APIs

## Project Setup

### 1. Clone and Install Dependencies

```bash
# Clone the repository
git clone https://github.com/xudong7/ecom.git
cd ecom

# Install Go dependencies
go mod tidy
```

### 2. Environment Configuration

Create a `.env` file in the root directory with the following variables:

```env
# Database Configuration
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_ADDRESS=localhost:3306
DB_NAME=ecom_db

# Server Configuration
PUBLIC_HOST=localhost
PORT=8080
```

### 3. Database Setup

```bash
# Install golang-migrate tool (if not already installed)
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create a new migration (example)
make migration add-new-table-name

# Run all pending migrations
make migrate-up

# Rollback the last migration (if needed)
make migrate-down
```

## Running the Project

### Development

```bash
# Clean previous builds, tidy dependencies, build and run
make run

# Or run individual commands:
make clean    # Clean build artifacts
make tidy     # Tidy Go modules
make build    # Build the application
make test     # Run tests
```

The server will start on `http://localhost:8080`

### Available Make Commands

| Command | Description |
|---------|-------------|
| `make run` | Clean, tidy, build and run the server |
| `make build` | Build the application binary |
| `make test` | Run all tests |
| `make tidy` | Tidy Go module dependencies |
| `make clean` | Remove build artifacts |
| `make migration <name>` | Create a new database migration |
| `make migrate-up` | Apply all pending migrations |
| `make migrate-down` | Rollback the last migration |

## Project Structure

```
├── cmd/
│   ├── main.go                 # Application entry point
│   ├── api/                    # API server and routes
│   └── migrate/                # Database migration tools
│       └── migrations/         # SQL migration files
├── config/                     # Configuration management
├── db/                         # Database connection and setup
├── service/                    # Business logic services
│   ├── auth/                   # Authentication services
│   └── user/                   # User management services
├── types/                      # Type definitions and models
├── utils/                      # Utility functions
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
├── Makefile                    # Build and development commands
└── README.md                   # Project documentation
```

## API Endpoints

The application provides RESTful API endpoints for:

- **User Management**: Registration, authentication, profile management
- **Product Catalog**: Product listing, details, search
- **Order Processing**: Cart management, order creation, order history

## Database Schema

The project includes the following main tables:

- `users` - User accounts and authentication
- `products` - Product catalog
- `orders` - Order information
- `order_items` - Individual items within orders

## Development Guidelines

### Creating New Migrations

```bash
# Create a new migration for adding a table
make migration add-table-name

# Create a new migration for modifying existing data
make migration update-table-name

# Always create both up and down migrations for reversibility
```

### Testing

```bash
# Run all tests
make test

# Run tests with verbose output
go test -v ./...

# Run tests for specific package
go test -v ./service/user
```

## Topics Covered

This project demonstrates:

This project demonstrates:

1. Setting up a Go web server with Gorilla Mux
2. Handling HTTP requests and responses with proper routing
3. Working with MySQL database and migrations
4. Implementing user authentication and password hashing
5. Building RESTful API endpoints
6. Database connection management and SQL operations
7. Project structure and organization best practices
8. Testing Go web applications
9. Environment configuration management
10. Build automation with Makefiles

## Troubleshooting

### Common Issues

**Migration tool not found:**

```bash
# Make sure golang-migrate is installed
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Check if $GOPATH/bin is in your PATH
echo $PATH | grep $GOPATH/bin
```

**Database connection issues:**

- Verify MySQL is running
- Check database credentials in `.env` file
- Ensure database exists before running migrations

**Build issues:**

```bash
# Clean and rebuild
make clean
make build
```

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Gorilla Mux Router](https://github.com/gorilla/mux)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- [Go Web Examples](https://gowebexamples.com/)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run tests and ensure they pass
6. Submit a pull request

## License

This project is for educational purposes.
