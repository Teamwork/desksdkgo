package models

// Contact related types
type Contact struct {
	BaseEntity
	Value  *string `json:"value,omitempty"`
	IsMain *bool   `json:"isMain,omitempty"`
}
