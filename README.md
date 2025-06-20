# iOS App Store Reviews Viewer

A full-stack application that monitors and displays iOS App Store reviews. Built with Go and React, it automatically fetches reviews every 5 minutes and shows the latest 48 hours of feedback.

## Features

- Track multiple iOS apps simultaneously
- Automatic review syncing every 5 minutes
- Shows reviews from the last 48 hours
- Clean React interface with pagination
- PostgreSQL database for persistence
- RESTful API for all operations
- Responsive design for desktop, tablet, and mobile
- Search and filtering capabilities
- Real-time data fetching with caching

## Architecture

```
React Frontend ←→ Go API Server ←→ Cron Service
                      ↓
              PostgreSQL Database
```

## Tech Stack

**Backend:**

- Go 1.24.4+ with Gorilla Mux
- PostgreSQL with SQLx
- Background jobs with GoCron

**Frontend:**

- React 18 with TypeScript
- TanStack Query for data fetching and caching
- Tailwind CSS for styling
- Vite for building
- Radix UI for accessible components
- Shadcn/ui component library

## Quick Start

### Prerequisites

- Go 1.24.4+
- Node.js 18+
- PostgreSQL 15+
- Docker (optional)

### Setup

1. **Clone and setup backend:**

```bash
cd api
cp .env-example .env
# Edit .env with your database credentials
```

2. **Start database:**

```bash
docker-compose up -d postgres
```

3. **Run migrations:**

```bash
make migration/up
```

4. **Start the API server:**

```bash
make api
# or: go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

5. **Start the background sync service:**

```bash
# In a new terminal
cd api
make cron
# or: go run cmd/cron/main.go
```

6. **Setup frontend:**

```bash
cd web
cp .env-example .env
# Edit .env with your API URL
```

7. **Install dependencies and start frontend:**

```bash
yarn
yarn dev
```

Visit `http://localhost:5173` to see the app.

## Project Structure

```
├── api/                    # Go backend
│   ├── cmd/               # Entry points (api + cron)
│   │   ├── api/           # Main API server
│   │   └── cron/          # Background sync service
│   ├── internal/          # Application code
│   │   ├── client/        # External API clients
│   │   ├── controller/    # HTTP handlers
│   │   ├── database/      # Database connection
│   │   ├── helpers/       # Utility functions
│   │   ├── repository/    # Data access layer
│   │   └── service/       # Business logic
│   ├── migrations/        # Database schema changes
│   ├── docker-compose.yaml # Local development
│   ├── go.mod             # Dependencies
│   ├── Makefile          # Build commands
│   └── sqlmigrate.yml    # Migration config
├── web/                   # React frontend
│   ├── src/               # Source code
│   │   ├── api/           # API client functions
│   │   ├── components/    # UI components
│   │   ├── pages/         # Page components
│   │   ├── app.tsx        # Main app component
│   │   ├── main.tsx       # Entry point
│   │   ├── query-client.tsx # TanStack Query config
│   │   └── types.ts       # TypeScript types
│   ├── lib/               # UI component library
│   │   ├── components/    # Shadcn/ui components
│   │   └── utils.ts       # Utility functions
│   ├── public/            # Static assets
│   ├── index.html         # HTML template
│   ├── package.json       # Dependencies
│   ├── tsconfig.json      # TypeScript config
│   ├── vite.config.ts     # Vite config
│   └── tailwind.config.js # Tailwind config
```

## API Endpoints

- `GET /api/v1/apps` - List all monitored apps
- `POST /api/v1/apps` - Add a new app to monitor
- `GET /api/v1/apps/{appId}` - Get app details
- `DELETE /api/v1/apps/{appId}` - Remove app from monitoring
- `GET /api/v1/apps/{appId}/reviews` - Get paginated reviews (last 48 hours)
- `POST /api/v1/apps/{appId}/sync` - Manually sync reviews
- `GET /api/v1/apps/{appId}/lookup` - Lookup app info from App Store

## Database Schema

### Monitored Apps Table

```sql
CREATE TABLE monitored_apps (
  app_id text NOT NULL PRIMARY KEY,
  app_name text NOT NULL,
  logo_url text NOT NULL,
  nickname text,
  last_synced_at timestamp,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
```

### App Reviews Table

```sql
CREATE TABLE app_reviews (
  id text NOT NULL PRIMARY KEY,
  app_id text NOT NULL,
  title text NOT NULL,
  content text NOT NULL,
  author text NOT NULL,
  rating int NOT NULL,
  submitted_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
```

## Development

### Backend Development

#### Running Tests

```bash
cd api
go test ./...
go test -cover ./...
```
