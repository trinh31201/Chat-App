# Chat Application with Kratos and Go

A real-time chat application built with Go, Kratos framework, gRPC, and WebSocket support.

## Architecture

- **Backend**: Go with Kratos framework
- **API**: gRPC + HTTP/REST (via gRPC-Gateway)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: JWT tokens

## Features

- User registration and authentication
- Create and join chat rooms
- Send and receive messages in real-time
- Message history with pagination
- Online/offline status tracking
- Read receipts
- Redis caching for performance

## Project Structure

```
chat-app/
├── api/              # Protocol buffer definitions
│   ├── chat/v1/      # Chat and room services
│   └── user/v1/      # User service
├── cmd/chat/         # Application entry point
├── configs/          # Configuration files
├── internal/
│   ├── conf/         # Configuration structures
│   ├── data/         # Data access layer (repositories)
│   ├── server/       # HTTP and gRPC servers
│   └── service/      # Business logic
└── third_party/      # Third-party proto files
```

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- Protocol Buffers compiler (protoc)

## Quick Start

### Using Docker Compose (Recommended)

1. Clone the repository
2. Start all services:
```bash
docker-compose up -d
```

This will start:
- PostgreSQL database (port 5432)
- Redis cache (port 6379)
- Chat application (HTTP: 8000, gRPC: 9000)

### Manual Setup

1. Install dependencies:
```bash
go mod download
```

2. Install protoc plugins:
```bash
make init
```

3. Generate code from proto files:
```bash
make api
```

4. Set up PostgreSQL:
```bash
createdb chatdb
psql chatdb < internal/data/schema.sql
```

5. Update `configs/config.yaml` with your database credentials

6. Generate Wire dependencies:
```bash
go install github.com/google/wire/cmd/wire@latest
wire ./cmd/chat
```

7. Run the application:
```bash
go run ./cmd/chat -conf ./configs
```

## API Endpoints

### REST API

#### User Service
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - User login
- `GET /api/v1/users/{id}` - Get user profile
- `PUT /api/v1/users/{user_id}/status` - Update user status

#### Room Service
- `POST /api/v1/rooms` - Create room
- `GET /api/v1/rooms/{id}` - Get room details
- `GET /api/v1/users/{user_id}/rooms` - List user's rooms
- `POST /api/v1/rooms/{room_id}/join` - Join room
- `POST /api/v1/rooms/{room_id}/leave` - Leave room

#### Chat Service
- `POST /api/v1/messages` - Send message
- `GET /api/v1/rooms/{room_id}/messages` - Get messages
- `POST /api/v1/messages/{message_id}/read` - Mark as read

### gRPC API

Connect to `localhost:9000` for gRPC services. See proto files in `api/` for service definitions.

## Testing

### Using gRPCurl

```bash
# List services
grpcurl -plaintext localhost:9000 list

# Register user
grpcurl -plaintext -d '{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}' localhost:9000 api.user.v1.UserService/Register

# Login
grpcurl -plaintext -d '{
  "email": "test@example.com",
  "password": "password123"
}' localhost:9000 api.user.v1.UserService/Login
```

### Using cURL (REST)

```bash
# Register user
curl -X POST http://localhost:8000/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login
curl -X POST http://localhost:8000/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

## Development

### Generate Proto Code
```bash
make api
```

### Build Application
```bash
make build
```

### Run Tests
```bash
make test
```

## Configuration

Edit `configs/config.yaml` to configure:
- Server ports
- Database connection
- Redis connection
- JWT settings
- Logging

## Docker Build

Build image:
```bash
docker build -t chat-app .
```

Run container:
```bash
docker run -p 8000:8000 -p 9000:9000 chat-app
```

## Next Steps

- [ ] Add WebSocket support for real-time messaging
- [ ] Implement message encryption
- [ ] Add file upload support
- [ ] Create web frontend
- [ ] Add push notifications
- [ ] Implement message search
- [ ] Add user presence tracking
- [ ] Create mobile apps

## License

MIT