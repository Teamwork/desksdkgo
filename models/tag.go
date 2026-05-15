package models

// Tag represents a tag in the system
type Tag struct {
	BaseEntity
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

// TagsResponse represents the response for a list of tags
type TagsResponse struct {
	Tags       []Tag        `json:"tags"`
	Included   IncludedData `json:"included"`
	Pagination Pagination   `json:"pagination"`
	Meta       Meta         `json:"meta"`
}

type TagResponse struct {
	Tag      Tag          `json:"tag"`
	Included IncludedData `json:"included"`
}
