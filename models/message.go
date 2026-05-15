package models

import (
	"encoding/json"
	"time"
)

// Message related types
type Message struct {
	BaseEntity
	AssigningUser      *EntityRef `json:"assigningUser,omitempty"`
	BCC                []string   `json:"bcc"`
	CC                 []string   `json:"cc"`
	Contact            *EntityRef `json:"contact,omitempty"`
	Delayed            *bool      `json:"delayed,omitempty"`
	EditMethod         *string    `json:"editMethod,omitempty"`
	Message            *string    `json:"message,omitempty"`
	IsPinned           *bool      `json:"isPinned,omitempty"`
	Status             *EntityRef `json:"status,omitempty"`
	ThreadType         *string    `json:"threadType,omitempty"`
	Ticket             EntityRef  `json:"ticket"`
	ViewedByCustomerAt *time.Time `json:"viewedByCustomerAt"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	aux := &struct {
		HtmlBody *string `json:"htmlBody"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	m.Message = aux.HtmlBody
	return nil
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
