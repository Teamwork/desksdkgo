package models

// Phone represents a phone number associated with a company
type Phone struct {
	BaseEntity
	Number      *string `json:"number,omitempty"`
	Type        *string `json:"type,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Extension   *string `json:"extension,omitempty"`
}

// Domain represents a domain associated with a company
type Domain struct {
	BaseEntity
	Name            *string `json:"name,omitempty"`
	Project         any     `json:"project"`
	Company         any     `json:"company"`
	ProjectsCompany any     `json:"projects_company"`
}

// Company represents a company in Desk.com
type Company struct {
	BaseEntity
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Details     *string     `json:"details,omitempty"`
	Industry    *string     `json:"industry,omitempty"`
	Website     *string     `json:"website,omitempty"`
	Permission  *string     `json:"permission,omitempty"`
	Kind        *string     `json:"kind,omitempty"`
	Domains     []EntityRef `json:"domains,omitempty"`
	Note        *string     `json:"note,omitempty"`
}

// CompaniesResponse represents the response for a list of companies
type CompaniesResponse struct {
	Companies  []Company    `json:"companies"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

// CompanyResponse represents the response for a single company
type CompanyResponse struct {
	Company  Company      `json:"company"`
	Included IncludedData `json:"included"`
}
