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
		Service: NewService[models.InboxResponse, models.InboxesResponse](client, NewDefaultPathHandler("inboxes")),
	}
}

// Get retrieves an inbox by ID
func (s *InboxService) Get(ctx context.Context, id int) (*models.InboxResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of inboxes with optional filters
func (s *InboxService) List(ctx context.Context, params url.Values) (*models.InboxesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new inbox
func (s *InboxService) Create(ctx context.Context, inbox *models.InboxResponse) (*models.InboxResponse, error) {
	return s.Service.Create(ctx, inbox)
}

// Update updates an existing inbox
func (s *InboxService) Update(ctx context.Context, id int, inbox *models.InboxResponse) (*models.InboxResponse, error) {
	return s.Service.Update(ctx, id, inbox)
}
