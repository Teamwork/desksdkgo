package models

import "time"

// Message related types
type Message struct {
	BaseEntity
	AssigningUser      EntityRef  `json:"assigningUser,omitempty"`
	BCC                []string   `json:"bcc"`
	CC                 []string   `json:"cc"`
	Contact            EntityRef  `json:"contact,omitempty"`
	Delayed            bool       `json:"delayed"`
	EditMethod         string     `json:"editMethod"`
	Message            string     `json:"message"`
	IsPinned           bool       `json:"isPinned"`
	Status             EntityRef  `json:"status,omitempty"`
	ThreadType         string     `json:"threadType"`
	Ticket             EntityRef  `json:"ticket"`
	ViewedByCustomerAt *time.Time `json:"viewedByCustomerAt"`
}

type MessageResponse struct {
	Message  Message      `json:"message"`
	Included IncludedData `json:"included"`
}

type MessagesResponse struct {
	Messages   []Message    `json:"messages"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}
