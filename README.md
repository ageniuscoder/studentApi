# 📚 Student API - A Go-based RESTful Service

This project is a straightforward and robust RESTful API for managing student data. Built with Go, it offers a full suite of **CRUD** (**C**reate, **R**ead, **U**pdate, **D**elete) operations and uses a file-based **SQLite** database for persistence. The application is designed with a clean architecture, featuring clear **separation of concerns**, structured logging, and graceful server management.

---

### ✨ Features

* **RESTful Endpoints**: Full CRUD operations for student records.
* **SQLite Database**: Uses SQLite for simple and efficient data storage.
* **Request Validation**: Ensures data integrity by validating incoming JSON requests.
* **Structured Logging**: Employs the `slog` package for clear and informative logging.
* **Graceful Shutdown**: Shuts down gracefully on interrupt signals (`SIGINT`, `SIGTERM`) to ensure no data is lost during termination.
* **Configuration Management**: Easy-to-use YAML configuration for environment-specific settings.

---

### 🛠️ Prerequisites

* Go 1.25.0 or later

---

### 🚀 Getting Started

To get the API up and running, follow these simple steps.

1.  **Clone the repository**:

    ```bash
    git clone [https://github.com/ageniouscoder/student-api](https://github.com/ageniouscoder/student-api)
    cd student-api
    ```

2.  **Install dependencies**:
    The project uses Go modules. The required dependencies will be downloaded automatically when you run the `go mod tidy` command.

    ```bash
    go mod tidy
    ```

3.  **Run the application**:
    Start the server by providing the path to your configuration file.

    ```bash
    go run ./cmd/students-api/main.go --config=config/local.yaml
    ```
    The server will start on `localhost:8082`.

---

### 📝 API Endpoints

The API interacts with a simple `Student` data model.

#### Student Model:

```go
type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}

POST /api/students
Creates a new student record.

Rqquest Body
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 25
}
Success Response(201 Created)
{
  "id": 123
}
Error Response (400 Bad Request):
{
  "status": "Error",
  "error": "name is required filed"
}
