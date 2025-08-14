package models

// TicketPriority related types
type TicketPriority struct {
	BaseEntity
	Name         string `json:"name"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
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
