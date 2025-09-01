### 📚 Student API - A Go-based RESTful Service

This is a simple, yet robust, RESTful API for managing student records, built with Go. It utilizes a SQLite database for data persistence and provides a set of endpoints to perform CRUD (Create, Read, Update, Delete) operations on student data. The project is designed with a clean architecture and includes request validation.

### ✨ Features
- **RESTful Endpoints**: Full CRUD operations for student records.
- **SQLite Database**: Simple and efficient file-based storage.
- **Request Validation**: Ensures data integrity for incoming requests.
- **Structured Logging**: Uses `slog` for clear and informative logging.
- **Graceful Shutdown**: Handles server shutdown elegantly on `SIGINT` or `SIGTERM` signals.

### 🛠️ Prerequisites
- Go 1.25.0 or later.

### 🚀 Getting Started

Clone the repository and navigate to the project directory:

```bash
git clone [https://github.com/ageniouscoder/student-api](https://github.com/ageniouscoder/student-api)
cd studentapi
