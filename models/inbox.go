package models

// Inbox related types
type Inbox struct {
	BaseEntity
	AutoReplyEnabled              *bool        `json:"autoReplyEnabled,omitempty"`
	AutoReplyMessage              *string      `json:"autoReplyMessage,omitempty"`
	AutoReplySubject              *string      `json:"autoReplySubject,omitempty"`
	ClientOnly                    *bool        `json:"clientOnly,omitempty"`
	DisplayOrder                  *int         `json:"displayOrder,omitempty"`
	Email                         *string      `json:"email,omitempty"`
	EmailForwardingState          *string      `json:"emailForwardingState,omitempty"`
	ForwardingAddress             *string      `json:"forwardingAddress,omitempty"`
	HappinessRatingEnabled        *bool        `json:"happinessRatingEnabled,omitempty"`
	HappinessRatingMessage        *string      `json:"happinessRatingMessage,omitempty"`
	IconImage                     *string      `json:"iconImage,omitempty"`
	Inboxaliases                  []EntityRef  `json:"inboxaliases"`
	Inboxcnames                   []InboxCname `json:"inboxcnames"`
	Inboxemailrefs                any          `json:"inboxemailrefs"`
	IncludeTicketHistoryOnForward *bool        `json:"includeTicketHistoryOnForward,omitempty"`
	IsAdmin                       *bool        `json:"isAdmin,omitempty"`
	IsFreeDomain                  *bool        `json:"isFreeDomain,omitempty"`
	LanguageCode                  *string      `json:"languageCode,omitempty"`
	LocalPart                     *string      `json:"localPart,omitempty"`
	Name                          *string      `json:"name,omitempty"`
	NotificationsOnly             *bool        `json:"notificationsOnly,omitempty"`
	Oauth2Token                   any          `json:"oauth2token"`
	OnClosedLock                  *string      `json:"onClosedLock,omitempty"`
	OnClosedWait                  *int         `json:"onClosedWait,omitempty"`
	Projects                      []EntityRef  `json:"projects"`
	PublicIconImage               *string      `json:"publicIconImage,omitempty"`
	Restricteddomains             any          `json:"restricteddomains"`
	SendEmailsFrom                *string      `json:"sendEmailsFrom,omitempty"`
	Signature                     *string      `json:"signature,omitempty"`
	SMTPPassword                  *string      `json:"smtpPassword,omitempty"`
	SMTPPort                      *int         `json:"smtpPort,omitempty"`
	SMTPProvider                  *string      `json:"smtpProvider,omitempty"`
	SMTPSecurity                  *string      `json:"smtpSecurity,omitempty"`
	SMTPServer                    *string      `json:"smtpServer,omitempty"`
	SMTPUsername                  *string      `json:"smtpUsername,omitempty"`
	SpamThreshold                 *int         `json:"spamThreshold,omitempty"`
	Starred                       *bool        `json:"starred,omitempty"`
	SyncAccountID                 any          `json:"syncAccountId"`
	SyncDays                      any          `json:"syncDays"`
	Synced                        *bool        `json:"synced,omitempty"`
	SyncSubscriptionID            any          `json:"syncSubscriptionId"`
	Ticketstatus                  *EntityRef   `json:"ticketstatus,omitempty"`
	Tickettypes                   []EntityRef  `json:"tickettypes"`
	TimeloggingEnabled            *bool        `json:"timeloggingEnabled,omitempty"`
	Triggers                      []Trigger    `json:"triggers"`
	Type                          *string      `json:"type,omitempty"`
	User                          any          `json:"user"`
	Users                         []InboxUser  `json:"users"`
	UseTeamworkMailServer         *bool        `json:"useTeamworkMailServer,omitempty"`
	UsingOfficeHours              *bool        `json:"usingOfficeHours,omitempty"`
	Verified                      *bool        `json:"verified,omitempty"`
}

type InboxesResponse struct {
	Inboxes    []Inbox      `json:"inboxes"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type InboxResponse struct {
	Inbox    Inbox        `json:"inbox"`
	Included IncludedData `json:"included"`
}

type InboxUser struct {
	EntityRef
	Meta InboxMeta `json:"meta"`
}

type InboxMeta struct {
	Access  *string `json:"access,omitempty"`
	IsAdmin *bool   `json:"isAdmin,omitempty"`
	Starred *bool   `json:"starred,omitempty"`
	State   *string `json:"state,omitempty"`
}

type Trigger struct {
	EntityRef
	Meta struct {
		DisplayOrder *int `json:"displayOrder,omitempty"`
	} `json:"meta"`
}

type InboxCname struct {
	EntityRef
	Meta struct {
		Domain *string `json:"domain,omitempty"`
	} `json:"meta"`
}
