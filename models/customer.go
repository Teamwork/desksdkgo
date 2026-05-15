package models

// Customer related types
type Customer struct {
	BaseEntity
	FirstName             *string     `json:"firstName,omitempty"`
	LastName              *string     `json:"lastName,omitempty"`
	Email                 *string     `json:"email,omitempty"`
	Organization          *string     `json:"organization,omitempty"`
	ExtraData             *string     `json:"extraData,omitempty"`
	Notes                 *string     `json:"notes,omitempty"`
	VerifiedEmail         *bool       `json:"verifiedEmail,omitempty"`
	LinkedinURL           *string     `json:"linkedinURL,omitempty"`
	FacebookURL           *string     `json:"facebookURL,omitempty"`
	TwitterHandle         *string     `json:"twitterHandle,omitempty"`
	NumTickets            *int        `json:"numTickets,omitempty"`
	JobTitle              any         `json:"jobTitle"`
	Phone                 *string     `json:"phone,omitempty"`
	Mobile                *string     `json:"mobile,omitempty"`
	Address               *string     `json:"address,omitempty"`
	ExternalID            *string     `json:"externalId,omitempty"`
	AvatarURL             *string     `json:"avatarURL,omitempty"`
	Contacts              []EntityRef `json:"contacts"`
	Customerwelcomeemails any         `json:"customerwelcomeemails"`
	Trusted               *bool       `json:"trusted,omitempty"`
	WelcomeEmailSent      *bool       `json:"welcomeEmailSent,omitempty"`
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
