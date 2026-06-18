package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sonh/qs"
	"github.com/teamwork/desksdkgo/models"
)

// HelpDocArticleService handles help doc article-related operations
type HelpDocArticleService struct {
	*Service[models.HelpDocArticleResponse, models.HelpDocArticlesResponse]
	client *Client
}

// NewHelpDocArticleService creates a new help doc article service
func NewHelpDocArticleService(client *Client) *HelpDocArticleService {
	return &HelpDocArticleService{
		Service: NewService[models.HelpDocArticleResponse, models.HelpDocArticlesResponse](client, NewDefaultPathHandler("helpdocssites/helpdocarticles")),
		client:  client,
	}
}

// Get retrieves a help doc article by ID
func (s *HelpDocArticleService) Get(ctx context.Context, id int, params url.Values) (*models.HelpDocArticleResponse, error) {
	return s.Service.Get(ctx, id, params)
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

// Search searches for help doc articles based on filter parameters
func (s *HelpDocArticleService) Search(ctx context.Context, filter *models.SearchHelpdocsFilter) (*models.HelpDocArticlesResponse, error) {
	encoder := qs.NewEncoder()
	values, err := encoder.Values(filter)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/search/helpdocs.json?%s", s.client.baseURL, values.Encode()), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var resources models.HelpDocArticlesResponse
	if err := json.NewDecoder(resp.Body).Decode(&resources); err != nil {
		return nil, err
	}

	return &resources, nil
}
