package models

// TicketType related types
type TicketType struct {
	BaseEntity
	Name                    *string     `json:"name,omitempty"`
	DisplayOrder            *int        `json:"displayOrder,omitempty"`
	Inboxes                 []EntityRef `json:"inboxes"`
	Default                 *bool       `json:"default,omitempty"`
	EnabledForFutureInboxes *bool       `json:"enabledForFutureInboxes,omitempty"`
}

type TicketTypesResponse struct {
	TicketTypes []TicketType `json:"tickettypes"`
	Meta        Meta         `json:"meta"`
	Pagination  Pagination   `json:"pagination"`
	Included    IncludedData `json:"included"`
}

type TicketTypeResponse struct {
	TicketType TicketType   `json:"tickettype"`
	Included   IncludedData `json:"included"`
}
