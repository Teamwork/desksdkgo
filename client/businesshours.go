package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// BusinessHourService handles ticket-related operations
type BusinessHourService struct {
	*Service[models.BusinessHourResponse, models.BusinessHoursResponse]
}

// NewBusinessHourService creates a new ticket service
func NewBusinessHourService(client *Client) *BusinessHourService {
	return &BusinessHourService{
		Service: NewService[models.BusinessHourResponse, models.BusinessHoursResponse](client, "businesshours"),
	}
}

// Get retrieves a businesshour by ID
func (s *BusinessHourService) Get(ctx context.Context, id int) (*models.BusinessHourResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of businesshours with optional filters
func (s *BusinessHourService) List(ctx context.Context, params url.Values) (*models.BusinessHoursResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new businesshour
func (s *BusinessHourService) Create(ctx context.Context, businesshour *models.BusinessHourResponse) (*models.BusinessHourResponse, error) {
	return s.Service.Create(ctx, businesshour)
}

// Update updates an existing businesshour
func (s *BusinessHourService) Update(ctx context.Context, id int, businesshour *models.BusinessHourResponse) (*models.BusinessHourResponse, error) {
	return s.Service.Update(ctx, id, businesshour)
}
