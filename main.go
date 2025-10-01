package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/teamwork/desksdkgo/api"
	"github.com/teamwork/desksdkgo/client"
	"github.com/teamwork/desksdkgo/models"
	"github.com/teamwork/desksdkgo/util"
)

func main() {
	// Load environment variables from .env file
	util.LoadEnv()

	// Define flags with environment variable fallbacks
	apiKey := flag.String("api-key", util.GetEnv("DESK_API_KEY", ""), "Desk API key (can also be set via DESK_API_KEY env var)")
	baseURL := flag.String("base-url", util.GetEnv("DESK_BASE_URL", "https://mycompany.teamwork.com/desk/api/v2"), "Desk API base URL (can also be set via DESK_BASE_URL env var)")
	resource := flag.String("resource", util.GetEnv("DESK_RESOURCE", "tickets"), "Resource to interact with (tickets, customers, companies, users) (can also be set via DESK_RESOURCE env var)")
	action := flag.String("action", util.GetEnv("DESK_ACTION", "list"), "Action to perform (get, list, create, update) (can also be set via DESK_ACTION env var)")
	envCount, _ := strconv.ParseInt(util.GetEnv("DESK_COUNT", "1"), 10, 64)
	count := flag.Int("count", int(envCount), "Number of resources to create (default: 1)")
	id := flag.Int("id", 0, "Resource ID for get/update actions")
	debug := flag.Bool("debug", false, "Enable debug logging")
	data := flag.String("data", "", "JSON data to merge with default values for create/update actions")
	flag.Parse()

	if action == nil || *action == "" {
		log.Fatal("Action is required. Set it via --action flag or DESK_ACTION environment variable")
	}

	if *action != "create" {
		*count = 1 // For get/list/update actions, count should be 1
	}

	// Validate required flags
	if *apiKey == "" {
		log.Fatal("API key is required. Set it via --api-key flag or DESK_API_KEY environment variable")
	}

	// Create client
	opts := []client.Option{}
	if *debug {
		opts = append(opts, client.WithLogLevel(slog.LevelDebug))
	}
	opts = append(opts, client.WithAPIKey(*apiKey))

	c := client.NewClient(*baseURL, opts...)

	// Create context
	ctx := context.Background()

	// Parse JSON data if provided
	var jsonData map[string]interface{}
	if *data != "" {
		if err := json.Unmarshal([]byte(*data), &jsonData); err != nil {
			log.Fatalf("Failed to parse JSON data: %v", err)
		}
	}

	resources := []string{*resource}
	if *resource == "all" {
		resources = []string{
			"businesshours",
			"companies",
			"customers",
			"inboxes",
			"priorities",
			"slas",
			"spamlists",
			"statuses",
			"tags",
			"tickets",
			"types",
		}
	}

	for _, resource := range resources {
		generateData(ctx, c, resource, *action, *count, *id, jsonData)
	}
}

