package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/teamwork/desksdkgo/models"
)

func TestMessageServiceCreateForTicket(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodPost, "/tickets/123/messages.json", http.StatusCreated, models.MessageResponse{
		Message: models.Message{
			BaseEntity: models.BaseEntity{ID: 1},
			Ticket:     models.EntityRef{ID: 123},
		},
	})

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	resp, err := client.Messages.CreateForTicket(context.Background(), 123, &models.MessageResponse{
		Message: models.Message{TextBody: "hello"},
	})
	if err != nil {
		t.Fatalf("CreateForTicket() returned error: %v", err)
	}

	if resp == nil {
		t.Fatal("CreateForTicket() returned nil response")
	}

	if resp.Message.ID != 1 {
		t.Fatalf("expected created message ID 1, got %d", resp.Message.ID)
	}

	requests := mockTransport.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}

	if requests[0].URL.Path != "/tickets/123/messages.json" {
		t.Fatalf("expected request path /tickets/123/messages.json, got %s", requests[0].URL.Path)
	}
}

func TestMessageServiceCreateUsesMessageTicketID(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodPost, "/tickets/321/messages.json", http.StatusCreated, models.MessageResponse{
		Message: models.Message{
			BaseEntity: models.BaseEntity{ID: 2},
			Ticket:     models.EntityRef{ID: 321},
		},
	})

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	_, err := client.Messages.Create(context.Background(), &models.MessageResponse{
		Message: models.Message{
			Ticket:   models.EntityRef{ID: 321},
			TextBody: "reply",
		},
	})
	if err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	requests := mockTransport.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}

	if requests[0].URL.Path != "/tickets/321/messages.json" {
		t.Fatalf("expected request path /tickets/321/messages.json, got %s", requests[0].URL.Path)
	}
}

func TestMessageServiceCreateRequiresTicketID(t *testing.T) {
	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: NewMockRoundTripper()}))

	_, err := client.Messages.Create(context.Background(), &models.MessageResponse{
		Message: models.Message{TextBody: "no ticket"},
	})
	if err == nil {
		t.Fatal("expected error when ticket ID is missing")
	}
}
