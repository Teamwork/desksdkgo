package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// TagService handles ticket-related operations
type TagService struct {
	*Service[models.TagResponse, models.TagsResponse]
}

// NewTagService creates a new ticket service
func NewTagService(client *Client) *TagService {
	return &TagService{
		Service: NewService[models.TagResponse, models.TagsResponse](client, NewDefaultPathHandler("tags")),
	}
}

// Get retrieves a tag by ID
func (s *TagService) Get(ctx context.Context, id int) (*models.TagResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of tages with optional filters
func (s *TagService) List(ctx context.Context, params url.Values) (*models.TagsResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new tag
func (s *TagService) Create(ctx context.Context, tag *models.TagResponse) (*models.TagResponse, error) {
	return s.Service.Create(ctx, tag)
}

// Update updates an existing tag
func (s *TagService) Update(ctx context.Context, id int, tag *models.TagResponse) (*models.TagResponse, error) {
	return s.Service.Update(ctx, id, tag)
}
