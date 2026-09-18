package client

import (
	"context"
	"net/http"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// HelpDocCategoryService handles help doc category-related operations
type HelpDocCategoryService struct {
	*Service[models.HelpDocCategoryResponse, models.HelpDocCategoriesResponse]
	client *Client
}

// NewHelpDocCategoryService creates a new help doc category service
func NewHelpDocCategoryService(client *Client) *HelpDocCategoryService {
	return &HelpDocCategoryService{
		Service: NewService[models.HelpDocCategoryResponse, models.HelpDocCategoriesResponse](
			client, NewDefaultPathHandlerWithUpdateMethod("helpdocscategories", http.MethodPatch),
		),
		client: client,
	}
}

// Get retrieves a help doc category by ID
func (s *HelpDocCategoryService) Get(
	ctx context.Context,
	id int,
	params url.Values,
) (*models.HelpDocCategoryResponse, error) {
	return s.Service.Get(ctx, id, params)
}

// List retrieves a list of help doc categories with optional filters. Narrow it
// to one site with the filter query parameter: {"sites.id": <siteID>}.
func (s *HelpDocCategoryService) List(
	ctx context.Context,
	params url.Values,
) (*models.HelpDocCategoriesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new help doc category
func (s *HelpDocCategoryService) Create(
	ctx context.Context,
	category *models.HelpDocCategoryResponse,
) (*models.HelpDocCategoryResponse, error) {
	return s.Service.Create(ctx, category)
}

// Update updates an existing help doc category
func (s *HelpDocCategoryService) Update(
	ctx context.Context,
	id int,
	category *models.HelpDocCategoryResponse,
) (*models.HelpDocCategoryResponse, error) {
	return s.Service.Update(ctx, id, category)
}
