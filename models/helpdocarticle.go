package models

type HelpDocArticle struct {
	BaseEntity
	Helpdocsite     EntityRef `json:"helpdocsite"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	Description     string    `json:"description"`
	OldURL          string    `json:"oldURL"`
	Popularity      int       `json:"popularity"`
	DisqusEnabled   bool      `json:"disqusEnabled"`
	IsPrivate       bool      `json:"isPrivate"`
	EditMethod      string    `json:"editMethod"`
	DisplayOrder    int       `json:"displayOrder"`
	Status          string    `json:"status"`
	Contents        string    `json:"contents"`
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
