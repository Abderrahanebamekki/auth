# Auth Service

A lightweight Go authentication service built on top of **Supabase Auth** REST API. It exposes signup, login, logout, a public endpoint, and a protected profile endpoint guarded by a custom authentication middleware.

## Features

- **Sign up** — create a new user via Supabase Auth
- **Login** — authenticate and receive access/refresh tokens
- **Sign out** — revoke the current session
- **Public route** — no authentication required
- **Protected profile** — requires a valid access token (verified against Supabase)
- **Auth middleware** — validates `Bearer` tokens using a custom `contextKey` to store the authenticated user in the request context
- **Swagger UI** — interactive API documentation

## Tech Stack

- **Go** (standard library `net/http`)
- **Gorilla Mux** — HTTP router
- **Supabase Auth** — user management and token verification
- **Swaggo / HTTP-Swagger** — API documentation

## Prerequisites

- Go 1.26+
- A Supabase project with Auth enabled

## Setup

1. Clone the repository

   ```bash
   git clone https://github.com/Abderrahanebamekki/auth.git
   cd auth
   ```

2. Create a `.env` file from the existing one or add the following variables

   ```env
   SUPABASE_URL=https://<your-project>.supabase.co
   SUPABASE_KEY=<your-anon-or-service-key>
   PORT=8080
   ```

3. Run the server

   ```bash
   go run .
   ```

The server will start on the port defined in `PORT` (default in docs config: `8080`).

## API Endpoints

| Method | Path                   | Auth Required | Description                       |
| ------ | ---------------------- | ------------- | --------------------------------- |
| POST   | `/auth/signup`         | No            | Create a new user account         |
| POST   | `/auth/login`          | No            | Authenticate and get tokens       |
| POST   | `/auth/signout`        | Yes (Bearer)  | Log out the current user          |
| GET    | `/public/info`         | No            | Public welcome message            |
| GET    | `/protected/profile`   | Yes (Bearer)  | Get the authenticated user's data |

### Example Requests

**Sign up**

```bash
curl -X POST http://localhost:8080/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

**Login**

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

**Protected profile**

```bash
curl -H "Authorization: Bearer <ACCESS_TOKEN>" \
  http://localhost:8080/protected/profile
```

**Sign out**

```bash
curl -X POST -H "Authorization: Bearer <ACCESS_TOKEN>" \
  http://localhost:8080/auth/signout
```

## Project Structure

```
.
├── docs/            # Generated Swagger documentation
├── handler/         # HTTP handlers (signup, login, signout, public, profile)
├── middleware/      # AuthGuard middleware
├── model/           # Data models (User)
├── router/          # Route registration
├── main.go          # Server entry point + Swagger general config
└── .env             # Environment variables (not committed)
```

## Swagger UI

Interactive API documentation is available at:

```
http://localhost:8080/swagger/index.html
```

The **profile** endpoint is protected and requires a `Bearer` token — open the "Authorize" dialog in Swagger UI and paste your access token.

To regenerate the docs after changing annotations:

```bash
swag init
```

## License

This project is for learning purposes.
