package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/teamwork/desksdkgo/models"
)

// MessageService handles message-related operations
type MessageService struct {
	*Service[models.MessageResponse, models.MessagesResponse]
	client *Client
}

// NewMessageService creates a new message service
func NewMessageService(client *Client) *MessageService {
	return &MessageService{
		Service: NewService[models.MessageResponse, models.MessagesResponse](client, NewDefaultPathHandler("messages")),
		client:  client,
	}
}

// Get retrieves a message by ID
func (s *MessageService) Get(ctx context.Context, id int) (*models.MessageResponse, error) {
	return s.Service.Get(ctx, id)
}

// List retrieves a list of messages with optional filters
func (s *MessageService) List(ctx context.Context, params url.Values) (*models.MessagesResponse, error) {
	return s.Service.List(ctx, params)
}

// Create creates a new message
func (s *MessageService) Create(ctx context.Context, message *models.MessageResponse) (*models.MessageResponse, error) {
	if message == nil {
		return nil, fmt.Errorf("message is required")
	}

	if message.Message.Ticket.ID <= 0 {
		return nil, fmt.Errorf("message.message.ticket.id is required")
	}

	return s.CreateForTicket(ctx, message.Message.Ticket.ID, message)
}

// CreateForTicket creates a new message scoped to a ticket
func (s *MessageService) CreateForTicket(ctx context.Context, ticketID int, message *models.MessageResponse) (*models.MessageResponse, error) {
	if ticketID <= 0 {
		return nil, fmt.Errorf("ticketID must be greater than 0")
	}

	if message == nil {
		return nil, fmt.Errorf("message is required")
	}

	body, err := json.Marshal(message.Message)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("%s/tickets/%d/messages.json", s.client.baseURL, ticketID), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(b))
	}

	var createdMessage models.MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&createdMessage); err != nil {
		return nil, err
	}

	return &createdMessage, nil
}

// Update updates an existing message
func (s *MessageService) Update(ctx context.Context, id int, message *models.MessageResponse) (*models.MessageResponse, error) {
	return s.Service.Update(ctx, id, message)
}
