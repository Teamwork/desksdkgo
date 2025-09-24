# Middleware System

The Desk SDK Go client supports a powerful middleware system that allows you to chain middleware functions to process HTTP requests before they are sent to the API. This enables you to implement cross-cutting concerns like authentication, logging, retry logic, rate limiting, and request/response modification.

## How Middleware Works

Middleware functions are executed in the order they are added to the client. Each middleware receives:

- `context.Context`: The request context which can carry values and cancellation signals
- `*http.Request`: The HTTP request that will be sent (which can be modified)
- `RequestHandler`: The next handler in the chain (which must be called to continue the request)

The middleware function signature is:

```go
type MiddlewareFunc func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error)
type RequestHandler func(ctx context.Context, req *http.Request) (*http.Response, error)
```

## Adding Middleware

You can add middleware to a client using the `WithMiddleware` option during client creation:

```go
client := client.NewClient("https://your-desk-instance.desk.com/api/v2",
    client.WithAPIKey("your-api-key"),
    client.WithMiddleware(client.LoggingMiddleware(logger)),
    client.WithMiddleware(client.RetryMiddleware(3, time.Second)),
)
```

## Built-in Middleware

The SDK provides several pre-built middleware functions:

### LoggingMiddleware

Logs HTTP requests and responses with timing information.

```go
client.WithMiddleware(client.LoggingMiddleware(logger))
```

### RetryMiddleware

Retries failed requests with configurable maximum retries and delay.

```go
client.WithMiddleware(client.RetryMiddleware(maxRetries, retryDelay))
```

### AuthMiddleware

Adds authentication headers to requests.

```go
client.WithMiddleware(client.AuthMiddleware("bearer-token"))
```

### UserAgentMiddleware

Sets a custom User-Agent header.

```go
client.WithMiddleware(client.UserAgentMiddleware("MyApp/1.0"))
```

### RateLimitMiddleware

Implements client-side rate limiting.

```go
client.WithMiddleware(client.RateLimitMiddleware(10.0)) // 10 requests per second
```

### RequestIDMiddleware

Adds unique request IDs to each request.

```go
client.WithMiddleware(client.RequestIDMiddleware())
```

### TimeoutMiddleware

Enforces request timeouts.

```go
client.WithMiddleware(client.TimeoutMiddleware(30 * time.Second))
```

### HeaderMiddleware

Adds custom headers to requests.

```go
client.WithMiddleware(client.HeaderMiddleware(map[string]string{
    "X-Custom-Header": "value",
}))
```

### ConditionalMiddleware

Only executes middleware when a condition is met.

```go
client.WithMiddleware(client.ConditionalMiddleware(
    func(req *http.Request) bool {
        return strings.Contains(req.URL.Path, "tickets")
    },
    client.HeaderMiddleware(map[string]string{
        "X-Ticket-Operation": "true",
    }),
))
```

## Creating Custom Middleware

You can create custom middleware by implementing the `MiddlewareFunc` signature:

```go
func customMiddleware() client.MiddlewareFunc {
    return func(ctx context.Context, req *http.Request, next client.RequestHandler) (*http.Response, error) {
        // Pre-request processing
        req.Header.Set("X-Custom-Header", "value")

        // Call the next handler
        resp, err := next(ctx, req)

        // Post-request processing
        if err == nil {
            // Process successful response
        }

        return resp, err
    }
}
```

### Modifying Request Body

Middleware can also modify the request body:

```go
func requestModificationMiddleware() client.MiddlewareFunc {
    return func(ctx context.Context, req *http.Request, next client.RequestHandler) (*http.Response, error) {
        if req.Method == "POST" && req.Header.Get("Content-Type") == "application/json" {
            if req.Body != nil {
                body, err := io.ReadAll(req.Body)
                if err == nil {
                    var data map[string]interface{}
                    if json.Unmarshal(body, &data) == nil {
                        // Modify the data
                        data["timestamp"] = time.Now().Unix()

                        if modifiedBody, err := json.Marshal(data); err == nil {
                            req.Body = io.NopCloser(bytes.NewBuffer(modifiedBody))
                            req.ContentLength = int64(len(modifiedBody))
                        }
                    }
                }
            }
        }

        return next(ctx, req)
    }
}
```

### Using Context Values

Middleware can add and read values from the context:

```go
type contextKey string
const userIDKey contextKey = "user_id"

func contextMiddleware(userID string) client.MiddlewareFunc {
    return func(ctx context.Context, req *http.Request, next client.RequestHandler) (*http.Response, error) {
        // Add value to context
        ctx = context.WithValue(ctx, userIDKey, userID)

        // Pass modified context to next handler
        return next(ctx, req)
    }
}

// Later, in another middleware or handler:
func someOtherMiddleware() client.MiddlewareFunc {
    return func(ctx context.Context, req *http.Request, next client.RequestHandler) (*http.Response, error) {
        if userID, ok := ctx.Value(userIDKey).(string); ok {
            req.Header.Set("X-User-ID", userID)
        }

        return next(ctx, req)
    }
}
```

## Middleware Execution Order

Middleware is executed in the order it's added to the client. The first middleware added will be the first to receive the request and the last to receive the response.

For example:

```go
client := client.NewClient("...",
    client.WithMiddleware(middleware1), // Executed first
    client.WithMiddleware(middleware2), // Executed second
    client.WithMiddleware(middleware3), // Executed third
)
```

The execution flow is:

1. middleware1 (pre-request)
2. middleware2 (pre-request)
3. middleware3 (pre-request)
4. HTTP request sent
5. middleware3 (post-request)
6. middleware2 (post-request)
7. middleware1 (post-request)

## Best Practices

1. **Order Matters**: Place authentication middleware before logging middleware if you don't want to log sensitive tokens.

2. **Error Handling**: Always handle errors appropriately in middleware. You can choose to return early or let errors propagate.

3. **Context Cancellation**: Respect context cancellation in long-running middleware (like rate limiting).

4. **Request Cloning**: When modifying requests in retry middleware, clone the request to avoid issues with body readers being consumed.

5. **Performance**: Be mindful of middleware performance, especially for high-throughput applications.

## Example Usage

See `examples/middleware_example.go` for a complete example showing how to use various middleware types together.
