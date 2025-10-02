package client

import (
	"context"
	"encoding/json"
	"fmt"
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
		Service: NewService[models.TicketResponse, models.TicketsResponse](client, NewDefaultPathHandler("tickets")),
		client:  client,
	}
}

// Get retrieves a ticket by ID
func (s *TicketService) Get(ctx context.Context, id int) (*models.TicketResponse, error) {
	return s.Service.Get(ctx, id)
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
	defer resp.Body.Close()

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
