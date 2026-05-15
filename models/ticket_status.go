package models

// TicketStatus related types
type TicketStatus struct {
	BaseEntity
	Code         *string `json:"code,omitempty"`
	Name         *string `json:"name,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
	IsCustom     *bool   `json:"isCustom,omitempty"`
	Color        *string `json:"color,omitempty"`
	Icon         *string `json:"icon,omitempty"`
}

type TicketStatusesResponse struct {
	TicketStatuses []TicketStatus `json:"ticketstatuses"`
	Meta           Meta           `json:"meta"`
	Pagination     Pagination     `json:"pagination"`
	Included       IncludedData   `json:"included"`
}

type TicketStatusResponse struct {
	TicketStatus TicketStatus `json:"ticketstatus"`
	Included     IncludedData `json:"included"`
}
