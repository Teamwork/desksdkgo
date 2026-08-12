package client

import (
	"context"
	"net/http"
	"testing"
)

func TestTicketServiceLinkTask(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodPost, "/tickets/123/tasks/456.json", http.StatusOK, "")

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	if err := client.Tickets.LinkTask(context.Background(), 123, 456); err != nil {
		t.Fatalf("LinkTask() returned error: %v", err)
	}

	requests := mockTransport.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}

	if requests[0].Method != http.MethodPost {
		t.Errorf("expected method POST, got %s", requests[0].Method)
	}

	if requests[0].URL.Path != "/tickets/123/tasks/456.json" {
		t.Errorf("expected request path /tickets/123/tasks/456.json, got %s", requests[0].URL.Path)
	}
}

func TestTicketServiceUnlinkTask(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodDelete, "/tickets/123/tasks/456.json", http.StatusNoContent, "")

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	if err := client.Tickets.UnlinkTask(context.Background(), 123, 456); err != nil {
		t.Fatalf("UnlinkTask() returned error: %v", err)
	}

	requests := mockTransport.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}

	if requests[0].Method != http.MethodDelete {
		t.Errorf("expected method DELETE, got %s", requests[0].Method)
	}

	if requests[0].URL.Path != "/tickets/123/tasks/456.json" {
		t.Errorf("expected request path /tickets/123/tasks/456.json, got %s", requests[0].URL.Path)
	}
}

func TestTicketServiceLinkTaskValidation(t *testing.T) {
	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: NewMockRoundTripper()}))

	tests := []struct {
		name     string
		ticketID int
		taskID   int
	}{
		{name: "zero ticket ID", ticketID: 0, taskID: 456},
		{name: "negative ticket ID", ticketID: -1, taskID: 456},
		{name: "zero task ID", ticketID: 123, taskID: 0},
		{name: "negative task ID", ticketID: 123, taskID: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := client.Tickets.LinkTask(context.Background(), tt.ticketID, tt.taskID); err == nil {
				t.Error("expected error from LinkTask, got nil")
			}
			if err := client.Tickets.UnlinkTask(context.Background(), tt.ticketID, tt.taskID); err == nil {
				t.Error("expected error from UnlinkTask, got nil")
			}
		})
	}
}

func TestTicketServiceLinkTaskUnexpectedStatus(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodPost, "/tickets/123/tasks/456.json", http.StatusBadRequest, `{"errors":[{"detail":"bad request"}]}`)

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	err := client.Tickets.LinkTask(context.Background(), 123, 456)
	if err == nil {
		t.Fatal("expected error for 400 response, got nil")
	}
}
