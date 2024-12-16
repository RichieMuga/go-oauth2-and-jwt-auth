# Motivation

The Go Gin Template aims to provide a robust and well-structured starting point for building scalable and maintainable web applications using Go and the Gin framework. I personally will be using this as microservices starterpack.

### Features

- [x] Swagger Documentation: Automatically generate and serve API documentation using Swagger.
- [x] Logging: Utilize slog for structured and JSON-formatted logging.
- [x] Testing: Implement tests with the testify library for comprehensive unit and integration testing.
- [x] ORM Integration: Leverage gorm as the Object-Relational Mapping (ORM) tool for seamless database interactions.
- [x] SQLite Support: Use SQLite for a lightweight and efficient database solution.

## Installation Instructions
### Prerequisites

Go: Ensure Go is installed on your machine. You can download it from [golang.org](https://go.dev/)

## Clone the Repository
```bash
git clone https://github.com/yourusername/go-gin-template.git
cd go-gin-template
```

## Install Dependencies

Run the following command to install the required dependencies:

```bash

go mod tidy
```

## Install Swagger CLI

Install the Swagger CLI tool for generating documentation:

```bash

go install github.com/swaggo/swag/cmd/swag@latest
```

## Install Testify

Install the testify package for testing:

```bash

go get github.com/stretchr/testify
```

## Install GORM and SQLite

Install GORM and SQLite dependencies if not installed in the go mod tidy:

```bash
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

### Important commands to run in the CLI
#### To generate Swagger Documentation

- To generate and update the Swagger documentation, run the following command:

```bash
swag init -g cmd/server/main.go -d ./ --exclude ./models ./cmd ./internals ./config ./database ./routes ./pkg
```

### To run the application, use the following command:

```bash
go run cmd/server/main.go
```

This starts the server, and you should be able to access it at http://localhost:8080 (or the port specified in your configuration).
Directory Structure

- cmd/server/main.go: Entry point for the application.
- config/config.go: Configuration settings.
- database/db.go: Database initialization and connection.
- models/note.go: Example model.
- pkg/logger/logger.go: Logging setup.
- gorm.db: Sqlite database
- README.md: This file.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
