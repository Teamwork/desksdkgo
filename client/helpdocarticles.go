package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// HelpDocArticleService handles help doc article-related operations
type HelpDocArticleService struct {
	*Service[models.HelpDocArticleResponse, models.HelpDocArticlesResponse]
}

// NewHelpDocArticleService creates a new help doc article service
func NewHelpDocArticleService(client *Client) *HelpDocArticleService {
	return &HelpDocArticleService{
		Service: NewService[models.HelpDocArticleResponse, models.HelpDocArticlesResponse](client, NewDefaultPathHandler("helpdocssites/helpdocarticles")),
	}
}

// Get retrieves a help doc article by ID
func (s *HelpDocArticleService) Get(ctx context.Context, id int) (*models.HelpDocArticleResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of help doc articles with optional filters
func (s *HelpDocArticleService) List(ctx context.Context, params url.Values) (*models.HelpDocArticlesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new help doc article
func (s *HelpDocArticleService) Create(ctx context.Context, article *models.HelpDocArticleResponse) (*models.HelpDocArticleResponse, error) {
	return s.Service.Create(ctx, article)
}

// Update updates an existing help doc article
func (s *HelpDocArticleService) Update(ctx context.Context, id int, article *models.HelpDocArticleResponse) (*models.HelpDocArticleResponse, error) {
	return s.Service.Update(ctx, id, article)
}
