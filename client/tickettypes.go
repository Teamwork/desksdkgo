package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// TicketTypeService handles ticket-related operations
type TicketTypeService struct {
	*Service[models.TicketTypeResponse, models.TicketTypesResponse]
}

// NewTicketTypeService creates a new ticket service
func NewTicketTypeService(client *Client) *TicketTypeService {
	return &TicketTypeService{
		Service: NewService[models.TicketTypeResponse, models.TicketTypesResponse](client, "tickettypes"),
	}
}

// Get retrieves a tickettype by ID
func (s *TicketTypeService) Get(ctx context.Context, id int) (*models.TicketTypeResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of tickettypees with optional filters
func (s *TicketTypeService) List(ctx context.Context, params url.Values) (*models.TicketTypesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new tickettype
func (s *TicketTypeService) Create(ctx context.Context, tickettype *models.TicketTypeResponse) (*models.TicketTypeResponse, error) {
	return s.Service.Create(ctx, tickettype)
}

// Update updates an existing tickettype
func (s *TicketTypeService) Update(ctx context.Context, id int, tickettype *models.TicketTypeResponse) (*models.TicketTypeResponse, error) {
	return s.Service.Update(ctx, id, tickettype)
}
