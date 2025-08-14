package client

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggingTransport wraps an http.RoundTripper and logs the request and response
type LoggingTransport struct {
	Transport http.RoundTripper
	Logger    *logrus.Logger
}

// RoundTrip implements the http.RoundTripper interface
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log request
	fields := logrus.Fields{
		"method":  req.Method,
		"url":     req.URL.String(),
		"headers": req.Header,
	}

	// Read and log request body if present
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			t.Logger.WithError(err).Error("Failed to read request body")
		} else {
			fields["request_body"] = string(bodyBytes)
			// Restore the request body
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	t.Logger.WithFields(fields).Debug("HTTP Request")

	// Make the request
	start := time.Now()
	resp, err := t.Transport.RoundTrip(req)
	duration := time.Since(start)

	// Log response
	respFields := logrus.Fields{
		"status_code": resp.StatusCode,
		"duration":    duration.String(),
		"headers":     resp.Header,
	}

	// Read and log response body if present
	if resp.Body != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Logger.WithError(err).Error("Failed to read response body")
		} else {
			respFields["response_body"] = string(bodyBytes)
			// Restore the response body
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	t.Logger.WithFields(respFields).Debug("HTTP Response")

	return resp, err
}

// NewLoggingClient creates a new HTTP client with logging
func NewLoggingClient(level logrus.Level) *http.Client {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logger.SetLevel(level)

	transport := &LoggingTransport{
		Transport: http.DefaultTransport,
		Logger:    logger,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}
}
