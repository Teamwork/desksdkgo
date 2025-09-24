package client

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestWithLogger(t *testing.T) {
	// Create a buffer to capture logger output
	var buf bytes.Buffer
	customLogger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Create client with custom logger (don't provide HTTP client so it creates one with our logger)
	c := NewClient(
		"https://test.teamwork.com/desk/api/v2",
		WithAPIKey("test-key"),
		WithLogger(customLogger),
	)

	// Check that the client has the logger set
	if c.logger != customLogger {
		t.Error("Expected client to have custom logger set")
	}
}

func TestWithoutLogger(t *testing.T) {
	// Create client without custom logger but with log level
	c := NewClient(
		"https://test.teamwork.com/desk/api/v2",
		WithAPIKey("test-key"),
		WithLogLevel(slog.LevelDebug),
	)

	// Check that the client doesn't have a custom logger set (should be nil)
	if c.logger != nil {
		t.Error("Expected client logger to be nil when not provided")
	}
}
