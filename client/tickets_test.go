package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/teamwork/desksdkgo/models"
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

func TestTicketServiceSearchDecodesTaskStatuses(t *testing.T) {
	// The API returns `status` inside linked-task metadata as either a status
	// name ("completed") or a status id (3), so both shapes must decode. The
	// same response mixes integer and float spam_score values.
	body := `{"tickets":[
		{"id":2143999,"type":{"id":13630,"type":"tickettypes"},"status":{"id":4,"type":"ticketstatuses"},"spam_score":0,
		 "tasks":[{"id":28423,"type":"tasktickets","meta":{"completed":true,"project":{"id":349856,"type":"projects"},"stateChanged":false,"status":"completed","task":{"completed":true,"id":26948402,"stateChanged":false,"status":"completed","type":"tasks"}}}],
		 "project":{"id":349856,"type":"projects"},"suggestions":{},"state":"active","sharedWith":[]},
		{"id":1894825,"spam_score":0.7,
		 "tasks":[{"id":8285,"type":"sites","meta":{"completed":false,"project":{"id":349856,"type":"oauth2tokens"},"stateChanged":false,"status":1,"task":{"completed":false,"id":26150701,"stateChanged":false,"status":1,"type":"sentiments"}}}],
		 "suggestions":{},"state":"active"}
	]}`

	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodGet, "/search/tickets.json", http.StatusOK, body)

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	resp, err := client.Tickets.Search(context.Background(), &models.SearchTicketsFilter{
		Customers: []int64{357836},
	})
	if err != nil {
		t.Fatalf("Search() returned error: %v", err)
	}

	if len(resp.Tickets) != 2 {
		t.Fatalf("expected 2 tickets, got %d", len(resp.Tickets))
	}

	named := resp.Tickets[0].Tasks
	if len(named) != 1 {
		t.Fatalf("expected 1 task on first ticket, got %d", len(named))
	}
	if named[0].Meta.Project.ID != 349856 {
		t.Errorf("expected project id 349856, got %d", named[0].Meta.Project.ID)
	}
	if named[0].Meta.Status != "completed" {
		t.Errorf("expected status \"completed\", got %v", named[0].Meta.Status)
	}
	if named[0].Meta.Task.ID != 26948402 {
		t.Errorf("expected task id 26948402, got %d", named[0].Meta.Task.ID)
	}

	if resp.Tickets[0].SpamScore == nil || *resp.Tickets[0].SpamScore != 0 {
		t.Errorf("expected integer spam_score 0, got %v", resp.Tickets[0].SpamScore)
	}
	if resp.Tickets[1].SpamScore == nil || *resp.Tickets[1].SpamScore != 0.7 {
		t.Errorf("expected float spam_score 0.7, got %v", resp.Tickets[1].SpamScore)
	}

	numbered := resp.Tickets[1].Tasks
	if len(numbered) != 1 {
		t.Fatalf("expected 1 task on second ticket, got %d", len(numbered))
	}
	if numbered[0].Meta.Task.ID != 26150701 {
		t.Errorf("expected task id 26150701, got %d", numbered[0].Meta.Task.ID)
	}
	if got, ok := numbered[0].Meta.Status.(float64); !ok || got != 1 {
		t.Errorf("expected status 1, got %v", numbered[0].Meta.Status)
	}
	if got, ok := numbered[0].Meta.Task.Status.(float64); !ok || got != 1 {
		t.Errorf("expected inner task status 1, got %v", numbered[0].Meta.Task.Status)
	}
}
