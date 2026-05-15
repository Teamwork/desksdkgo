package models

type SLANotificationConditionType string

const (
	SLANotificationConditionTypeWarning SLANotificationConditionType = "warning"
	SLANotificationConditionTypeBreach  SLANotificationConditionType = "breach"
)

type SLANotificationType string

const (
	SLANotificationTypeFirstResponse  SLANotificationType = "firstResponse"
	SLANotificationTypeReplyTime      SLANotificationType = "replyTime"
	SLANotificationTypeResolutionTime SLANotificationType = "resolutionTime"
)

type SLAConditionOption string

const (
	SLAConditionOptionEqual    SLAConditionOption = "eq"
	SLAConditionOptionNotEqual SLAConditionOption = "ne"
)

// SLA represents a SLA in the system
type SLA struct {
	BaseEntity
	Name             *string     `json:"name,omitempty"`
	Description      *string     `json:"description,omitempty"`
	DisplayOrder     *int        `json:"displayOrder,omitempty"`
	Enabled          *bool       `json:"enabled,omitempty"`
	BusinessHour     *EntityRef  `json:"businesshours,omitempty"`
	Customers        []EntityRef `json:"slacustomers"`
	Companies        []EntityRef `json:"slacompanies"`
	Inboxes          []EntityRef `json:"slainboxes"`
	TicketTypes      []EntityRef `json:"slatickettypes"`
	TicketPriorities []EntityRef `json:"slaticketpriorities"`
	Tags             []EntityRef `json:"slatags"`
	Notifications    []EntityRef `json:"slanotifications"`
	Threads          []EntityRef `json:"threads"`
}

type SLANotification struct {
	BaseEntity
	Condition          *SLANotificationConditionType `json:"condition,omitempty" db:"condition"`
	Type               *SLANotificationType          `json:"type,omitempty"`
	Duration           *int                          `json:"duration,omitempty"`
	User               *EntityRef                    `json:"user,omitempty"`
	NotifyAssignedUser *bool                         `json:"notifyAssignedUser,omitempty"`
}

type SLATicketPriority struct {
	BaseEntity
	Hours          *int       `json:"hours,omitempty"`
	Minutes        *int       `json:"minutes,omitempty"`
	Description    *string    `json:"description,omitempty"`
	TicketPriority *EntityRef `json:"priority"`
}

type SLACustomer struct {
	BaseEntity
	Customer  *EntityRef          `json:"customer"`
	Condition *SLAConditionOption `json:"conditionoption,omitempty" db:"conditionoption"`
}

type SLACompany struct {
	BaseEntity
	Company   *EntityRef          `json:"company"`
	Condition *SLAConditionOption `json:"conditionoption,omitempty" db:"conditionoption"`
}

type SLAInbox struct {
	BaseEntity
	Inbox     *EntityRef          `json:"inbox"`
	Condition *SLAConditionOption `json:"conditionoption,omitempty" db:"conditionoption"`
}

type SLATag struct {
	BaseEntity
	Tag       *EntityRef          `json:"tag"`
	Condition *SLAConditionOption `json:"conditionoption,omitempty" db:"conditionoption"`
}

// SLAsResponse represents the response for a list of slas
type SLAsResponse struct {
	SLAs       []SLA        `json:"slas"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type SLAResponse struct {
	SLA      SLA          `json:"sla"`
	Included IncludedData `json:"included"`
}
