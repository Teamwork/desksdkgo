package models

// BusinessHour represents a tag in the system
type BusinessHour struct {
	BaseEntity

	Name              *string `json:"name,omitempty"`
	Description       *string `json:"description,omitempty"`
	IsDefault         *bool   `json:"isDefault,omitempty"`
	TimezoneID        *int64  `json:"timezoneId,omitempty"`
	TimezoneReference *string `json:"timezone_name,omitempty"`
}

// BusinessHoursResponse represents the response for a list of businesshours
type BusinessHoursResponse struct {
	BusinessHours []BusinessHour `json:"businesshours"`
	Included      IncludedData   `json:"included"`
	Pagination    Pagination     `json:"pagination"`
	Meta          Meta           `json:"meta"`
}

type BusinessHourResponse struct {
	BusinessHour BusinessHour `json:"businesshour"`
	Included     IncludedData `json:"included"`
}
