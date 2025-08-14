package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// UserService handles user-related operations
type UserService struct {
	*Service[models.UserResponse, models.UsersResponse]
}

// NewUserService creates a new user service
func NewUserService(client *Client) *UserService {
	return &UserService{
		Service: NewService[models.UserResponse, models.UsersResponse](client, "users"),
	}
}

// Get retrieves a user by ID
func (s *UserService) Get(ctx context.Context, id int) (*models.UserResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of users with optional filters
func (s *UserService) List(ctx context.Context, params url.Values) (*models.UsersResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new user
func (s *UserService) Create(ctx context.Context, user *models.UserResponse) (*models.UserResponse, error) {
	return s.Service.Create(ctx, user)
}

// Update updates an existing user
func (s *UserService) Update(ctx context.Context, id int, user *models.UserResponse) (*models.UserResponse, error) {
	return s.Service.Update(ctx, id, user)
}
