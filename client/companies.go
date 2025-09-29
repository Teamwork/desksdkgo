package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// CompanyService handles company-related operations
type CompanyService struct {
	*Service[models.CompanyResponse, models.CompaniesResponse]
}

// NewCompanyService creates a new company service
func NewCompanyService(client *Client) *CompanyService {
	return &CompanyService{
		Service: NewService[models.CompanyResponse, models.CompaniesResponse](client, NewDefaultPathHandler("companies")),
	}
}

// Get retrieves a company by ID
func (s *CompanyService) Get(ctx context.Context, id int) (*models.CompanyResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of companies with optional filters
func (s *CompanyService) List(ctx context.Context, params url.Values) (*models.CompaniesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new company
func (s *CompanyService) Create(ctx context.Context, company *models.CompanyResponse) (*models.CompanyResponse, error) {
	return s.Service.Create(ctx, company)
}

// Update updates an existing company
func (s *CompanyService) Update(ctx context.Context, id int, company *models.CompanyResponse) (*models.CompanyResponse, error) {
	return s.Service.Update(ctx, id, company)
}
