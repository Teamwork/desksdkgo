package models

// Contact related types
type Contact struct {
	BaseEntity
	Value  string `json:"value"`
	IsMain bool   `json:"isMain"`
}
