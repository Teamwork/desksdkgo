package models

import "time"

// Ticket related types
type Ticket struct {
	BaseEntity
	Activities            []EntityRef `json:"activities,omitempty"`
	Agent                 *EntityRef  `json:"agent,omitempty"`
	BCC                   []string    `json:"bcc,omitempty"`
	Body                  string      `json:"message"`
	CC                    []string    `json:"cc,omitempty"`
	Contact               *EntityRef  `json:"contact,omitempty"`
	Customer              *EntityRef  `json:"customer,omitempty"`
	Files                 []EntityRef `json:"files,omitempty"`
	HappinessSurveySentAt *time.Time  `json:"happinessSurveySentAt"`
	ImagesHidden          bool        `json:"imagesHidden"`
	Inbox                 *EntityRef  `json:"inbox,omitempty"`
	IsRead                bool        `json:"isRead"`
	MessageCount          int         `json:"messageCount"`
	Messages              []EntityRef `json:"messages,omitempty"`
	NotifyCustomer        bool        `json:"notifyCustomer"`
	OriginalRecipient     string      `json:"originalRecipient"`
	PreviewText           string      `json:"previewText"`
	Priority              *EntityRef  `json:"priority,omitempty"`
	Readonly              bool        `json:"readonly"`
	ResolutionTimeMins    int         `json:"resolutionTimeMins"`
	ResponseTimeMins      int         `json:"responseTimeMins"`
	Source                *EntityRef  `json:"source,omitempty"`
	SpamRules             any         `json:"spam_rules"`
	SpamScore             float64     `json:"spam_score"`
	Status                *EntityRef  `json:"status,omitempty"`
	Subject               string      `json:"subject"`
	Suggestions           struct{}    `json:"suggestions"`
	Tags                  []EntityRef `json:"tags,omitempty"`
	Tasks                 []Task      `json:"tasks,omitempty"`
	Timelogs              []EntityRef `json:"timelogs,omitempty"`
	Type                  *EntityRef  `json:"type,omitempty"`
}

// Response types for tickets
type TicketsResponse struct {
	Tickets    []Ticket     `json:"tickets"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type TicketResponse struct {
	Ticket   Ticket       `json:"ticket"`
	Included IncludedData `json:"included"`
}

type CustomFieldsSearch []CustomFieldSearch

type CustomFieldSearch struct {
	ID        int64   `qs:"id"`
	Value     string  `qs:"value"`
	Values    []int64 `qs:"values"`
	Operation string  `qs:"operation"`
}

type Task struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Meta struct {
		Completed bool `json:"completed"`
		Project   struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
		} `json:"project"`
		StateChanged bool   `json:"stateChanged"`
		Status       string `json:"status"`
		Task         struct {
			Completed    bool   `json:"completed"`
			ID           int    `json:"id"`
			StateChanged bool   `json:"stateChanged"`
			Status       string `json:"status"`
			Type         string `json:"type"`
		} `json:"task"`
	} `json:"meta"`
}

type SearchTicketsFilter struct {
	Agents                []int64            `qs:"agents"`
	Companies             []int64            `qs:"companies"`
	Customers             []int64            `qs:"customers"`
	CustomFields          CustomFieldsSearch `qs:"customfields"`
	EndDate               *time.Time         `qs:"endDate,omitempty"`
	Exact                 bool               `qs:"exact"`
	ExcludeInboxes        []int64            `qs:"excludeInboxes"`
	ExcludeTags           []int64            `qs:"excludeTags"`
	ExcludeWorkEmails     bool               `qs:"excludeWorkEmails"`
	Filter                string             `qs:"filter"`
	HelpdocSites          []int64            `qs:"helpdocSites"`
	Inboxes               []int64            `qs:"inboxes"`
	IncludeArchivedAgents bool               `qs:"includeArchivedAgents"`
	LastUpdated           *time.Time         `qs:"lastUpdated,omitempty"`
	OmitMerged            bool               `qs:"omitMerged"`
	OnlyUntagged          bool               `qs:"onlyUntagged"`
	OnlyWithAttachment    bool               `qs:"onlyWithAttachment"`
	Priorities            []int64            `qs:"priorities"`
	ProjectID             *int64             `qs:"project,omitempty"`
	RequireAllTags        bool               `qs:"tagRequireAll"`
	Search                string             `qs:"search"`
	Sources               []int64            `qs:"sources"`
	StartDate             *time.Time         `qs:"startDate,omitempty"`
	Statuses              []int64            `qs:"statuses"`
	SubjectKeywords       []string           `qs:"subjectKeywords"`
	Tags                  []int64            `qs:"tags"`
	TaskID                int64              `qs:"task"`
	TaskStatuses          []string           `qs:"taskStatuses"`
	Teams                 []int64            `qs:"teams"`
	TicketID              *int64             `qs:"ticket,omitempty"`
	TimeRange             string             `qs:"timeRange"`
	TWCompanyIDs          []int64            `qs:"twCompanyIds"`
	Types                 []int64            `qs:"types"`
	Unassigned            bool               `qs:"unassigned"`
}
