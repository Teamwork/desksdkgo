package models

import "time"

// Ticket related types
type Ticket struct {
	BaseEntity
	Activities            []EntityRef `json:"activities"`
	Agent                 EntityRef   `json:"agent"`
	Contact               EntityRef   `json:"contact"`
	Customer              EntityRef   `json:"customer"`
	HappinessSurveySentAt time.Time   `json:"happinessSurveySentAt"`
	ImagesHidden          bool        `json:"imagesHidden"`
	Inbox                 EntityRef   `json:"inbox"`
	IsRead                bool        `json:"isRead"`
	MessageCount          int         `json:"messageCount"`
	Messages              []EntityRef `json:"messages"`
	OriginalRecipient     string      `json:"originalRecipient"`
	PreviewText           string      `json:"previewText"`
	Readonly              bool        `json:"readonly"`
	ResolutionTimeMins    int         `json:"resolutionTimeMins"`
	ResponseTimeMins      int         `json:"responseTimeMins"`
	Source                EntityRef   `json:"source"`
	SpamRules             any         `json:"spam_rules"`
	SpamScore             int         `json:"spam_score"`
	Status                EntityRef   `json:"status"`
	Subject               string      `json:"subject"`
	Suggestions           struct{}    `json:"suggestions"`
	Timelogs              []EntityRef `json:"timelogs"`
	Type                  EntityRef   `json:"type"`
}

// Response types for tickets
type TicketsResponse struct {
	Tickets    []Ticket     `json:"tickets"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type TicketResponse struct {
	Ticket   Ticket       `json:"ticket"`
	Included IncludedData `json:"included"`
}
