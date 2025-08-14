package models

// Inbox related types
type Inbox struct {
	BaseEntity
	AutoReplyEnabled              bool         `json:"autoReplyEnabled"`
	AutoReplyMessage              string       `json:"autoReplyMessage"`
	AutoReplySubject              string       `json:"autoReplySubject"`
	ClientOnly                    bool         `json:"clientOnly"`
	DisplayOrder                  int          `json:"displayOrder"`
	Email                         string       `json:"email"`
	EmailForwardingState          string       `json:"emailForwardingState"`
	ForwardingAddress             string       `json:"forwardingAddress"`
	HappinessRatingEnabled        bool         `json:"happinessRatingEnabled"`
	HappinessRatingMessage        string       `json:"happinessRatingMessage"`
	IconImage                     string       `json:"iconImage"`
	Inboxaliases                  []EntityRef  `json:"inboxaliases"`
	Inboxcnames                   []InboxCname `json:"inboxcnames"`
	Inboxemailrefs                any          `json:"inboxemailrefs"`
	IncludeTicketHistoryOnForward bool         `json:"includeTicketHistoryOnForward"`
	IsAdmin                       bool         `json:"isAdmin"`
	IsFreeDomain                  bool         `json:"isFreeDomain"`
	LanguageCode                  string       `json:"languageCode"`
	LocalPart                     string       `json:"localPart"`
	Name                          string       `json:"name"`
	NotificationsOnly             bool         `json:"notificationsOnly"`
	Oauth2Token                   any          `json:"oauth2token"`
	OnClosedLock                  string       `json:"onClosedLock"`
	OnClosedWait                  int          `json:"onClosedWait"`
	Projects                      []EntityRef  `json:"projects"`
	PublicIconImage               string       `json:"publicIconImage"`
	Restricteddomains             any          `json:"restricteddomains"`
	SendEmailsFrom                string       `json:"sendEmailsFrom"`
	Signature                     string       `json:"signature"`
	SMTPPassword                  string       `json:"smtpPassword"`
	SMTPPort                      int          `json:"smtpPort"`
	SMTPProvider                  string       `json:"smtpProvider"`
	SMTPSecurity                  string       `json:"smtpSecurity"`
	SMTPServer                    string       `json:"smtpServer"`
	SMTPUsername                  string       `json:"smtpUsername"`
	SpamThreshold                 int          `json:"spamThreshold"`
	Starred                       bool         `json:"starred"`
	SyncAccountID                 any          `json:"syncAccountId"`
	SyncDays                      any          `json:"syncDays"`
	Synced                        bool         `json:"synced"`
	SyncSubscriptionID            any          `json:"syncSubscriptionId"`
	Ticketstatus                  EntityRef    `json:"ticketstatus"`
	Tickettypes                   []EntityRef  `json:"tickettypes"`
	TimeloggingEnabled            bool         `json:"timeloggingEnabled"`
	Triggers                      []Trigger    `json:"triggers"`
	Type                          string       `json:"type"`
	User                          any          `json:"user"`
	Users                         []InboxUser  `json:"users"`
	UseTeamworkMailServer         bool         `json:"useTeamworkMailServer"`
	UsingOfficeHours              bool         `json:"usingOfficeHours"`
	Verified                      bool         `json:"verified"`
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
	Access  string `json:"access"`
	IsAdmin bool   `json:"isAdmin"`
	Starred bool   `json:"starred"`
	State   string `json:"state"`
}

type Trigger struct {
	EntityRef
	Meta struct {
		DisplayOrder int `json:"displayOrder"`
	} `json:"meta"`
}

type InboxCname struct {
	EntityRef
	Meta struct {
		Domain string `json:"domain"`
	} `json:"meta"`
}
