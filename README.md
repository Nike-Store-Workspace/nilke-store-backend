# 👟 Nike Store Backend API

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Gin Framework](https://img.shields.io/badge/Gin-008080?style=for-the-badge&logo=gin&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white)

The **Nike Store Backend** is a high-performance RESTful API built with **Go** using **Clean Architecture** principles. It provides essential services for a modern e-commerce mobile application, including user authentication, product management, dual-language comment systems, pagination, and robust database constraints.

---

## 🏗️ Architecture

This project follows **Clean Architecture / Layered Architecture** to ensure scalability, maintainability, and clear separation of concerns:

```text
├── internal/
│   ├── domain/           # Core models, interfaces, and DTOs
│   ├── repository/       # Database access layer (PostgreSQL)
│   ├── service/          # Business logic layer
│   └── delivery/
│       └── http/         # Gin handlers and HTTP routes
└── main.go               # Application entry point & setup
✨ Features
🔐 JWT Authentication: Secure sign-up and sign-in flow with custom JWT middleware.

💬 Localization & Comments System:

Multi-language support (FA / EN).

Strict Foreign Key integrity and clear PostgreSQL error handling.

⚡ Pagination: Optimized data fetching using dynamic LIMIT and OFFSET.

📱 Mobile-First Design: Built specifically to serve Flutter and React Native clients seamlessly.

🚀 Quick Start
Prerequisites
Go (version 1.21 or higher)

PostgreSQL instance running locally or hosted

Setup Instructions
Clone the repository:

Bash
git clone [https://github.com/Nike-Store-Workspace/nilke-store-backend.git](https://github.com/Nike-Store-Workspace/nilke-store-backend.git)
cd nilke-store-backend
Download dependencies:

Bash
go mod download
Configure Environment Variables (.env):
Create a .env file in the root directory:

Code snippet
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=nike_store
JWT_SECRET=your_jwt_secret_key
Run the server:

Bash
go run main.go
🤝 Contributing
We love open-source collaboration! Contributions, ideas, and feature enhancements are warmly welcomed.

If you'd like to help improve the architecture, add new endpoints (such as cart management or payment gateways), or optimize SQL performance:

Fork the repository.

Create your feature branch (git checkout -b feature/AmazingFeature).

Commit your changes (git commit -m 'Add some AmazingFeature').

Push to the branch (git push origin feature/AmazingFeature).Here is a clean, professional, and visually engaging README.md content in English text format. You can copy and paste this directly into your README.md file in the root of your repository:

👟 Nike Store Backend API
The Nike Store Backend is a high-performance RESTful API built with Go following the principles of Clean Architecture. It provides a robust core backend for modern e-commerce mobile applications (such as Flutter or React Native), handling user authentication, product management, multi-language comments, and efficient data pagination.

🏗️ Architecture Overview
This application adheres to Clean Architecture to maintain clear separation of concerns, testability, and scalability across all layers:

domain/: Enterprise business models, data structures, and repository/service interface contracts.

internal/repository/: Database persistence layer utilizing raw SQL queries and PostgreSQL drivers.

internal/service/: Business logic layer managing data validation, pagination rules, and multi-language fallbacks.

internal/delivery/http/: Transport layer containing Gin route handlers, query parsing, and custom error responses.

✨ Key Features
🔐 Secure Authentication: User sign-up, login, and protected routes using JWT tokens.

💬 Localization & Comments: Product comment system with support for Persian and English content (fa / en).

⚡ Optimized Pagination: Efficient database fetching using offset-based pagination (LIMIT & OFFSET).

🛡️ Robust Error Handling: Clean foreign key constraint validations and domain-specific error feedback.

📱 Mobile-Ready: Formatted responses designed specifically for seamless mobile client integration.

🚀 Quick Start
Prerequisites
Go (version 1.21 or later)

PostgreSQL database instance

Installation & Run
Clone the repository:

Bash
git clone [https://github.com/Nike-Store-Workspace/nilke-store-backend.git](https://github.com/Nike-Store-Workspace/nilke-store-backend.git)
cd nilke-store-backend
Download dependencies:

Bash
go mod download
Configure environment variables:
Create a .env file in the root directory and specify your configuration:

Code snippet
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=nike_store
JWT_SECRET=your_jwt_secret_key
SERVER_PORT=8090
Start the server:

Bash
go run main.go
🤝 Open for Contributions!
We love open-source collaboration! Whether you want to refactor code, improve overall architecture, write tests, or add major features like cart management, checkout, and payment gateways — your contributions are warmly welcome!

How to Contribute:
Fork the repository.

Create your feature branch (git checkout -b feature/AmazingFeature).

Commit your changes (git commit -m 'Add some AmazingFeature').

Normally I can help with things like this, but I don't seem to have access to that content. You can try again or ask me for something else.