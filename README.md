# Go REST API (Learning & Practice Project)

A robust, modular REST API built with Go, utilizing the **Gin Web Framework**, **GORM** for MySQL interactions, **Redis** for caching and API rate limiting, **JWT Authentication**, and **Bcrypt** for secure password hashing. This project follows clean architecture principles by separating concerns into Handlers, Services, and Repositories, and is fully containerized using Docker Compose.

---

## 🚀 Tech Stack

- **Language:** [Go (Golang)](https://go.dev/)
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **Database & ORM:** [MySQL 8.0](https://www.mysql.com/) & [GORM](https://gorm.io/)
- **Caching & Rate Limiting:** [Redis 7](https://redis.io/) (`go-redis/v9`)
- **Containerization:** [Docker & Docker Compose](https://www.docker.com/)
- **Authentication:** [JWT (JSON Web Tokens)](https://github.com/golang-jwt/jwt/v5)
- **Password Hashing:** [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- **Environment Management:** [Godotenv](https://github.com/joho/godotenv)

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
│   └── services/         # Business logic & caching workflows
├── pkg/
│   ├── hashing/          # Password hashing utilities (bcrypt)
│   ├── jwt/              # JWT token generation & validation
│   ├── redis_cache/      # Redis client connection setup
│   └── response/         # Standardized API response format helper
├── .env                  # Environment variables
├── Dockerfile            # Multi-stage build for Go backend
├── docker-compose.yml    # Multi-container orchestration (App, MySQL, Redis)
├── go.mod
└── go.sum
```


## 🐳 Running with Docker Compose

To spin up the entire stack (Go backend, MySQL database, and Redis cache) simultaneously:

1. Create your `.env` file in the root directory based on your required environment variables.
2. Start the containers using Docker Compose:
   ```bash
   docker compose up --build -d
   ```
3. To stop all service using Docker Compose:
    ```bash
    docker compose down
    ```