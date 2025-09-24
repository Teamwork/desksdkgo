# Teamwork Desk Go SDK

A Go SDK for interacting with the Teamwork Desk API. This SDK provides a simple and intuitive way to interact with Teamwork Desk's REST API.

## Features

- Support for all major Teamwork Desk API endpoints
- Simple and intuitive client interface
- Built-in logging and debugging support
- Command-line interface for quick operations
- Environment variable support for configuration
- `.env` file support for easy configuration

## Installation

```bash
go get github.com/teamwork/desksdkgo
```

## Usage

### Basic Client Usage

```go
package main

import (
    "context"
    "log/slog"
    "github.com/teamwork/desksdkgo/client"
)

func main() {
    // Create a new client
    c := client.NewClient(
        "https://yourcompany.teamwork.com/desk/api/v2",
        client.WithAPIKey("your-api-key"), // Get this in your profile settings
        client.WithLogLevel(slog.LevelDebug),
    )

    // Use the client
    ctx := context.Background()

    // List tickets
    tickets, err := c.Tickets.List(ctx, nil)
    if err != nil {
        panic(err)
    }

    // Get a specific ticket
    ticket, err := c.Tickets.Get(ctx, 123)
    if err != nil {
        panic(err)
    }
}
```

### Using a Custom Logger

You can provide your own log/slog logger to the client:

```go
package main

import (
    "context"
    "log/slog"
    "os"
    "github.com/teamwork/desksdkgo/client"
)

func main() {
    // Create a custom logger
    customLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
        AddSource: true,
    }))

    // Create client with custom logger
    c := client.NewClient(
        "https://yourcompany.teamwork.com/desk/api/v2",
        client.WithAPIKey("your-api-key"),
        client.WithLogger(customLogger), // Use your custom logger
    )

    // The client will now use your custom logger for all HTTP request/response logging
    ctx := context.Background()
    tickets, err := c.Tickets.List(ctx, nil)
    if err != nil {
        panic(err)
    }
}
```

### Available Resources

The SDK supports the following resources:

- **Business Hours**: Manage business hours
- **Companies**: Manage company information
- **Customers**: Manage customer information
- **Help Doc Articles**: Manage help documentation articles
- **Help Doc Sites**: Manage help documentation sites
- **SLAs**: Manage service level agreements
- **Tags**: Manage ticket tags
- **Ticket Priorities**: Manage ticket priorities
- **Ticket Statuses**: Manage ticket statuses
- **Ticket Types**: Manage ticket types
- **Tickets**: Manage support tickets
- **Users**: Manage user accounts

Each resource supports the following operations:

- `Get`: Retrieve a single resource by ID
- `List`: Retrieve a list of resources with optional filters
- `Create`: Create a new resource
- `Update`: Update an existing resource

### Command Line Interface

The SDK includes a command-line interface for quick operations:

```bash
# List tickets
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action list

# Get a specific ticket
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action get --id 123

# Create a new ticket
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action create --data '{"subject": "New Ticket", "description": "Ticket description"}'

# Update a ticket
./desksdkgo --api-key YOUR_API_KEY --resource tickets --action update --id 123 --data '{"status": "resolved"}'
```

### Configuration

The CLI supports the following configuration options:

- `--api-key`: Teamwork Desk API key (required)
- `--base-url`: Teamwork Desk API base URL (default: https://mycompany.teamwork.com/desk/api/v2)
- `--resource`: Resource to interact with (default: tickets)
- `--action`: Action to perform (get, list, create, update) (default: list)
- `--id`: Resource ID for get/update actions
- `--debug`: Enable debug logging
- `--data`: JSON data to merge with default values for create/update actions

All configuration options can be set in multiple ways, in order of precedence:

1. Command-line flags
2. Environment variables
3. `.env` file

#### Environment Variables

The following environment variables are supported:

- `DESK_API_KEY`
- `DESK_BASE_URL`
- `DESK_RESOURCE`
- `DESK_ACTION`
- `DEBUG`

#### .env File Support

You can create a `.env` file in your project directory to set default values. Copy `.env.example` to `.env` and modify the values:

```bash
# Copy the example file
cp .env.example .env

# Edit the .env file with your values
```

Example `.env` file:

```env
# Desk API Configuration
DESK_API_KEY=your_api_key_here
DESK_BASE_URL=https://yourcompany.teamwork.com/desk/api/v2

# Default Resource and Action
DESK_RESOURCE=tickets
DESK_ACTION=list

# Debug Mode (true/false)
DEBUG=false
```

### Filtering

The SDK includes a filter builder for creating complex queries:

```go
filter := client.NewFilter().
    Eq("status", "open").
    And(
        client.NewFilter().Gt("created_at", "2023-01-01"),
        client.NewFilter().Lt("created_at", "2023-12-31"),
    )

tickets, err := c.Tickets.List(ctx, filter.Build())
```

Available filter operators:

- `$eq`: Equal to
- `$ne`: Not equal to
- `$lt`: Less than
- `$lte`: Less than or equal to
- `$gt`: Greater than
- `$gte`: Greater than or equal to
- `$in`: In list
- `$nin`: Not in list
- `$and`: Logical AND
- `$or`: Logical OR

## License

This project is licensed under the MIT License - see the LICENSE file for details.
