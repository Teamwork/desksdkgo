package client

import (
	"context"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// FileService handles ticket-related operations
type FileService struct {
	*Service[models.FileResponse, models.FilesResponse]
}

type FilePathHandler struct {
	DefaultPathHandler
}

func NewFilePathHandler() FilePathHandler {
	return FilePathHandler{DefaultPathHandler: NewDefaultPathHandler("files")}
}

func (f FilePathHandler) Create() string {
	return "files/ref"
}

// NewFileService creates a new ticket service
func NewFileService(client *Client) *FileService {
	return &FileService{
		Service: NewService[models.FileResponse, models.FilesResponse](client, NewFilePathHandler()),
	}
}

// Get retrieves a file by ID
func (s *FileService) Get(ctx context.Context, id int) (*models.FileResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of files with optional filters
func (s *FileService) List(ctx context.Context, params url.Values) (*models.FilesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new file
func (s *FileService) Create(ctx context.Context, file *models.FileResponse) (*models.FileResponse, error) {
	return s.Service.Create(ctx, file)
}

// Update updates an existing file
func (s *FileService) Update(ctx context.Context, id int, file *models.FileResponse) (*models.FileResponse, error) {
	return s.Service.Update(ctx, id, file)
}
