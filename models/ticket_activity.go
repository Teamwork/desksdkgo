package models

// TicketActivity related types
type TicketActivity struct {
	BaseEntity
	EventType   *string    `json:"eventType,omitempty"`
	Icon        *string    `json:"icon,omitempty"`
	Color       *string    `json:"color,omitempty"`
	TargetAgent *EntityRef `json:"targetAgent,omitempty"`
	Ticket      EntityRef  `json:"ticket"`
	Inbox       any        `json:"inbox"`
	OldInbox    any        `json:"oldInbox"`
	Status      *EntityRef `json:"status,omitempty"`
}
