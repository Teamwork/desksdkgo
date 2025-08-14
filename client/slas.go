package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// SLAService handles ticket-related operations
type SLAService struct {
	*Service[models.SLAResponse, models.SLAsResponse]
}

// NewSLAService creates a new ticket service
func NewSLAService(client *Client) *SLAService {
	return &SLAService{
		Service: NewService[models.SLAResponse, models.SLAsResponse](client, "slas"),
	}
}

// Get retrieves a sla by ID
func (s *SLAService) Get(ctx context.Context, id int) (*models.SLAResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of slas with optional filters
func (s *SLAService) List(ctx context.Context, params url.Values) (*models.SLAsResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new sla
func (s *SLAService) Create(ctx context.Context, sla *models.SLAResponse) (*models.SLAResponse, error) {
	return s.Service.Create(ctx, sla)
}

// Update updates an existing sla
func (s *SLAService) Update(ctx context.Context, id int, sla *models.SLAResponse) (*models.SLAResponse, error) {
	return s.Service.Update(ctx, id, sla)
}
