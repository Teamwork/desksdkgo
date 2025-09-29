package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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

// Create creates a new file reference.  This does not upload the file to s3,
// but returns the necessary information to do so.
func (s *FileService) Create(ctx context.Context, file *models.FileResponse) (*models.FileResponse, error) {
	return s.Service.Create(ctx, file)
}

// Upload uploads a file to s3.  This is a helper method that uses the
// information returned from the Create method.
func (s *FileService) Upload(ctx context.Context, file *models.FileResponse, f []byte) error {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	fields := map[string]string{
		"Content-Type":          file.Params.ContentType,
		"bucket":                file.Params.Bucket,
		"key":                   file.Params.Key,
		"policy":                file.Params.Policy,
		"success_action_status": file.Params.SuccessActionStatus,
		"x-amz-algorithm":       file.Params.XAmzAlgorithm,
		"x-amz-credential":      file.Params.XAmzCredential,
		"x-amz-date":            file.Params.XAmzDate,
		"x-amz-signature":       file.Params.XAmzSignature,
	}

	for k, v := range fields {
		if v != "" {
			err := writer.WriteField(k, v)
			if err != nil {
				return err
			}
		}
	}

	part, err := writer.CreateFormFile("file", file.File.Filename)
	if err != nil {
		return fmt.Errorf("create form file: %w", err)
	}

	_, err = io.Copy(part, bytes.NewReader(f))
	if err != nil {
		return fmt.Errorf("copy file data: %w", err)
	}

	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, file.URL, &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to upload file, status code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, body)
	}

	return nil
}

// Update updates an existing file
func (s *FileService) Update(ctx context.Context, id int, file *models.FileResponse) (*models.FileResponse, error) {
	return s.Service.Update(ctx, id, file)
}
