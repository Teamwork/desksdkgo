package models

// TicketSource related types
type TicketSource struct {
	BaseEntity
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	DisplayOrder int    `json:"displayOrder"`
	IsCustom     bool   `json:"isCustom"`
}

type TicketSourcesResponse struct {
	TicketSources []TicketSource `json:"ticketSources"`
	Meta          Meta           `json:"meta"`
	Pagination    Pagination     `json:"pagination"`
	Included      IncludedData   `json:"included"`
}

type TicketSourceResponse struct {
	TicketSource TicketSource `json:"ticketSource"`
	Included     IncludedData `json:"included"`
}
