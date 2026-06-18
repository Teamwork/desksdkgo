package models

type SearchHelpdocsFilter struct {
	Search     string   `qs:"search,omitempty"`
	SiteID     int64    `qs:"siteId,omitempty"`
	Status     string   `qs:"status,omitempty"`
	CategoryID int64    `qs:"categoryId,omitempty"`
	Page       int      `qs:"page,omitempty"`
	PageSize   int      `qs:"pageSize,omitempty"`
	Includes   []string `qs:"includes,omitempty"`
}

type HelpDocArticle struct {
	BaseEntity
	Helpdocsite     EntityRef `json:"helpdocsite"`
	Title           *string   `json:"title,omitempty"`
	Slug            *string   `json:"slug,omitempty"`
	Description     *string   `json:"description,omitempty"`
	OldURL          *string   `json:"oldURL,omitempty"`
	Popularity      *int      `json:"popularity,omitempty"`
	DisqusEnabled   *bool     `json:"disqusEnabled,omitempty"`
	IsPrivate       *bool     `json:"isPrivate,omitempty"`
	EditMethod      *string   `json:"editMethod,omitempty"`
	DisplayOrder    *int      `json:"displayOrder,omitempty"`
	Status          *string   `json:"status,omitempty"`
	Contents        *string   `json:"contents,omitempty"`
	Categories      []int     `json:"categories"`
	RelatedArticles []int     `json:"relatedArticles,omitempty"`
}

type HelpDocArticlesResponse struct {
	HelpDocArticles []HelpDocArticle `json:"helpdocarticles"`
	Included        IncludedData     `json:"included"`
	Pagination      Pagination       `json:"pagination"`
	Meta            Meta             `json:"meta"`
}

type HelpDocArticleResponse struct {
	HelpDocArticle HelpDocArticle `json:"helpDocArticle"`
	Included       IncludedData   `json:"included"`
}
