# 🧑‍🧒‍🧒 User Service

This is a Golang-based API for managing users, supporting operations such as user creation, retrieval, and health checks.

The API is defined using OpenAPI 3.0 specifications, making it easy to understand, extend, and integrate with.

## 📑 Table of Contents

- [📋 Overview](#-overview)
- [▶️ Getting Started](#-getting-started)
- [🗂️ Project Structure](#-project-structure)
- [🧪 Testing](#-testing)

## 📋 Overview

- **API Version**: 1.0.0
- **Description**: The User Service allows for managing users in a structured, RESTful manner.
- **OpenAPI Specification**: `api/openapi/user-service.yaml`
- **Language**: Go

## ▶️ Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/maslechkin86/go-api-example.git
   cd user-service
   ```
   
2. Install dependencies:
   ```bash
   go mod download
   ```

3. Download tools:
   ```bash
   make tools-install
   ```
   
4. Generate API code from the OpenAPI spec:
   ```bash
   make gen
   ```

5. Run the service:
   ```bash
   make go-run
   ```

## 🗂️ Project Structure
    ```text
    user-service/
    ├── api/                 # API specifications (e.g., OpenAPI)
    ├── cmd/                 # Main application entry points
    ├── internal/            # Private application code (business logic, handlers)
    ├── pkg/                 # Public packages
    └── Makefile             # Build, test, and generate commands
    ```
## 🧪 Testing
Run unit and integration tests using:
   ```bash
   make test
   ```

