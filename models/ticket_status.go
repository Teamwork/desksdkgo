package models

// TicketStatus related types
type TicketStatus struct {
	BaseEntity
	Code         string `json:"code"`
	Name         string `json:"name"`
	DisplayOrder int    `json:"displayOrder"`
	IsCustom     bool   `json:"isCustom"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
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
