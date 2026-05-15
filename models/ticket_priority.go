package models

// TicketPriority related types
type TicketPriority struct {
	BaseEntity
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
}

type TicketPrioritiesResponse struct {
	TicketPriorities []TicketStatus `json:"ticketpriorities"`
	Meta             Meta           `json:"meta"`
	Pagination       Pagination     `json:"pagination"`
	Included         IncludedData   `json:"included"`
}

type TicketPriorityResponse struct {
	TicketPriority TicketPriority `json:"ticketpriority"`
	Included       IncludedData   `json:"included"`
}
