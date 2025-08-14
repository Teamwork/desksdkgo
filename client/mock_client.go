package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strings"
	"sync"
)

// MockRoundTripper implements http.RoundTripper for testing
type MockRoundTripper struct {
	mu        sync.Mutex
	requests  []*http.Request
	responses map[string]*http.Response
	basePath  string
}

// MockReadCloser implements io.ReadCloser for testing
type MockReadCloser struct {
	io.Reader
	closed bool
}

// Close implements io.Closer
func (m *MockReadCloser) Close() error {
	m.closed = true
	return nil
}

// NewMockReadCloser creates a new MockReadCloser from a string
func NewMockReadCloser(s string) *MockReadCloser {
	return &MockReadCloser{
		Reader: strings.NewReader(s),
	}
}

// NewMockRoundTripper creates a new mock round tripper
func NewMockRoundTripper() *MockRoundTripper {
	return &MockRoundTripper{
		responses: make(map[string]*http.Response),
	}
}

// Reset clears all stored requests and responses
func (m *MockRoundTripper) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.requests = nil
	m.responses = make(map[string]*http.Response)
}

// getPathKey returns a consistent key for the given method and path
func getPathKey(method, urlPath string) string {
	// Clean the path to handle trailing slashes and normalize
	cleanPath := path.Clean(urlPath)
	// Remove any query parameters
	if idx := strings.Index(cleanPath, "?"); idx != -1 {
		cleanPath = cleanPath[:idx]
	}
	return method + " " + cleanPath
}

// RoundTrip implements http.RoundTripper
func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Store the request
	m.requests = append(m.requests, req)

	// Get the response for this request
	key := getPathKey(req.Method, req.URL.Path)
	resp, ok := m.responses[key]
	if !ok {
		// Try without the base path if it exists
		if m.basePath != "" && strings.HasPrefix(req.URL.Path, m.basePath) {
			key = getPathKey(req.Method, strings.TrimPrefix(req.URL.Path, m.basePath))
			resp, ok = m.responses[key]
		}
		if !ok {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(bytes.NewBufferString("Not Found")),
			}, nil
		}
	}

	return resp, nil
}

// AddResponse adds a mock response for a given method and path
func (m *MockRoundTripper) AddResponse(method, urlPath string, statusCode int, body interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var respBody io.ReadCloser
	switch b := body.(type) {
	case io.ReadCloser:
		respBody = b
	case string:
		respBody = io.NopCloser(strings.NewReader(b))
	default:
		bodyBytes, _ := json.Marshal(body)
		respBody = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	key := getPathKey(method, urlPath)
	m.responses[key] = &http.Response{
		StatusCode: statusCode,
		Body:       respBody,
		Header:     make(http.Header),
	}
}

// GetRequests returns all requests made to the mock
func (m *MockRoundTripper) GetRequests() []*http.Request {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.requests
}
