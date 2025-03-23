# Kotync

Rewrite of [Kotatsu Synchronization Server](https://github.com/KotatsuApp/kotatsu-syncserver) in Go using the Echo framework.

## Compatibility

- Can work with MySQL database from original server. See [Installing](/docs/install.md).

## Differences

- Uses Blake3 for hashing passwords
- Configurable via environment variables and a configuration file

### API differences

- `get /manga`: max `limit` is 1000

## Why?

- To run on low hardware
- To run as single-binary
- To not use database server (SQLite is used instead by default)

## Installing

See [Installing](/docs/install.md).

## Usage

1. Clone the repository
2. Install Go
3. Run `go build` to build the application
4. Run `./kotync` to start the server

## API Documentation

### Authentication

- `POST /auth`: Register or login a user
  - Request body: `{ "email": "user@example.com", "password": "password" }`
  - Response: `{ "token": "jwt_token" }`

### User Management

- `GET /me`: Retrieve authenticated user's information
  - Headers: `Authorization: Bearer jwt_token`
  - Response: `{ "id": 1, "email": "user@example.com", "nickname": "User" }`

### Manga Operations

- `GET /manga`: Retrieve a list of manga
  - Query parameters: `offset`, `limit`
  - Response: `[ { "id": 1, "title": "Manga Title", ... } ]`

- `GET /manga/:id`: Retrieve a specific manga by its ID
  - Response: `{ "id": 1, "title": "Manga Title", ... }`

- `POST /manga`: Create a new manga
  - Request body: `{ "title": "Manga Title", ... }`
  - Response: `{ "id": 1, "title": "Manga Title", ... }`

- `PUT /manga/:id`: Update an existing manga
  - Request body: `{ "title": "Updated Manga Title", ... }`
  - Response: `{ "id": 1, "title": "Updated Manga Title", ... }`
