package models

// TicketType related types
type TicketType struct {
	BaseEntity
	Name                    string      `json:"name"`
	DisplayOrder            int         `json:"displayOrder"`
	Inboxes                 []EntityRef `json:"inboxes"`
	Default                 bool        `json:"default"`
	EnabledForFutureInboxes bool        `json:"enabledForFutureInboxes"`
}

type TicketTypesResponse struct {
	TicketTypes []TicketStatus `json:"tickettypes"`
	Meta        Meta           `json:"meta"`
	Pagination  Pagination     `json:"pagination"`
	Included    IncludedData   `json:"included"`
}

type TicketTypeResponse struct {
	TicketType TicketType   `json:"tickettype"`
	Included   IncludedData `json:"included"`
}
