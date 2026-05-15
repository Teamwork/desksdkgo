package models

// TicketSource related types
type TicketSource struct {
	BaseEntity
	Name         *string `json:"name,omitempty"`
	Icon         *string `json:"icon,omitempty"`
	DisplayOrder *int    `json:"displayOrder,omitempty"`
	IsCustom     *bool   `json:"isCustom,omitempty"`
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
