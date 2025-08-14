package models

// TicketActivity related types
type TicketActivity struct {
	BaseEntity
	EventType   string    `json:"eventType"`
	Icon        string    `json:"icon"`
	Color       string    `json:"color"`
	TargetAgent EntityRef `json:"targetAgent,omitempty"`
	Ticket      EntityRef `json:"ticket"`
	Inbox       any       `json:"inbox"`
	OldInbox    any       `json:"oldInbox"`
	Status      EntityRef `json:"status,omitempty"`
}
