package client

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware creates middleware that logs HTTP requests and responses
func LoggingMiddleware(logger *slog.Logger) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		start := time.Now()

		logger.InfoContext(ctx, "Making HTTP request",
			slog.String("method", req.Method),
			slog.String("url", req.URL.String()),
		)

		resp, err := next(ctx, req)
		duration := time.Since(start)

		if err != nil {
			logger.ErrorContext(ctx, "HTTP request failed",
				slog.String("method", req.Method),
				slog.String("url", req.URL.String()),
				slog.Duration("duration", duration),
				slog.String("error", err.Error()),
			)
		} else {
			logger.InfoContext(ctx, "HTTP request completed",
				slog.String("method", req.Method),
				slog.String("url", req.URL.String()),
				slog.Int("status", resp.StatusCode),
				slog.Duration("duration", duration),
			)
		}

		return resp, err
	}
}

// RetryMiddleware creates middleware that retries requests on failure
func RetryMiddleware(maxRetries int, retryDelay time.Duration) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		var resp *http.Response
		var err error

		for attempt := 0; attempt <= maxRetries; attempt++ {
			// Clone the request for retry attempts
			clonedReq := req.Clone(ctx)

			resp, err = next(ctx, clonedReq)

			// If successful or on last attempt, return the result
			if err == nil || attempt == maxRetries {
				break
			}

			// Wait before retrying (except on last attempt)
			if attempt < maxRetries {
				select {
				case <-time.After(retryDelay):
				case <-ctx.Done():
					return nil, ctx.Err()
				}
			}
		}

		return resp, err
	}
}

// AuthMiddleware creates middleware that adds authentication headers
func AuthMiddleware(token string) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		return next(ctx, req)
	}
}

// UserAgentMiddleware creates middleware that adds a User-Agent header
func UserAgentMiddleware(userAgent string) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		req.Header.Set("User-Agent", userAgent)
		return next(ctx, req)
	}
}

// RateLimitMiddleware creates middleware that implements rate limiting
func RateLimitMiddleware(requestsPerSecond float64) MiddlewareFunc {
	limiter := make(chan time.Time, 1)
	interval := time.Duration(1000000000 / requestsPerSecond) // Convert to nanoseconds

	// Start the ticker
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for t := range ticker.C {
			select {
			case limiter <- t:
			default:
				// Channel is full, skip this tick
			}
		}
	}()

	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		select {
		case <-limiter:
			return next(ctx, req)
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// RequestIDMiddleware creates middleware that adds a unique request ID header
func RequestIDMiddleware() MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		requestID := fmt.Sprintf("req_%d", time.Now().UnixNano())
		req.Header.Set("X-Request-ID", requestID)
		return next(ctx, req)
	}
}

// TimeoutMiddleware creates middleware that enforces request timeouts
func TimeoutMiddleware(timeout time.Duration) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		req = req.WithContext(timeoutCtx)
		return next(timeoutCtx, req)
	}
}

// HeaderMiddleware creates middleware that adds custom headers to requests
func HeaderMiddleware(headers map[string]string) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
		return next(ctx, req)
	}
}

// ConditionalMiddleware creates middleware that only executes when a condition is met
func ConditionalMiddleware(condition func(*http.Request) bool, middleware MiddlewareFunc) MiddlewareFunc {
	return func(ctx context.Context, req *http.Request, next RequestHandler) (*http.Response, error) {
		if condition(req) {
			return middleware(ctx, req, next)
		}
		return next(ctx, req)
	}
}
