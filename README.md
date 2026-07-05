# GoSkeleton
**GoSkeleton** is a simple, lightweight, and production-ready boilerplate for hitting the ground running with a web site or API in Go (Golang). 

This project is designed as an ideal starting point (a "skeleton") that eliminates the routine chore of setting up basic architecture, allowing you to dive straight into writing your business logic.

## 🚀 Features

* **Clean Architecture:** Clear separation of layers (routes, handlers, services).
* **Minimal Dependencies:** Built primarily using the Go standard library and proven tools.
* **Ready-to-use Config:** Environment variable support via `.env`.
* **Logging:** Built-in basic request and error logging.
* **Docker Ready:** Pre-configured `Dockerfile` for rapid deployment.

## 🛠 Tech Stack

* **Language:** Go (Golang)
* **Router:** Clean `net/http` (or optionally `chi` / `gin`)
* **Config:** `godotenv`
* **Containerization:** Docker

## 📋 Project Structure

```text
goskeleton/
├── cmd/
│   └── server/
│       └── main.go       # Application entry point
├── internal/
│   ├── config/           # Configuration loading and management
│   ├── handlers/         # HTTP request handlers (controllers)
│   └── server/           # HTTP server setup and startup
├── .env.example          # Sample environment configuration file
├── Dockerfile            # Instructions for building the Docker image
├── go.mod                # Go module files
└── README.md             # Project documentation
