package models

import "time"

// Message related types
type Message struct {
	BaseEntity
	HTMLBody           string     `json:"htmlBody"`
	TextBody           string     `json:"textBody"`
	EditMethod         string     `json:"editMethod"`
	EmailMessageID     string     `json:"emailMessageId"`
	S3Link             *string    `json:"s3link"`
	ViewedByCustomerAt *time.Time `json:"viewedByCustomerAt"`
	Delayed            bool       `json:"delayed"`
	Contact            EntityRef  `json:"contact,omitempty"`
	ThreadType         string     `json:"threadType"`
	Ticket             EntityRef  `json:"ticket"`
	DeliveryStatus     string     `json:"deliveryStatus"`
	DeliveryReason     string     `json:"deliveryReason"`
	DeliveryMethod     string     `json:"deliveryMethod"`
	IsPinned           bool       `json:"isPinned"`
	AssigningUser      EntityRef  `json:"assigningUser,omitempty"`
	Status             EntityRef  `json:"status,omitempty"`
}
