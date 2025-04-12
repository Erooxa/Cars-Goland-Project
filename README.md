# Cars API

A REST API for managing cars, built with Go, Gin, and PostgreSQL.

## Docker Setup

This project includes Docker configuration for easy deployment.

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Running with Docker Compose

1. Clone the repository:
   ```
   git clone <repository-url>
   cd Cars
   ```

2. Build and start the containers:
   ```
   docker-compose up -d
   ```

   This will:
   - Build the Go application
   - Start the PostgreSQL database
   - Start the application server

3. The API will be available at `http://localhost:8080`

4. To stop the containers:
   ```
   docker-compose down
   ```

### Database Management

#### Rolling Back Migrations

If you need to roll back the latest migration:

```
docker-compose -f docker-compose.rollback.yml up rollback
```

This command will:
1. Connect to the same database used by the main application
2. Execute the rollback operation
3. Exit when complete

#### Database Persistence

The PostgreSQL data is persisted in a Docker volume named `postgres-data`. This ensures your data remains even after containers are stopped or removed.

## API Endpoints

### Authentication

- **POST /signup** - Register a new user
  ```json
  {
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }
  ```

- **POST /login** - Login and get access token
  ```json
  {
    "email": "test@example.com",
    "password": "password123"
  }
  ```

### Car Management

The following endpoints require authentication with the token received from login:

- **GET /api/cars** - Get all cars
- **GET /api/cars/:id** - Get a specific car by ID
- **POST /api/admin/cars** - Create a new car (admin only)
- **PUT /api/admin/cars/:id** - Update a car (admin only)
- **DELETE /api/admin/cars/:id** - Delete a car (admin only)

## Default Users

The system creates two default users during initialization:

1. **Admin User**
   - Email: admin@example.com
   - Password: admin123
   - Role: admin

2. **Regular User**
   - Email: user@example.com
   - Password: user123
   - Role: user 