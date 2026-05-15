package models

type HelpDocSite struct {
	BaseEntity
	Name                 *string          `json:"name,omitempty"`
	Description          *string          `json:"description,omitempty"`
	MetaSiteDescription  *string          `json:"metaSiteDescription,omitempty"`
	Subdomain            *string          `json:"subdomain,omitempty"`
	ContactFormEnabled   *bool            `json:"contactFormEnabled,omitempty"`
	ShowDateLastModified *bool            `json:"showDateLastModified,omitempty"`
	CustomDomain         *string          `json:"customDomain,omitempty"`
	CustomStyleSheet     *string          `json:"customStyleSheet,omitempty"`
	HomePageLinkEnabled  *bool            `json:"homePageLinkEnabled,omitempty"`
	HomePageLinkText     *string          `json:"homePageLinkText,omitempty"`
	HomePageURL          *string          `json:"homePageURL,omitempty"`
	HTMLHeadCode         *string          `json:"htmlHeadCode,omitempty"`
	LogoImage            *string          `json:"logoImage,omitempty"`
	Favicon              *string          `json:"favicon,omitempty"`
	TouchIcon            *string          `json:"touchIcon,omitempty"`
	PublicSiteEnabled    *bool            `json:"publicSiteEnabled,omitempty"`
	SendEmailsToInboxID  *int             `json:"sendEmailsToInboxId,omitempty"`
	ShowOnHomePage       *string          `json:"showOnHomePage,omitempty"`
	HeaderBGColor        *string          `json:"headerBGColor,omitempty"`
	NavActiveColor       *string          `json:"navActiveColor,omitempty"`
	NavTextColor         *string          `json:"navTextColor,omitempty"`
	PageBGColor          *string          `json:"pageBGColor,omitempty"`
	LinkColor            *string          `json:"linkColor,omitempty"`
	TextColor            *string          `json:"textColor,omitempty"`
	LanguageCode         *string          `json:"languageCode,omitempty"`
	Password             *string          `json:"password,omitempty"`
	ShowSocialIcons      *bool            `json:"showSocialIcons,omitempty"`
	DisqusShortname      any              `json:"disqusShortname"`
	AuthenticationType   *string          `json:"authenticationType,omitempty"`
	AuthenticationTypeID *int             `json:"authenticationTypeId,omitempty"`
	EditMethod           *string          `json:"editMethod,omitempty"`
	Stats                *HelpDocSiteStats `json:"stats,omitempty"`
	SearchTemplate       *string          `json:"searchTemplate,omitempty"`
	HomeTemplate         *string          `json:"homeTemplate,omitempty"`
	HeadTemplate         *string          `json:"headTemplate,omitempty"`
	FooterTemplate       *string          `json:"footerTemplate,omitempty"`
	CategoryTemplate     *string          `json:"categoryTemplate,omitempty"`
	ArticleTemplate      *string          `json:"articleTemplate,omitempty"`
	Contributors         []EntityRef      `json:"contributors"`
}

type HelpDocSiteStats struct {
	ArticleCount     *int `json:"articleCount,omitempty"`
	PublishedCount   *int `json:"publishedCount,omitempty"`
	UnpublishedCount *int `json:"unpublishedCount,omitempty"`
	DraftCount       *int `json:"draftCount,omitempty"`
}

type HelpDocSitesResponse struct {
	HelpDocSites []HelpDocSite `json:"helpdocssites"`
	Included     IncludedData  `json:"included"`
	Pagination   Pagination    `json:"pagination"`
	Meta         Meta          `json:"meta"`
}

type HelpDocSiteResponse struct {
	HelpDocSite HelpDocSite  `json:"helpdocssite"`
	Included    IncludedData `json:"included"`
}
