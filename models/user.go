package models

// User related types
type User struct {
	BaseEntity
	Email                    string    `json:"email"`
	FirstName                string    `json:"firstName"`
	LastName                 string    `json:"lastName"`
	AvatarURL                string    `json:"avatarURL"`
	EditMethod               string    `json:"editMethod"`
	IsPartTime               bool      `json:"isPartTime"`
	TicketReplyRedirect      string    `json:"ticketReplyRedirect"`
	Reviewer                 bool      `json:"reviewer"`
	TrainingWheelsEnrollment EntityRef `json:"trainingWheelsEnrollment"`
	Role                     string    `json:"role"`
	SendPushNotifications    bool      `json:"sendPushNotifications"`
	SendWebNotifications     bool      `json:"sendWebNotifications"`
	AutoFollowOnCC           bool      `json:"autoFollowOnCC"`
	TimeFormatID             int       `json:"timeFormatId"`
	TimezoneID               int       `json:"timezoneId"`
	ProjectsCompanyID        int       `json:"projectsCompanyId"`
	IsAppOwner               bool      `json:"isAppOwner"`
	LdKey                    string    `json:"ldKey"`
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
