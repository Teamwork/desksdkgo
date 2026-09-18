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
	// Helpdocsite is the site the article belongs to. It is reported on a read;
	// on a write the route carries the site, so an unset one is left out.
	Helpdocsite   EntityRef `json:"helpdocsite,omitzero"`
	Title         *string   `json:"title,omitempty"`
	Slug          *string   `json:"slug,omitempty"`
	Description   *string   `json:"description,omitempty"`
	OldURL        *string   `json:"oldURL,omitempty"`
	Popularity    *int      `json:"popularity,omitempty"`
	DisqusEnabled *bool     `json:"disqusEnabled,omitempty"`
	IsPrivate     *bool     `json:"isPrivate,omitempty"`
	EditMethod    *string   `json:"editMethod,omitempty"`
	DisplayOrder  *int      `json:"displayOrder,omitempty"`
	Status        *string   `json:"status,omitempty"`
	Contents      *string   `json:"contents,omitempty"`
	// Categories is the list of categories the article is filed under. The
	// update route binds the body over the stored article and revalidates the
	// result, so an unset list must be left out rather than sent as null: null
	// clears the categories, and the revalidation then rejects the update for
	// having none.
	Categories      []int `json:"categories,omitempty"`
	RelatedArticles []int `json:"relatedArticles,omitempty"`
}

type HelpDocArticlesResponse struct {
	HelpDocArticles []HelpDocArticle `json:"helpdocarticles"`
	Included        IncludedData     `json:"included"`
	Pagination      Pagination       `json:"pagination"`
	Meta            Meta             `json:"meta"`
}

type HelpDocArticleResponse struct {
	HelpDocArticle HelpDocArticle `json:"helpdocarticle"`
	Included       IncludedData   `json:"included"`
}
