package models

// HelpDocCategory is a category of a help doc site. Every article is filed
// under at least one, and a category belongs to a single site.
type HelpDocCategory struct {
	BaseEntity
	Name                 *string `json:"name,omitempty"`
	Slug                 *string `json:"slug,omitempty"`
	ParentID             *int    `json:"parent_id,omitempty"`
	SiteID               *int    `json:"helpDocsSites_id,omitempty"`
	DisplayOrder         *int    `json:"displayOrder,omitempty"`
	OldURL               *string `json:"oldURL,omitempty"`
	DisplayOnDocHomepage *bool   `json:"displayOnDocHomepage,omitempty"`
}

type HelpDocCategoriesResponse struct {
	HelpDocCategories []HelpDocCategory `json:"helpdocscategories"`
	Included          IncludedData      `json:"included"`
	Pagination        Pagination        `json:"pagination"`
	Meta              Meta              `json:"meta"`
}

type HelpDocCategoryResponse struct {
	HelpDocCategory HelpDocCategory `json:"helpdocscategory"`
	Included        IncludedData    `json:"included"`
}
