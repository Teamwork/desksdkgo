package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/sonh/qs"
	"github.com/teamwork/desksdkgo/models"
)

// TicketService handles ticket-related operations
type TicketService struct {
	*Service[models.TicketResponse, models.TicketsResponse]
	client *Client
}

// NewTicketService creates a new ticket service
func NewTicketService(client *Client) *TicketService {
	return &TicketService{
		Service: NewService[models.TicketResponse, models.TicketsResponse](
			client,
			NewDefaultPathHandlerWithUpdateMethod("tickets", http.MethodPatch),
		),
		client: client,
	}
}

// Get retrieves a ticket by ID
func (s *TicketService) Get(ctx context.Context, id int, params url.Values) (*models.TicketResponse, error) {
	return s.Service.Get(ctx, id, params)
}

// List retrieves a list of tickets with optional filters
func (s *TicketService) List(ctx context.Context, params url.Values) (*models.TicketsResponse, error) {
	return s.Service.List(ctx, params)
}

// Search searches for tickets based on query parameters
func (s *TicketService) Search(ctx context.Context, filter *models.SearchTicketsFilter) (*models.TicketsResponse, error) {
	encoder := qs.NewEncoder()
	values, err := encoder.Values(filter)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/search/tickets.json?%s", s.client.baseURL, values.Encode()), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resources models.TicketsResponse
	if err := json.NewDecoder(resp.Body).Decode(&resources); err != nil {
		return nil, err
	}

	return &resources, nil
}

// Create creates a new ticket
func (s *TicketService) Create(ctx context.Context, ticket *models.TicketResponse) (*models.TicketResponse, error) {
	return s.Service.Create(ctx, ticket)
}

// Update updates an existing ticket
func (s *TicketService) Update(ctx context.Context, id int, ticket *models.TicketResponse) (*models.TicketResponse, error) {
	return s.Service.Update(ctx, id, ticket)
}

// LinkTask links a project task to a ticket
func (s *TicketService) LinkTask(ctx context.Context, ticketID, taskID int) error {
	return s.taskLink(ctx, http.MethodPost, ticketID, taskID)
}

// UnlinkTask removes the link between a project task and a ticket
func (s *TicketService) UnlinkTask(ctx context.Context, ticketID, taskID int) error {
	return s.taskLink(ctx, http.MethodDelete, ticketID, taskID)
}

// taskLink performs a request against the ticket task link endpoint
func (s *TicketService) taskLink(ctx context.Context, method string, ticketID, taskID int) error {
	if ticketID <= 0 {
		return fmt.Errorf("ticketID must be greater than 0")
	}

	if taskID <= 0 {
		return fmt.Errorf("taskID must be greater than 0")
	}

	req, err := http.NewRequestWithContext(ctx, method,
		fmt.Sprintf("%s/tickets/%d/tasks/%d.json", s.client.baseURL, ticketID, taskID), nil)
	if err != nil {
		s.logError("failed to create request", slog.Any("error", err))
		return err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		s.logError("request failed",
			slog.Any("error", err),
			slog.String("method", method),
			slog.String("url", req.URL.String()),
		)
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent {
		return nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logError("failed to read response body",
			slog.Any("error", err),
			slog.Int("status_code", resp.StatusCode),
			slog.String("method", method),
			slog.String("url", req.URL.String()),
		)
		return err
	}

	s.logError("unexpected status code",
		slog.Int("status_code", resp.StatusCode),
		slog.String("method", method),
		slog.String("url", req.URL.String()),
		slog.String("response_body", string(b)),
	)
	return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(b))
}
