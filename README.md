## Project Overview
This is a simple microservice that tracks CI/CD build This project is a simple microservice that tracks CI/CD build results.
Each build record contains the Git commit ID, Docker image tag, and build status.
The service provides RESTful API endpoints to create, retrieve, update, and delete build records, and stores the data in a MySQL database.


## Tech Stack

- Language: Golang
- Framework: Gin Web Framework
- ORM: GORM
- Database: MySQL 8.0.43
- Environment: Docker & Docker Compose
- API Documentation: Swagger (using swaggo/swag)
<br>

## Getting Started

1. Create a `.env` file in the root directory with the following content:

```env
PORT=8080

DB_USER=dreamus
DB_PASSWORD=dreamus123
DB_HOST=mysql
DB_PORT=3306
DB_NAME=dreamus_db

MYSQL_ROOT_PASSWORD=dreamusAdmin123
MYSQL_DATABASE=dreamus_db
MYSQL_USER=dreamus
MYSQL_PASSWORD=dreamus123
```
<br>

## Build & Run

2. Make sure Docker and Docker Compose are installed and running.
3. Start the full development environment (MySQL + application) with:

```bash
docker-compose up --build
```

This will:
- Build the application Docker image
- Start a MySQL 8.0.43 database container
- Launch the Gin web server on port `8080`

You can access the Swagger documentation at:
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## API Endpoints (Examples)

### 1. Create a new build (POST `/builds`)
```bash
curl -i -X POST http://localhost:8080/builds \
-H "Content-Type: application/json" \
-d '{
    "commit_id": "a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2",
    "docker_tag": "my-app:1.0.0",
    "status": "PENDING"
}'
```

### 2. Get a build records by ID (GET `/builds/:id`)
```bash
curl -i http://localhost:8080/builds/1
```

### 3. Update a build (PUT `/builds/:id`)
```bash
curl -i -X PUT http://localhost:8080/builds/1 \
-H "Content-Type: application/json" \
-d '{
    "status": "SUCCESS"
}'
```

### 4. Delete a build (DELETE `/builds/:id`)
```bash
curl -i -X DELETE http://localhost:8080/builds/1
```
<br>

## Notes
- ChatGPT was used for:
  - Writing this README
  - Generating Swagger annotations in controller files
  - Providing help with unfamiliar parts of the task
- Docker Compose reference:
  https://github.com/rickseven/dockerized-golang-mysql-api-sample/blob/master/docker-compose.yml