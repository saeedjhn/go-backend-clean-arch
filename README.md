# Go Backend Clean Architecture

A Go (Golang) Backend Clean Architecture Project Based
on <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html"> **Uncle Bob's Clean
Architecture** </a> And <a href="https://github.com/golang-standards/project-layout">
**Standard Go Project
Layout** </a> with Echo, MysqlDB, JWT Authentication Middleware, Test, and Docker.

<section style="text-align: center;margin: auto">

![Uncle-Bobs-Clean-Arch](./assets/CleanArchitecture.jpg)
</section>

---

## Motivation

In the Golang ecosystem, there is no strict structure for backend projects.  
This sometimes causes confusion for developers, especially **junior developers**.

This project can be used as a **template** to build your own backend project in Go.

Before creating this project, I have gone through more than **XX** projects related to Go (Golang) Clean Architecture on
GitHub.

Thanks to all those projects, I learned a lot from them.  
As I keep saying:

> The best way to learn to code is to code.  
> But, to write good code, you also need to read good code.  
> Make a habit of reading good code.  
> You can find many open-source projects on GitHub and start reading.

For the implementation part, I combined all of my ideas, experiences, and learnings from those projects to create this
project.

And as always, I would love to get **feedback** on my project.  
This helps everyone and most importantly **me**.

---

## Architecture Layers

The architecture of this project follows **Clean Architecture** principles:

- **Router**
- **Controller / Handler**
- **Service / Use-case**
- **Repository**
- **Domain / Entities**

![go-backend-clean-arch](./assets/diagram.architecture.png)

---

## Project Development Guide

This project provides a `Makefile` to simplify development, database migrations, testing, and CI/CD.  
Below is a guide for the most frequently used commands.

### Development & Build

- **Start development environment**
  ```bash
  make development-up
  ```
  Starts and builds all services using Docker Compose.  
  It also enables live reload via air.

- **Build services**
  ```bash
  make build
  ```
  Builds all services via Docker Compose without starting them.

- **Stop and clean up**
  ```bash
  make down
  ```
  Stops and removes containers, networks, images, and volumes.

- **Cross-platform builds**
  ```bash
  make build/linux
  make build/darwin
  make build/windows
  ```
  Compile Go binaries for Linux, macOS (Darwin), or Windows.

---

### Database Migrations

- **Run migrations**
  ```bash
  make migrate-up
  ```
  Applies all available migrations to your database.

- **Rollback migrations**
  ```bash
  make migrate-down
  ```
  Rolls back previously applied migrations.

- **Force migration version**
  ```bash
  make migrate-force
  ```
  Manually sets the migration version.

- **Drop database**
  ```bash
  make migrate-drop
  ```
  Drops the entire database.

---

### Testing

- **Run tests with tparse**
  ```bash
  make tparse
  ```
  Runs tests, generates a JSON test report, and summarizes the output with tparse.

---

## Running the Project

Before running the project, make sure the following environment variables are set in your Makefile:

```makefile
MYSQL_USER ?= admin
MYSQL_PASSWORD ?= password123
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= go-backend-clean-arch_db
MYSQL_DSN := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)"
MYSQL_MIGRATION_PATH := "./internal/repository/mysql/migrations"
```

### Step 1: Run migrations

```bash
make migrate-up
```

This applies all database migrations.

If you change the database, go to the following path and update your Docker database service configuration:

```bash
deployments/development/docker-compose
```

By default, this repository uses MySQL.  
However, Docker configuration for PostgreSQL is also available. You only need to update the values in your Makefile.

### Step 2: Start development environment

```bash
make development-up
```

This command will start all Docker services and enable live reload for development.

---

## CI/CD

This project uses **GitHub Actions** for CI/CD.

- **setup** → prepares the environment (Go, Protobuf, caching modules)
- **audit** → code checks (security, formatting, linting, static analysis)
- **test** → runs tests with `tparse`
- **deploy** → builds and pushes Docker images to DockerHub
    - You must configure the following GitHub Secrets:
        - `DOCKERHUB_USERNAME` → your DockerHub username
        - `DOCKERHUB_TOKEN` → your DockerHub access token
        - `DOCKERHUB_REPOSITORY` → your DockerHub repository name (e.g., `yourusername/yourproject`)

---

## API documentation of Go Backend Clean Architecture

<a href="https://documenter.getpostman.com/view/10523269/2sA3QqgD5c" target="_blank" style="text-decoration: none;">
    <button style="display: inline-block;background-color: lightgreen;color: black;padding: 1.4rem;border: none">
        View API DOCS
    </button>
</a>

--- 

## Feedback

If you have any suggestions, improvements, or feedback, feel free to open an issue or PR.  
Let's learn and grow together 🚀  
