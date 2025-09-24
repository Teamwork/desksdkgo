package client

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// LoggingTransport wraps an http.RoundTripper and logs the request and response
type LoggingTransport struct {
	Transport http.RoundTripper
	Logger    *slog.Logger
}

// RoundTrip implements the http.RoundTripper interface
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log request
	attrs := []slog.Attr{
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		slog.Any("headers", req.Header),
	}

	// Read and log request body if present
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			t.Logger.Error("Failed to read request body", slog.Any("error", err))
		} else {
			attrs = append(attrs, slog.String("request_body", string(bodyBytes)))
			// Restore the request body
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	t.Logger.LogAttrs(nil, slog.LevelDebug, "HTTP Request", attrs...)

	// Make the request
	start := time.Now()
	resp, err := t.Transport.RoundTrip(req)
	duration := time.Since(start)

	// Log response
	respAttrs := []slog.Attr{
		slog.Int("status_code", resp.StatusCode),
		slog.String("duration", duration.String()),
		slog.Any("headers", resp.Header),
	}

	// Read and log response body if present
	if resp.Body != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Logger.Error("Failed to read response body", slog.Any("error", err))
		} else {
			respAttrs = append(respAttrs, slog.String("response_body", string(bodyBytes)))
			// Restore the response body
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	t.Logger.LogAttrs(nil, slog.LevelDebug, "HTTP Response", respAttrs...)

	return resp, err
}

// NewLoggingClient creates a new HTTP client with logging
func NewLoggingClient(level slog.Level) *http.Client {
	return NewLoggingClientWithLogger(level, nil)
}

// NewLoggingClientWithLogger creates a new HTTP client with logging using a custom logger
func NewLoggingClientWithLogger(level slog.Level, logger *slog.Logger) *http.Client {
	if logger == nil {
		logger = slog.New(slog.NewJSONHandler(nil, &slog.HandlerOptions{
			Level: level,
		}))
	}

	transport := &LoggingTransport{
		Transport: http.DefaultTransport,
		Logger:    logger,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}
}
