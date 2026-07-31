# 👟 Nike Store Backend API

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/Gin_Framework-008080?style=for-the-badge&logo=gin&logoColor=white">
  <img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white">
  <img src="https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white">
</p>

<p align="center">
A scalable RESTful API for a modern Nike Store application built with Go, Gin, and PostgreSQL.
</p>

---

## 📖 Overview

Nike Store Backend is a RESTful API built with **Go** for powering modern e-commerce applications.

The project follows a **Layered Architecture** with a clear separation between handlers, services, repositories, and domain models, making the codebase easier to maintain and extend.

It is designed primarily for mobile clients such as **Flutter** and **React Native**, but can be consumed by any REST client.

---

# ✨ Features

- 🔐 JWT Authentication
- 👤 User Registration & Login
- 👟 Product Management
- 💬 Product Comments
- 🌍 Persian & English Comment Support
- 📄 Pagination (LIMIT / OFFSET)
- 🛡️ PostgreSQL Constraint Validation
- 🚀 High Performance REST API
- 📱 Mobile-Friendly JSON Responses
- 🐳 Docker & Docker Compose Support

---

# 🛠 Tech Stack

| Technology | Description |
|------------|-------------|
| Go | Programming Language |
| Gin | HTTP Framework |
| PostgreSQL | Database |
| JWT | Authentication |
| Docker | Containerization |

---

# 🏗 Project Structure

The project is organized using a layered architecture.

```text
.
├── cmd
│   ├── api
│   ├── comments
│   ├── crawler
│   └── seed
│
├── internal
│   ├── data
│   │   └── repository
│   │       ├── queries
│   │       ├── comment_repository.go
│   │       ├── postgres_product_repository.go
│   │       └── user_repository.go
│   │
│   ├── domain
│   │   ├── category.go
│   │   ├── comment.go
│   │   ├── errors.go
│   │   ├── product.go
│   │   └── user.go
│   │
│   ├── handler
│   │   ├── middleware
│   │   ├── auth.go
│   │   ├── signup.go
│   │   ├── get_all_product_handler.go
│   │   ├── get_product_by_id_handler.go
│   │   ├── get_comments_handler.go
│   │   └── ping.go
│   │
│   └── services
│       ├── auth_service.go
│       ├── signup_service.go
│       ├── product_service.go
│       └── comment_service.go
│
├── Dockerfile
├── docker-compose.yml
├── init.sql
├── go.mod
└── README.md
```

---

## 📂 Directory Overview

| Folder | Responsibility |
|---------|----------------|
| `cmd/` | Application entry points (API, Seeder, Crawler) |
| `internal/domain` | Domain models |
| `internal/services` | Business logic |
| `internal/data/repository` | Database access layer |
| `internal/data/repository/queries` | Raw SQL queries |
| `internal/handler` | HTTP handlers |
| `internal/handler/middleware` | Authentication middleware |

---

# 🚀 Getting Started

## Prerequisites

- Go 1.21+
- PostgreSQL
- Docker (Optional)

---

## Clone Repository

```bash
git clone https://github.com/Nike-Store-Workspace/nilke-store-backend.git

cd nilke-store-backend
```

---

## Install Dependencies

```bash
go mod download
```

---

## Configure Environment Variables

Create a `.env` file in the project root.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=nike_store

JWT_SECRET=your_secret_key

SERVER_PORT=8090
```

---

## Run the Application

```bash
go run ./cmd/api
```

or

```bash
go run main.go
```

(depending on your project configuration)

---

## Run with Docker

```bash
docker compose up --build
```

---

# 📚 Current API Modules

- Authentication
- Users
- Products
- Categories
- Comments
- Pagination

---

# 🎯 Project Goals

- Clean and readable code
- High performance
- Easy maintenance
- Separation of concerns
- Mobile-first API
- Easy scalability

---

# 🤝 Contributing

Contributions are welcome!

1. Fork the repository.

2. Create a feature branch.

```bash
git checkout -b feature/AmazingFeature
```

3. Commit your changes.

```bash
git commit -m "Add AmazingFeature"
```

4. Push to your branch.

```bash
git push origin feature/AmazingFeature
```

5. Open a Pull Request.

---

# 📌 Roadmap

- [x] JWT Authentication
- [x] Product APIs
- [x] Comment System
- [x] Pagination
- [ ] Shopping Cart
- [ ] Wishlist
- [ ] Order Management
- [ ] Payment Gateway
- [ ] Admin Dashboard
- [ ] Unit Tests
- [ ] Swagger Documentation

---

# ⭐ Support

If you like this project, please consider giving it a **Star ⭐**.

It helps the project grow and motivates future development.

---

# 📄 License

This project is licensed under the MIT License.

---

<p align="center">
Made with ❤️ using Go & Gin
</p>