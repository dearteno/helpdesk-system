# Helpdesk Management System

A modern, microservice-based Helpdesk Management System built with Go 1.24, gRPC, Supabase, and SvelteKit.

## 📖 Overview

This Helpdesk Management System is designed according to industry-standard UML diagrams to provide a robust solution for managing support tickets, issues, tracking, and helpdesk operations. The system is built as a monorepo containing:

- **API Gateway**: Go-based gateway for routing requests to microservices
- **Microservices**: Six separate services for different business domains
- **Frontend**: Modern SvelteKit UI with TypeScript and TailwindCSS
- **Database**: Supabase for database operations and authentication
- **API Communication**: gRPC for efficient service-to-service communication

## 🏗️ Architecture

The system follows a microservice architecture pattern with:

```
                    ┌─────────────────┐
                    │    Frontend     │
                    │  (SvelteKit)    │
                    └────────┬────────┘
                             │
                             ▼
┌────────────────────────────────────────────────┐
│                API Gateway                     │
│          (Go 1.24 with Gin & gRPC)            │
└───┬─────────┬────────┬─────────┬─────┬────────┘
    │         │        │         │     │
    ▼         ▼        ▼         ▼     ▼        ▼
┌─────────┐ ┌──────┐ ┌──────┐ ┌─────┐ ┌───┐ ┌────────┐
│ Helpdesk│ │Ticket│ │Issues│ │Track│ │FAQ│ │Network │
│ Service │ │Service│ │Service│ │Service│ │Service│ │Service │
└────┬────┘ └───┬───┘ └───┬───┘ └──┬──┘ └─┬─┘ └────┬───┘
     │          │         │        │      │        │
     └──────────┴─────────┴────────┴──────┴────────┘
                          │
                          ▼
                    ┌─────────────┐
                    │  Supabase   │
                    │ (Database & │
                    │    Auth)    │
                    └─────────────┘
```

### Microservices

Based on the UML class diagram, the system is divided into the following microservices:

1. **Helpdesk Service**: Manages helpdesk entities and operations
2. **Ticket Service**: Handles ticket creation, tracking, and resolution
3. **Issues Service**: Manages issue reporting and resolution tracking
4. **Track Service**: Provides tracking functionality for issues and tickets
5. **Network Service**: Handles network-related operations
6. **FAQ Service**: Manages knowledge base and frequently asked questions

## 🚀 Technology Stack

### Backend
- **Language**: Go 1.24
- **API Gateway**: Gin HTTP framework
- **Service Communication**: gRPC
- **Database**: Supabase (PostgreSQL)
- **Authentication**: Supabase Auth

### Frontend
- **Framework**: SvelteKit 2.x with SSR
- **Styling**: TailwindCSS with Typography and Forms plugins
- **State Management**: Built-in Svelte stores
- **API Communication**: Fetch API with REST endpoints
- **Authentication**: Supabase Auth
- **Component Development**: Storybook
- **Testing**: Vitest, Playwright

## 📂 Project Structure

```
helpdesk-system/
├── README.md
├── api-gateway/             # API Gateway service
│   ├── cmd/                 # Command-line entry points
│   │   └── server/          # API Gateway server
│   └── pkg/                 # Reusable packages
│       ├── config/          # Configuration management
│       └── middleware/      # HTTP middleware components
├── docs/                    # Documentation
├── proto/                   # Shared Protocol Buffer definitions
│   └── common.proto         # Common message definitions
├── scripts/                 # Build and deployment scripts
├── services/                # Microservices
│   ├── faq/                 # FAQ service
│   ├── helpdesk/            # Helpdesk service
│   ├── issues/              # Issues service
│   ├── network/             # Network service
│   ├── ticket/              # Ticket service
│   └── track/               # Track service
└── ui/                      # Frontend SvelteKit application
    ├── src/                 # SvelteKit source code
    │   ├── lib/             # Shared components and utilities
    │   ├── routes/          # SvelteKit routes and pages
    │   └── stories/         # Storybook components
    └── static/              # Static assets
```

## 🔧 Setup and Installation

### Prerequisites

- Go 1.24 or higher
- Node.js 18+ and pnpm
- Docker (optional, for local development)
- Supabase account or self-hosted Supabase instance

### Environment Setup

1. **Clone the repository**

```bash
git clone <repository-url>
cd helpdesk-system
```

2. **Set up Supabase**

- Create a new Supabase project
- Set up the required tables (Helpdesk, Ticket, Issues, Track, Network, FAQ)
- Copy the Supabase URL and API Key for configuration

3. **Configure environment variables**

Create a `.env` file in the root directory with:

```
SUPABASE_URL=your_supabase_url
SUPABASE_KEY=your_supabase_key
```

And a `.env` file in the `ui` directory with:

```
VITE_SUPABASE_URL=your_supabase_url
VITE_SUPABASE_ANON_KEY=your_supabase_anon_key
```

### Starting the Services

1. **Start the API Gateway**

```bash
cd api-gateway
go mod tidy
go run cmd/server/main.go
```

2. **Start each microservice**

```bash
# Example for ticket service
cd services/ticket
go mod tidy
go run cmd/server/main.go
```

3. **Start the UI development server**

```bash
cd ui
pnpm install
pnpm run dev
```

## 📱 Features

Based on the UML diagrams, the system provides:

### User Roles (from Use Case Diagram)

- **Super Admin**: Complete system management
- **System User**: Manage helpdesk, tickets, issues, track, network, FAQ
- **Client**: View tickets, add comments, change status, resolve tickets
- **User**: Create tickets, view status, add comments

### Core Features

- Ticket Management (create, update, view, search)
- Helpdesk Operations
- Issue Tracking and Resolution
- FAQ Management
- User Authentication and Authorization

## 🧪 Testing

### Backend Testing

```bash
cd services/ticket
go test ./...
```

### Frontend Testing

```bash
cd ui
pnpm test
```

## 📝 Development Guidelines

### Adding a New Microservice

1. Create a new directory under `services/`
2. Define Protocol Buffer messages and services in the `proto/` directory
3. Generate gRPC code with `protoc`
4. Implement the service handler in the `pkg/handlers` directory
5. Create a repository implementation in `pkg/repository`
6. Update the API Gateway to route requests to the new service

### Extending the Frontend

1. Create new routes in the `ui/src/routes` directory
2. Add API client functions in the `ui/src/lib` directory
3. Update the navigation in the `ui/src/routes/+layout.svelte` file

## 📄 License

[MIT](LICENSE)

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
