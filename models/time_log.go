package models

import "time"

// TimeLog related types
type TimeLog struct {
	BaseEntity
	Billable            *bool      `json:"billable,omitempty"`
	Description         *string    `json:"description,omitempty"`
	Date                *time.Time `json:"date,omitempty"`
	Seconds             *int       `json:"seconds,omitempty"`
	TimezoneOffset      *int       `json:"timezoneOffset,omitempty"`
	TimelogsID          any        `json:"timelogs_id"`
	AssignToCurrentUser *bool      `json:"assignToCurrentUser,omitempty"`
	Ticket              EntityRef  `json:"ticket"`
	User                EntityRef  `json:"user"`
}

type TimeLogsResponse struct {
	TimeLogs   []TimeLog    `json:"timeLogs"`
	Meta       Meta         `json:"meta"`
	Pagination Pagination   `json:"pagination"`
	Included   IncludedData `json:"included"`
}

type TimeLogResponse struct {
	TimeLog  TimeLog      `json:"timeLog"`
	Included IncludedData `json:"included"`
}
