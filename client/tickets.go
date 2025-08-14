package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// TicketService handles ticket-related operations
type TicketService struct {
	*Service[models.TicketResponse, models.TicketsResponse]
}

// NewTicketService creates a new ticket service
func NewTicketService(client *Client) *TicketService {
	return &TicketService{
		Service: NewService[models.TicketResponse, models.TicketsResponse](client, "tickets"),
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

// Create creates a new ticket
func (s *TicketService) Create(ctx context.Context, ticket *models.TicketResponse) (*models.TicketResponse, error) {
	return s.Service.Create(ctx, ticket)
}

// Update updates an existing ticket
func (s *TicketService) Update(ctx context.Context, id int, ticket *models.TicketResponse) (*models.TicketResponse, error) {
	return s.Service.Update(ctx, id, ticket)
}
