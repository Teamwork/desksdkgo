package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// InboxService handles ticket-related operations
type InboxService struct {
	*Service[models.InboxResponse, models.InboxesResponse]
}

// NewInboxService creates a new ticket service
func NewInboxService(client *Client) *InboxService {
	return &InboxService{
		Service: NewService[models.InboxResponse, models.InboxesResponse](client, "inboxes"),
	}
}

// Get retrieves a sla by ID
func (s *InboxService) Get(ctx context.Context, id int) (*models.InboxResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of slas with optional filters
func (s *InboxService) List(ctx context.Context, params url.Values) (*models.InboxesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new sla
func (s *InboxService) Create(ctx context.Context, sla *models.InboxResponse) (*models.InboxResponse, error) {
	return s.Service.Create(ctx, sla)
}

// Update updates an existing sla
func (s *InboxService) Update(ctx context.Context, id int, sla *models.InboxResponse) (*models.InboxResponse, error) {
	return s.Service.Update(ctx, id, sla)
}
