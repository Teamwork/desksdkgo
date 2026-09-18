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
		Service: newHelpDocArticleSiteService(client, 0),
		client:  client,
	}
}

// newHelpDocArticleSiteService builds the generic service for one site's
// articles, or for the cross-site list when siteID is zero.
//
// Every article route but that list is keyed on the site
// ("helpdocssites/{siteID}/helpdocarticles/..."), and the update is a PATCH, so
// the site cannot be baked into one service the way a top-level resource's path
// is. The embedded Service is the site-less one, and serves List alone.
func newHelpDocArticleSiteService(
	client *Client,
	siteID int,
) *Service[models.HelpDocArticleResponse, models.HelpDocArticlesResponse] {
	base := "helpdocssites/helpdocarticles"
	if siteID > 0 {
		base = fmt.Sprintf("helpdocssites/%d/helpdocarticles", siteID)
	}
	return NewService[models.HelpDocArticleResponse, models.HelpDocArticlesResponse](client,
		NewDefaultPathHandlerWithUpdateMethod(base, http.MethodPatch))
}

// Get is not reachable without a site: an article is addressed through the site
// it belongs to, so there is nothing to read from an ID alone. Use GetForSite.
//
// It is kept to the shape the other services share, rather than dropped, so a
// caller ranging over services still compiles — and is told what to call.
func (s *HelpDocArticleService) Get(_ context.Context, _ int, _ url.Values) (*models.HelpDocArticleResponse, error) {
	return nil, fmt.Errorf("help doc articles are addressed through their site: use GetForSite")
}

// GetForSite retrieves a help doc article by ID from the site it belongs to.
func (s *HelpDocArticleService) GetForSite(
	ctx context.Context,
	siteID, id int,
	params url.Values,
) (*models.HelpDocArticleResponse, error) {
	if siteID <= 0 {
		return nil, fmt.Errorf("siteID must be greater than 0")
	}
	return newHelpDocArticleSiteService(s.client, siteID).Get(ctx, id, params)
}

// List retrieves a list of help doc articles across the sites the caller can
// see, with optional filters. It is the one article route that is not keyed on
// a site.
func (s *HelpDocArticleService) List(ctx context.Context, params url.Values) (*models.HelpDocArticlesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new help doc article on the site its helpdocsite reference
// names.
func (s *HelpDocArticleService) Create(
	ctx context.Context,
	article *models.HelpDocArticleResponse,
) (*models.HelpDocArticleResponse, error) {
	if article == nil {
		return nil, fmt.Errorf("article is required")
	}
	if article.HelpDocArticle.Helpdocsite.ID <= 0 {
		return nil, fmt.Errorf("article.helpdocarticle.helpdocsite.id is required")
	}
	return s.CreateForSite(ctx, article.HelpDocArticle.Helpdocsite.ID, article)
}

// CreateForSite creates a new help doc article on the given site.
//
// The route carries the site and the endpoint takes it from there, so what the
// body's helpdocsite says is ignored.
func (s *HelpDocArticleService) CreateForSite(
	ctx context.Context,
	siteID int,
	article *models.HelpDocArticleResponse,
) (*models.HelpDocArticleResponse, error) {
	if siteID <= 0 {
		return nil, fmt.Errorf("siteID must be greater than 0")
	}
	if article == nil {
		return nil, fmt.Errorf("article is required")
	}
	return newHelpDocArticleSiteService(s.client, siteID).Create(ctx, article)
}

// Update updates an existing help doc article on the site its helpdocsite
// reference names.
func (s *HelpDocArticleService) Update(
	ctx context.Context,
	id int,
	article *models.HelpDocArticleResponse,
) (*models.HelpDocArticleResponse, error) {
	if article == nil {
		return nil, fmt.Errorf("article is required")
	}
	if article.HelpDocArticle.Helpdocsite.ID <= 0 {
		return nil, fmt.Errorf("article.helpdocarticle.helpdocsite.id is required")
	}
	return s.UpdateForSite(ctx, article.HelpDocArticle.Helpdocsite.ID, id, article)
}

// UpdateForSite updates an existing help doc article on the given site.
//
// The endpoint binds the body over the stored article and revalidates the
// result, so only the properties the body carries change — and one carried as
// its zero value is an erasure rather than a no-op.
func (s *HelpDocArticleService) UpdateForSite(
	ctx context.Context,
	siteID, id int,
	article *models.HelpDocArticleResponse,
) (*models.HelpDocArticleResponse, error) {
	if siteID <= 0 {
		return nil, fmt.Errorf("siteID must be greater than 0")
	}
	if article == nil {
		return nil, fmt.Errorf("article is required")
	}
	return newHelpDocArticleSiteService(s.client, siteID).Update(ctx, id, article)
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
