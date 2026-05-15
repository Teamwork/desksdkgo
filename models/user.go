package models

// User related types
type User struct {
	BaseEntity
	Email                    *string    `json:"email,omitempty"`
	FirstName                *string    `json:"firstName,omitempty"`
	LastName                 *string    `json:"lastName,omitempty"`
	AvatarURL                *string    `json:"avatarURL,omitempty"`
	EditMethod               *string    `json:"editMethod,omitempty"`
	IsPartTime               *bool      `json:"isPartTime,omitempty"`
	TicketReplyRedirect      *string    `json:"ticketReplyRedirect,omitempty"`
	Reviewer                 *bool      `json:"reviewer,omitempty"`
	TrainingWheelsEnrollment *EntityRef `json:"trainingWheelsEnrollment,omitempty"`
	Role                     *string    `json:"role,omitempty"`
	SendPushNotifications    *bool      `json:"sendPushNotifications,omitempty"`
	SendWebNotifications     *bool      `json:"sendWebNotifications,omitempty"`
	AutoFollowOnCC           *bool      `json:"autoFollowOnCC,omitempty"`
	TimeFormatID             *int       `json:"timeFormatId,omitempty"`
	TimezoneID               *int       `json:"timezoneId,omitempty"`
	ProjectsCompanyID        *int       `json:"projectsCompanyId,omitempty"`
	IsAppOwner               *bool      `json:"isAppOwner,omitempty"`
	LdKey                    *string    `json:"ldKey,omitempty"`
}

type UsersResponse struct {
	Users      []User       `json:"users"`
	Included   IncludedData `json:"included"`
	Meta       Meta         `json:"meta"`
	Pagination Pagination   `json:"pagination"`
}

type UserResponse struct {
	User     User         `json:"user"`
	Included IncludedData `json:"included"`
}