func generateData(
	ctx context.Context,
	c *client.Client,
	resource string,
	action string,
	count int,
	id int,
	jsonData map[string]any,
) {
	// Execute action based on resource and action
	for range count {
		switch strings.ToLower(resource) {
		case "tickets":
			api.Call(ctx, c.Tickets, action, id, func() *models.TicketResponse {
				inboxes, err := c.Inboxes.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list inboxes: %v", err)
				}

				if len(inboxes.Inboxes) == 0 {
					log.Fatal("No inboxes found. Please create an inbox first.")
				}

				customers, err := c.Customers.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list customers: %v", err)
				}

				if len(customers.Customers) == 0 {
					log.Fatal("No customers found. Please create a customer first.")
				}

				types, err := c.TicketTypes.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list ticket types: %v", err)
				}

				var t models.TicketType
				for _, tt := range types.TicketTypes {
					for _, ibx := range inboxes.Inboxes {
						for _, ttibx := range tt.Inboxes {
							if ttibx.ID == ibx.ID {
								t = tt
								break
							}
						}
					}
				}

				if t.ID == 0 {
					log.Fatal("No ticket types associated with the available inboxes.")
				}

				if len(types.TicketTypes) == 0 {
					log.Fatal("No ticket types found. Please create a ticket type first.")
				}

				sources, err := c.TicketSources.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list ticket sources: %v", err)
				}

				if len(sources.TicketSources) == 0 {
					log.Fatal("No ticket sources found. Please create a ticket source first.")
				}

				statuses, err := c.TicketStatuses.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list ticket statuses: %v", err)
				}

				if len(statuses.TicketStatuses) == 0 {
					log.Fatal("No ticket statuses found. Please create a ticket status first.")
				}

				agents, err := c.Users.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list users: %v", err)
				}

				if len(agents.Users) == 0 {
					log.Fatal("No users found. Please create a user first.")
				}

				resp := &models.TicketResponse{Ticket: models.Ticket{
					Subject:           gofakeit.Sentence(1),
					PreviewText:       gofakeit.Paragraph(1, 2, 3, " "),
					OriginalRecipient: gofakeit.Email(),
					Inbox: models.EntityRef{
						ID: inboxes.Inboxes[0].ID,
					},
					Customer: models.EntityRef{
						ID: customers.Customers[0].ID,
					},
					Body: gofakeit.Paragraph(3, 5, 10, "\n"),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.Ticket, jsonData)
				}
				return resp
			})
		case "customers":
			api.Call(ctx, c.Customers, action, id, func() *models.CustomerResponse {
				email := gofakeit.Email()
				resp := &models.CustomerResponse{
					Customer: models.Customer{
						FirstName: gofakeit.FirstName(),
						LastName:  gofakeit.LastName(),
						Email:     email,
					},
					Included: models.IncludedData{
						Contacts: []models.Contact{
							{
								BaseEntity: models.BaseEntity{
									Type: "email",
								},
								Value:  email,
								IsMain: true,
							},
						},
					},
				}
				if jsonData != nil {
					util.MergeJSONData(&resp.Customer, jsonData)
				}
				return resp
			})
		case "companies":
			api.Call(ctx, c.Companies, action, id, func() *models.CompanyResponse {
				resp := &models.CompanyResponse{
					Company: models.Company{
						Name:        gofakeit.Company(),
						Description: gofakeit.Paragraph(1, 2, 3, " "),
					},
					Included: models.IncludedData{
						Domains: []models.Domain{
							{
								Name: gofakeit.DomainName(),
							},
						},
					},
				}
				if jsonData != nil {
					util.MergeJSONData(&resp.Company, jsonData)
				}
				return resp
			})
		case "users":
			api.Call(ctx, c.Users, action, id, func() *models.UserResponse {
				resp := &models.UserResponse{User: models.User{
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
					Email:     gofakeit.Email(),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.User, jsonData)
				}
				return resp
			})
		case "tags":
			api.Call(ctx, c.Tags, action, id, func() *models.TagResponse {
				resp := &models.TagResponse{Tag: models.Tag{
					Name: gofakeit.Word(),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.Tag, jsonData)
				}
				return resp
			})
		case "files":
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			f := &models.FileResponse{File: models.File{
				Filename:    gofakeit.LoremIpsumWord() + "." + gofakeit.FileExtension(),
				MIMEType:    "image/jpeg",
				Type:        models.FileTypeAttachment,
				Disposition: models.DispositionAttachment,
			}}
			if jsonData != nil {
				util.MergeJSONData(&f.File, jsonData)
			}

			resp, err := c.Files.Create(ctx, f)
			if err != nil {
				log.Fatalf("Failed to create file reference: %v", err)
			}

			err = c.Files.Upload(ctx, resp, []byte(gofakeit.ImageJpeg(800, 600)))
			if err != nil {
				log.Fatalf("Failed to upload file: %v", err)
			}

			enc.Encode(resp)
		case "spamlists":
			api.Call(ctx, c.Spamlists, action, id, func() *models.SpamlistResponse {
				resp := &models.SpamlistResponse{Spamlist: models.Spamlist{
					Term: gofakeit.Email(),
					Type: "blacklist",
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.Spamlist, jsonData)
				}
				return resp
			})
		case "statuses":
			api.Call(ctx, c.TicketStatuses, action, id, func() *models.TicketStatusResponse {
				resp := &models.TicketStatusResponse{TicketStatus: models.TicketStatus{
					Name: gofakeit.Word(),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.TicketStatus, jsonData)
				}
				return resp
			})
		case "types":
			api.Call(ctx, c.TicketTypes, action, id, func() *models.TicketTypeResponse {
				resp := &models.TicketTypeResponse{TicketType: models.TicketType{
					Name: gofakeit.Word(),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.TicketType, jsonData)
				}
				return resp
			})
		case "priorities":
			api.Call(ctx, c.TicketPriorities, action, id, func() *models.TicketPriorityResponse {
				resp := &models.TicketPriorityResponse{TicketPriority: models.TicketPriority{
					Name:  gofakeit.Word(),
					Color: gofakeit.SafeColor(),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.TicketPriority, jsonData)
				}
				return resp
			})
		case "helpdocsites":
			api.Call(ctx, c.HelpDocSites, action, id, func() *models.HelpDocSiteResponse {
				resp := &models.HelpDocSiteResponse{HelpDocSite: models.HelpDocSite{
					Name: gofakeit.Company() + " Help Center",
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.HelpDocSite, jsonData)
				}
				return resp
			})
		case "helpdocarticles":
			api.Call(ctx, c.HelpDocArticles, action, id, func() *models.HelpDocArticleResponse {
				resp := &models.HelpDocArticleResponse{HelpDocArticle: models.HelpDocArticle{
					Title:    gofakeit.Sentence(5),
					Contents: gofakeit.Paragraph(3, 5, 10, "\n"),
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.HelpDocArticle, jsonData)
				}
				return resp
			})
		case "businesshours":
			api.Call(ctx, c.BusinessHours, action, id, func() *models.BusinessHourResponse {
				resp := &models.BusinessHourResponse{BusinessHour: models.BusinessHour{
					Name:      gofakeit.Company() + " Business Hours",
					IsDefault: true,
				}}
				if jsonData != nil {
					util.MergeJSONData(&resp.BusinessHour, jsonData)
				}
				return resp
			})

		case "inboxes":
			api.Call(ctx, c.Inboxes, action, id, func() *models.InboxResponse {
				users, err := c.Users.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list users: %v", err)
				}
				if len(users.Users) == 0 {
					log.Fatal("No users found. Please create a user first.")
				}

				resp := &models.InboxResponse{Inbox: models.Inbox{
					Name:      gofakeit.Company() + " Inbox",
					Email:     gofakeit.Email(),
					LocalPart: strings.SplitN(gofakeit.Email(), "@", 2)[0],
				}}

				for _, user := range users.Users {
					resp.Inbox.Users = append(resp.Inbox.Users, models.InboxUser{
						EntityRef: models.EntityRef{
							ID: user.ID,
						},
						Meta: models.InboxMeta{
							Access: "write",
						},
					})
				}

				if jsonData != nil {
					util.MergeJSONData(&resp.Inbox, jsonData)
				}
				return resp
			})
		case "slas":
			api.Call(ctx, c.SLAs, action, id, func() *models.SLAResponse {
				priorities, err := c.TicketPriorities.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list ticketpriorities: %v", err)
				}

				if len(priorities.TicketPriorities) == 0 {
					log.Fatal("No ticketpriorities found. Please create a ticketpriority first.")
				}

				tags, err := c.Tags.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list tags: %v", err)
				}

				if len(tags.Tags) == 0 {
					log.Fatal("No tags found. Please create a tag first.")
				}

				companies, err := c.Companies.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list companies: %v", err)
				}

				if len(companies.Companies) == 0 {
					log.Fatal("No companies found. Please create a company first.")
				}

				customers, err := c.Customers.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list customers: %v", err)
				}

				if len(customers.Customers) == 0 {
					log.Fatal("No customers found. Please create a customer first.")
				}

				inboxes, err := c.Inboxes.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list inboxes: %v", err)
				}

				if len(inboxes.Inboxes) == 0 {
					log.Fatal("No inboxes found. Please create an inbox first.")
				}

				businesshours, err := c.BusinessHours.List(ctx, nil)
				if err != nil {
					log.Fatalf("Failed to list businesshours: %v", err)
				}

				if len(businesshours.BusinessHours) == 0 {
					log.Fatal("No businesshours found. Please create a businesshour first.")
				}

				resp := &models.SLAResponse{
					SLA: models.SLA{
						Name: gofakeit.Company() + " SLA Policy",
						BusinessHour: &models.EntityRef{
							ID: businesshours.BusinessHours[0].ID,
						},
					},
					Included: models.IncludedData{
						SLANotifications: []models.SLANotification{
							{
								Condition:          models.SLANotificationConditionTypeWarning,
								Type:               models.SLANotificationTypeFirstResponse,
								Duration:           gofakeit.Number(1, 10),
								NotifyAssignedUser: true,
							},
							{
								Condition:          models.SLANotificationConditionTypeBreach,
								Type:               models.SLANotificationTypeFirstResponse,
								Duration:           0,
								NotifyAssignedUser: true,
							},
						},
					},
				}

				for _, priority := range priorities.TicketPriorities {
					resp.Included.SLAPriorities = append(resp.Included.SLAPriorities, models.SLATicketPriority{
						Hours:       gofakeit.Number(1, 10),
						Minutes:     gofakeit.Number(1, 59),
						Description: "SLA for " + priority.Name,
						TicketPriority: &models.EntityRef{
							ID: priority.ID,
						},
					})
				}

				resp.Included.SLAPriorities = append(resp.Included.SLAPriorities, models.SLATicketPriority{
					Hours:       gofakeit.Number(1, 10),
					Minutes:     gofakeit.Number(1, 59),
					Description: "SLA for None",
				})

				for _, inbox := range inboxes.Inboxes {
					resp.Included.SLAInboxes = append(resp.Included.SLAInboxes, models.SLAInbox{
						Inbox: &models.EntityRef{
							ID: inbox.ID,
						},
						Condition: models.SLAConditionOptionEqual,
					})

					if len(resp.Included.SLAInboxes) > 4 {
						break
					}
				}

				for _, company := range companies.Companies {
					resp.Included.SLACompanies = append(resp.Included.SLACompanies, models.SLACompany{
						Company: &models.EntityRef{
							ID: company.ID,
						},
						Condition: models.SLAConditionOptionEqual,
					})

					if len(resp.Included.SLACompanies) > 4 {
						break
					}
				}

				for _, customer := range customers.Customers {
					resp.Included.SLACustomers = append(resp.Included.SLACustomers, models.SLACustomer{
						Customer: &models.EntityRef{
							ID: customer.ID,
						},
						Condition: models.SLAConditionOptionEqual,
					})

					if len(resp.Included.SLACustomers) > 3 {
						break
					}
				}

				for _, tag := range tags.Tags {
					resp.Included.SLATags = append(resp.Included.SLATags, models.SLATag{
						Tag: &models.EntityRef{
							ID: tag.ID,
						},
						Condition: models.SLAConditionOptionEqual,
					})

					if len(resp.Included.SLATags) > 6 {
						break
					}
				}

				if jsonData != nil {
					util.MergeJSONData(&resp.SLA, jsonData)
				}
				return resp
			})
		default:
			log.Fatalf("Unsupported resource: %s", resource)
		}
	}
}
