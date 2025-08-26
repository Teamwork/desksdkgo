package models

// Spamlist represents a spamlist entry.  Term can be an email address, domain,
// or IP address.  Type is whitelist or blacklist.
type Spamlist struct {
	BaseEntity
	Term string `json:"term"`
	Type string `json:"type"`
}

// SpamlistsResponse represents the response for a list of spam lists
type SpamlistsResponse struct {
	Spamlists  []Spamlist   `json:"spamlists"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type SpamlistResponse struct {
	Spamlist Spamlist     `json:"spamlist"`
	Included IncludedData `json:"included"`
}
