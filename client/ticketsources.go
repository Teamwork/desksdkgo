package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// TicketSourceService handles ticket-related operations
type TicketSourceService struct {
	*Service[models.TicketSourceResponse, models.TicketSourcesResponse]
}

// NewTicketSourceService creates a new ticket service
func NewTicketSourceService(client *Client) *TicketSourceService {
	return &TicketSourceService{
		Service: NewService[models.TicketSourceResponse, models.TicketSourcesResponse](client, NewDefaultPathHandler("ticketsources")),
	}
}

// Get retrieves a ticketsource by ID
func (s *TicketSourceService) Get(ctx context.Context, id int) (*models.TicketSourceResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of ticketsources with optional filters
func (s *TicketSourceService) List(ctx context.Context, params url.Values) (*models.TicketSourcesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new ticketsource
func (s *TicketSourceService) Create(ctx context.Context, ticketsource *models.TicketSourceResponse) (*models.TicketSourceResponse, error) {
	return s.Service.Create(ctx, ticketsource)
}

// Update updates an existing ticketsource
func (s *TicketSourceService) Update(ctx context.Context, id int, ticketsource *models.TicketSourceResponse) (*models.TicketSourceResponse, error) {
	return s.Service.Update(ctx, id, ticketsource)
}
