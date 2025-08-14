package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// HelpDocSiteService handles help doc site-related operations
type HelpDocSiteService struct {
	*Service[models.HelpDocSiteResponse, models.HelpDocSitesResponse]
}

// NewHelpDocSiteService creates a new help doc site service
func NewHelpDocSiteService(client *Client) *HelpDocSiteService {
	return &HelpDocSiteService{
		Service: NewService[models.HelpDocSiteResponse, models.HelpDocSitesResponse](client, "helpdocssites"),
	}
}

// Get retrieves a help doc site by ID
func (s *HelpDocSiteService) Get(ctx context.Context, id int) (*models.HelpDocSiteResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of help doc sites with optional filters
func (s *HelpDocSiteService) List(ctx context.Context, params url.Values) (*models.HelpDocSitesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new ticket
func (s *HelpDocSiteService) Create(ctx context.Context, helpDocSite *models.HelpDocSiteResponse) (*models.HelpDocSiteResponse, error) {
	return s.Service.Create(ctx, helpDocSite)
}

// Update updates an existing ticket
func (s *HelpDocSiteService) Update(ctx context.Context, id int, helpDocSite *models.HelpDocSiteResponse) (*models.HelpDocSiteResponse, error) {
	return s.Service.Update(ctx, id, helpDocSite)
}
