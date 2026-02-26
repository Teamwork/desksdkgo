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
	DeliveryMethod     string     `json:"deliveryMethod"`
	DeliveryReason     string     `json:"deliveryReason"`
	DeliveryStatus     string     `json:"deliveryStatus"`
	EditMethod         string     `json:"editMethod"`
	EmailMessageID     string     `json:"emailMessageId"`
	Message            string     `json:"message"`
	IsPinned           bool       `json:"isPinned"`
	S3Link             *string    `json:"s3link"`
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
