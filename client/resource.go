package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

// Service handles generic resource operations
type Service[T any, L any] struct {
	client *Client
	router PathHandler
}

type PathHandler interface {
	Get(id int) string
	List() string
	Create() string
	Update(id int) string
}

// NewService creates a new generic service
func NewService[T any, L any](client *Client, router PathHandler) *Service[T, L] {
	return &Service[T, L]{
		client: client,
		router: router,
	}
}

// Get retrieves a resource by ID
func (s *Service[T, L]) Get(ctx context.Context, id int) (*T, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/%s.json?includes=all", s.client.baseURL, s.router.Get(id)), nil)
	if err != nil {
		s.client.logger.Error("failed to create request", slog.Any("error", err))
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		s.client.logger.Error("request failed", slog.Any("error", err), slog.String("method", http.MethodGet), slog.String("url", req.URL.String()))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		s.client.logger.Error("unexpected status code",
			slog.Int("status_code", resp.StatusCode),
			slog.String("method", http.MethodGet),
			slog.String("url", req.URL.String()),
			slog.String("response_body", string(body)),
		)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resource T
	if err := json.NewDecoder(resp.Body).Decode(&resource); err != nil {
		s.client.logger.Error("failed to decode response",
			slog.Any("error", err),
			slog.String("method", http.MethodGet),
			slog.String("url", req.URL.String()),
		)
		return nil, err
	}

	return &resource, nil
}

// List retrieves a list of resources with optional filters
func (s *Service[T, L]) List(ctx context.Context, params url.Values) (*L, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/%s.json?%s", s.client.baseURL, s.router.List(), params.Encode()), nil)
	if err != nil {
		s.client.logger.Error("failed to create request", slog.Any("error", err))
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		s.client.logger.Error("request failed", slog.Any("error", err), slog.String("method", http.MethodGet), slog.String("url", req.URL.String()))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		s.client.logger.Error("unexpected status code",
			slog.Int("status_code", resp.StatusCode),
			slog.String("method", http.MethodGet),
			slog.String("url", req.URL.String()),
			slog.String("response_body", string(body)),
		)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resources L
	if err := json.NewDecoder(resp.Body).Decode(&resources); err != nil {
		s.client.logger.Error("failed to decode response",
			slog.Any("error", err),
			slog.String("method", http.MethodGet),
			slog.String("url", req.URL.String()),
		)
		return nil, err
	}

	return &resources, nil
}

// Create creates a new resource
func (s *Service[T, L]) Create(ctx context.Context, resource *T) (*T, error) {
	body, err := json.Marshal(resource)
	if err != nil {
		s.client.logger.Error("failed to marshal request body", slog.Any("error", err))
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/%s.json", s.client.baseURL, s.router.Create()), bytes.NewBuffer(body))
	if err != nil {
		s.client.logger.Error("failed to create request", slog.Any("error", err))
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		s.client.logger.Error("request failed", slog.Any("error", err), slog.String("method", http.MethodPost), slog.String("url", req.URL.String()))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			s.client.logger.Error("failed to read response body",
				slog.Any("error", err),
				slog.Int("status_code", resp.StatusCode),
				slog.String("method", http.MethodPost),
				slog.String("url", req.URL.String()),
			)
			return nil, err
		}

		s.client.logger.Error("unexpected status code",
			slog.Int("status_code", resp.StatusCode),
			slog.String("method", http.MethodPost),
			slog.String("url", req.URL.String()),
			slog.String("response_body", string(b)),
		)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(b))
	}

	var createdResource T
	if err := json.NewDecoder(resp.Body).Decode(&createdResource); err != nil {
		s.client.logger.Error("failed to decode response",
			slog.Any("error", err),
			slog.String("method", http.MethodPost),
			slog.String("url", req.URL.String()),
		)
		return nil, err
	}

	return &createdResource, nil
}

// Update updates an existing resource
func (s *Service[T, L]) Update(ctx context.Context, id int, resource *T) (*T, error) {
	body, err := json.Marshal(resource)
	if err != nil {
		s.client.logger.Error("failed to marshal request body", slog.Any("error", err))
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut,
		fmt.Sprintf("%s/%s.json", s.client.baseURL, s.router.Update(id)), bytes.NewBuffer(body))
	if err != nil {
		s.client.logger.Error("failed to create request", slog.Any("error", err))
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		s.client.logger.Error("request failed", slog.Any("error", err), slog.String("method", http.MethodPut), slog.String("url", req.URL.String()))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		s.client.logger.Error("unexpected status code",
			slog.Int("status_code", resp.StatusCode),
			slog.String("method", http.MethodPut),
			slog.String("url", req.URL.String()),
			slog.String("response_body", string(body)),
		)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var updatedResource T
	if err := json.NewDecoder(resp.Body).Decode(&updatedResource); err != nil {
		s.client.logger.Error("failed to decode response",
			slog.Any("error", err),
			slog.String("method", http.MethodPut),
			slog.String("url", req.URL.String()),
		)
		return nil, err
	}

	return &updatedResource, nil
}
