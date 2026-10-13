# Go REST API (Learning & Practice Project)

A robust, modular REST API built with Go, utilizing the **Gin Web Framework**, **GORM** for MySQL interactions, **JWT Authentication**, and **Bcrypt** for secure password hashing. This project follows clean architecture principles by separating concerns into Handlers, Services, and Repositories.

---

## 🚀 Tech Stack

* **Language:** [Go (Golang)](https://go.dev/)
* **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
* **ORM & Database:** [GORM](https://gorm.io/) & [MySQL](https://www.mysql.com/)
* **Authentication:** [JWT (JSON Web Tokens)](https://github.com/golang-jwt/jwt/v5)
* **Password Hashing:** [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
* **Environment Management:** [Godotenv](https://github.com/joho/godotenv)

---

## 📂 Project Architecture

The project follows a clean, layered directory structure:

```text
go-rest-api/
├── cmd/
│   └── server/
│       └── main.go       # Application entry point
├── internal/
│   ├── handlers/         # HTTP request parsers & response senders (Gin)
│   ├── models/           # Database entities / structs (GORM)
│   ├── repositories/     # Database queries & data access layer
│   └── services/         # Business logic & authentication workflows
├── pkg/
│   ├── hashing/          # Password hashing utilities (bcrypt)
│   ├── jwt/              # JWT token generation & validation
│   └── response/         # Standardized API response format helper
├── .env                  # Environment variables
├── go.mod
└── go.sum