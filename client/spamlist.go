package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// SpamlistService handles spamlist-related operations
type SpamlistService struct {
	*Service[models.SpamlistResponse, models.SpamlistsResponse]
}

// NewSpamlistService creates a new spamlist service
func NewSpamlistService(client *Client) *SpamlistService {
	return &SpamlistService{
		Service: NewService[models.SpamlistResponse, models.SpamlistsResponse](client, "spamlists"),
	}
}

// Get retrieves a spamlist by ID
func (s *SpamlistService) Get(ctx context.Context, id int) (*models.SpamlistResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of spamlistes with optional filters
func (s *SpamlistService) List(ctx context.Context, params url.Values) (*models.SpamlistsResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new spamlist
func (s *SpamlistService) Create(ctx context.Context, spamlist *models.SpamlistResponse) (*models.SpamlistResponse, error) {
	return s.Service.Create(ctx, spamlist)
}

// Update updates an existing spamlist
func (s *SpamlistService) Update(ctx context.Context, id int, spamlist *models.SpamlistResponse) (*models.SpamlistResponse, error) {
	return s.Service.Update(ctx, id, spamlist)
}
