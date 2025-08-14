package models

// Customer related types
type Customer struct {
	BaseEntity
	FirstName             string      `json:"firstName"`
	LastName              string      `json:"lastName"`
	Email                 string      `json:"email"`
	Organization          string      `json:"organization"`
	ExtraData             string      `json:"extraData"`
	Notes                 string      `json:"notes"`
	VerifiedEmail         bool        `json:"verifiedEmail"`
	LinkedinURL           string      `json:"linkedinURL"`
	FacebookURL           string      `json:"facebookURL"`
	TwitterHandle         string      `json:"twitterHandle"`
	NumTickets            int         `json:"numTickets"`
	JobTitle              any         `json:"jobTitle"`
	Phone                 string      `json:"phone"`
	Mobile                string      `json:"mobile"`
	Address               string      `json:"address"`
	ExternalID            string      `json:"externalId"`
	AvatarURL             string      `json:"avatarURL"`
	Contacts              []EntityRef `json:"contacts"`
	Customerwelcomeemails any         `json:"customerwelcomeemails"`
	Trusted               bool        `json:"trusted"`
	WelcomeEmailSent      bool        `json:"welcomeEmailSent"`
}

// Response types for customers
type CustomersResponse struct {
	Customers  []Customer   `json:"customers"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type CustomerResponse struct {
	Customer Customer     `json:"customer"`
	Included IncludedData `json:"included"`
}
