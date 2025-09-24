package client

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
)

// Client represents the Desk API client
type Client struct {
	baseURL    string
	apiKey     string
	logLevel   slog.Level
	logger     *slog.Logger
	httpClient *http.Client
	middleware []MiddlewareFunc

	// Services
	BusinessHours    *BusinessHourService
	Companies        *CompanyService
	Customers        *CustomerService
	HelpDocArticles  *HelpDocArticleService
	HelpDocSites     *HelpDocSiteService
	Inboxes          *InboxService
	SLAs             *SLAService
	Spamlists        *SpamlistService
	Tags             *TagService
	TicketPriorities *TicketPriorityService
	Tickets          *TicketService
	TicketStatuses   *TicketStatusService
	TicketTypes      *TicketTypeService
	Users            *UserService
}

type MiddlewareFunc func(ctx context.Context, next HandlerFunc) HandlerFunc
type HandlerFunc func(c Client) error

// Config represents the client configuration
type Config struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// Option is a function that configures a Client
type Option func(*Client)

// WithAPIKey sets the API key for the client
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithLogLevel sets the log level for the client
func WithLogLevel(level slog.Level) Option {
	return func(c *Client) {
		c.logLevel = level
	}
}

// WithLogger sets a custom logger for the client
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithMiddleware adds middleware to the client
func WithMiddleware(mw MiddlewareFunc) Option {
	return func(c *Client) {
		c.middleware = append(c.middleware, mw)
	}
}

// NewClient creates a new Desk.com API client
func NewClient(baseURL string, opts ...Option) *Client {
	client := &Client{
		baseURL: baseURL,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		client.httpClient = NewLoggingClientWithLogger(client.logLevel, client.logger)
	}

	// Initialize services
	client.BusinessHours = NewBusinessHourService(client)
	client.Companies = NewCompanyService(client)
	client.Customers = NewCustomerService(client)
	client.HelpDocArticles = NewHelpDocArticleService(client)
	client.HelpDocSites = NewHelpDocSiteService(client)
	client.Inboxes = NewInboxService(client)
	client.SLAs = NewSLAService(client)
	client.Spamlists = NewSpamlistService(client)
	client.Tags = NewTagService(client)
	client.TicketPriorities = NewTicketPriorityService(client)
	client.Tickets = NewTicketService(client)
	client.TicketStatuses = NewTicketStatusService(client)
	client.TicketTypes = NewTicketTypeService(client)
	client.Users = NewUserService(client)

	return client
}

// doRequest performs an HTTP request with the client's configuration
func (c *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	// Add API key if set
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}

	// Add content type
	req.Header.Set("Content-Type", "application/json")

	// Add accept header
	req.Header.Set("Accept", "application/json")

	// Run the middleware chain
	if len(c.middleware) > 0 {
		var finalHandler HandlerFunc = func(c Client) error {
			return nil
		}

		for i := len(c.middleware) - 1; i >= 0; i-- {
			finalHandler = c.middleware[i](ctx, finalHandler)
		}

		if err := finalHandler(*c); err != nil {
			return nil, err
		}
	}

	return c.httpClient.Do(req)
}

// ListOptions represents options for list operations
type ListOptions struct {
	Page    int
	PerPage int
	SortBy  string
	SortDir string
	Embed   string
	Fields  string
	Q       string
}

// Encode encodes the options into a query string
func (o *ListOptions) Encode() string {
	if o == nil {
		return ""
	}

	v := url.Values{}
	if o.Page > 0 {
		v.Set("page", strconv.Itoa(o.Page))
	}
	if o.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(o.PerPage))
	}
	if o.SortBy != "" {
		v.Set("sort_by", o.SortBy)
	}
	if o.SortDir != "" {
		v.Set("sort_dir", o.SortDir)
	}
	if o.Embed != "" {
		v.Set("embed", o.Embed)
	}
	if o.Fields != "" {
		v.Set("fields", o.Fields)
	}
	if o.Q != "" {
		v.Set("q", o.Q)
	}

	return v.Encode()
}
