package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// CustomerService handles customer-related operations
type CustomerService struct {
	*Service[models.CustomerResponse, models.CustomersResponse]
}

// NewCustomerService creates a new customer service
func NewCustomerService(client *Client) *CustomerService {
	return &CustomerService{
		Service: NewService[models.CustomerResponse, models.CustomersResponse](client, "customers"),
	}
}

// Get retrieves a customer by ID
func (s *CustomerService) Get(ctx context.Context, id int) (*models.CustomerResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of customers with optional filters
func (s *CustomerService) List(ctx context.Context, params url.Values) (*models.CustomersResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new customer
func (s *CustomerService) Create(ctx context.Context, customer *models.CustomerResponse) (*models.CustomerResponse, error) {
	return s.Service.Create(ctx, customer)
}

// Update updates an existing customer
func (s *CustomerService) Update(ctx context.Context, id int, customer *models.CustomerResponse) (*models.CustomerResponse, error) {
	return s.Service.Update(ctx, id, customer)
}
