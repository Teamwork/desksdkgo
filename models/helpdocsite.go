package models

type HelpDocSite struct {
	BaseEntity
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	MetaSiteDescription  string           `json:"metaSiteDescription"`
	Subdomain            string           `json:"subdomain"`
	ContactFormEnabled   bool             `json:"contactFormEnabled"`
	ShowDateLastModified bool             `json:"showDateLastModified"`
	CustomDomain         string           `json:"customDomain"`
	CustomStyleSheet     string           `json:"customStyleSheet"`
	HomePageLinkEnabled  bool             `json:"homePageLinkEnabled"`
	HomePageLinkText     string           `json:"homePageLinkText"`
	HomePageURL          string           `json:"homePageURL"`
	HTMLHeadCode         string           `json:"htmlHeadCode"`
	LogoImage            string           `json:"logoImage"`
	Favicon              string           `json:"favicon"`
	TouchIcon            string           `json:"touchIcon"`
	PublicSiteEnabled    bool             `json:"publicSiteEnabled"`
	SendEmailsToInboxID  int              `json:"sendEmailsToInboxId"`
	ShowOnHomePage       string           `json:"showOnHomePage"`
	HeaderBGColor        string           `json:"headerBGColor"`
	NavActiveColor       string           `json:"navActiveColor"`
	NavTextColor         string           `json:"navTextColor"`
	PageBGColor          string           `json:"pageBGColor"`
	LinkColor            string           `json:"linkColor"`
	TextColor            string           `json:"textColor"`
	LanguageCode         string           `json:"languageCode"`
	Password             string           `json:"password"`
	ShowSocialIcons      bool             `json:"showSocialIcons"`
	DisqusShortname      any              `json:"disqusShortname"`
	AuthenticationType   string           `json:"authenticationType"`
	AuthenticationTypeID int              `json:"authenticationTypeId"`
	EditMethod           string           `json:"editMethod"`
	Stats                HelpDocSiteStats `json:"stats"`
	SearchTemplate       string           `json:"searchTemplate"`
	HomeTemplate         string           `json:"homeTemplate"`
	HeadTemplate         string           `json:"headTemplate"`
	FooterTemplate       string           `json:"footerTemplate"`
	CategoryTemplate     string           `json:"categoryTemplate"`
	ArticleTemplate      string           `json:"articleTemplate"`
	Contributors         []EntityRef      `json:"contributors"`
}

type HelpDocSiteStats struct {
	ArticleCount     int `json:"articleCount"`
	PublishedCount   int `json:"publishedCount"`
	UnpublishedCount int `json:"unpublishedCount"`
	DraftCount       int `json:"draftCount"`
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
