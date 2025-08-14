package models

// Tag represents a tag in the system
type Tag struct {
	BaseEntity
	Name  string `json:"name"`
	Color string `json:"color"`
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
