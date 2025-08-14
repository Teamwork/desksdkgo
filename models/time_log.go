package models

import "time"

// TimeLog related types
type TimeLog struct {
	BaseEntity
	Billable            bool      `json:"billable"`
	Description         string    `json:"description"`
	Date                time.Time `json:"date"`
	Seconds             int       `json:"seconds"`
	TimezoneOffset      int       `json:"timezoneOffset"`
	TimelogsID          any       `json:"timelogs_id"`
	AssignToCurrentUser bool      `json:"assignToCurrentUser"`
	Ticket              EntityRef `json:"ticket"`
	User                EntityRef `json:"user"`
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
