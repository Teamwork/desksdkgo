package models

import "time"

type State string

const (
	// StateActive represents an active state
	StateActive State = "active"
	// StateDeleted represents a deleted state
	StateDeleted State = "deleted"
)

// Base types for common fields
type BaseEntity struct {
	ID        int        `json:"id"`
	Type      any        `json:"type"` // Can be string or object
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	CreatedBy *UserRef   `json:"createdBy,omitempty"`
	UpdatedBy *UserRef   `json:"updatedBy,omitempty"`
	State     State      `json:"state"`
}

type UserRef struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type EntityRef struct {
	ID   int            `json:"id"`
	Type string         `json:"type"`
	Meta map[string]any `json:"meta"`
}
