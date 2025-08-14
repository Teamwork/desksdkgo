package models

// BusinessHour represents a tag in the system
type BusinessHour struct {
	BaseEntity

	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDefault         bool   `json:"isDefault"`
	TimezoneID        int64  `json:"timezoneId"`
	TimezoneReference string `json:"timezone_name"`
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
