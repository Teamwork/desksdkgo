package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/teamwork/desksdkgo/models"
)

// TestHelpDocArticleServiceRoutesAreSiteScoped pins the route and verb of every
// article operation.
//
// An article is addressed through its site — helpdocssites/{siteID}/
// helpdocarticles — and the update is a PATCH. Addressing them without the site
// reaches no route at all: the endpoint answers 404, which reads as a missing
// article rather than as a wrong path.
func TestHelpDocArticleServiceRoutesAreSiteScoped(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*Client) error
		wantMethod string
		wantPath   string
	}{{
		name: "get",
		call: func(c *Client) error {
			_, err := c.HelpDocArticles.GetForSite(context.Background(), 7, 42, url.Values{})
			return err
		},
		wantMethod: http.MethodGet,
		wantPath:   "/helpdocssites/7/helpdocarticles/42.json",
	}, {
		name: "create",
		call: func(c *Client) error {
			_, err := c.HelpDocArticles.CreateForSite(context.Background(), 7, &models.HelpDocArticleResponse{
				HelpDocArticle: models.HelpDocArticle{Title: ptr("New Article")},
			})
			return err
		},
		wantMethod: http.MethodPost,
		wantPath:   "/helpdocssites/7/helpdocarticles.json",
	}, {
		name: "update",
		call: func(c *Client) error {
			_, err := c.HelpDocArticles.UpdateForSite(context.Background(), 7, 42, &models.HelpDocArticleResponse{
				HelpDocArticle: models.HelpDocArticle{Title: ptr("Updated")},
			})
			return err
		},
		wantMethod: http.MethodPatch,
		wantPath:   "/helpdocssites/7/helpdocarticles/42.json",
	}, {
		name: "list",
		call: func(c *Client) error {
			_, err := c.HelpDocArticles.List(context.Background(), url.Values{})
			return err
		},
		wantMethod: http.MethodGet,
		wantPath:   "/helpdocssites/helpdocarticles.json",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTransport := NewMockRoundTripper()
			mockTransport.AddResponse(tt.wantMethod, tt.wantPath, http.StatusOK, models.HelpDocArticleResponse{
				HelpDocArticle: models.HelpDocArticle{BaseEntity: models.BaseEntity{ID: 42}},
			})

			client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))
			if err := tt.call(client); err != nil {
				t.Fatalf("call returned error: %v", err)
			}

			requests := mockTransport.GetRequests()
			if len(requests) != 1 {
				t.Fatalf("expected 1 request, got %d", len(requests))
			}
			if got := requests[0].Method; got != tt.wantMethod {
				t.Errorf("expected method %s, got %s", tt.wantMethod, got)
			}
			if got := requests[0].URL.Path; got != tt.wantPath {
				t.Errorf("expected path %s, got %s", tt.wantPath, got)
			}
		})
	}
}

// TestHelpDocArticleServiceDerivesTheSiteFromTheArticle pins that the plain
// Create and Update take the site from the article's own helpdocsite reference,
// the same way MessageService.Create takes the ticket from the message.
func TestHelpDocArticleServiceDerivesTheSiteFromTheArticle(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodPost, "/helpdocssites/7/helpdocarticles.json",
		http.StatusCreated, models.HelpDocArticleResponse{
			HelpDocArticle: models.HelpDocArticle{BaseEntity: models.BaseEntity{ID: 42}},
		})

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	if _, err := client.HelpDocArticles.Create(context.Background(), &models.HelpDocArticleResponse{
		HelpDocArticle: models.HelpDocArticle{
			Helpdocsite: models.EntityRef{ID: 7, Type: "helpdocsites"},
			Title:       ptr("New Article"),
		},
	}); err != nil {
		t.Fatalf("Create() returned error: %v", err)
	}

	requests := mockTransport.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
	if got, want := requests[0].URL.Path, "/helpdocssites/7/helpdocarticles.json"; got != want {
		t.Errorf("expected path %s, got %s", want, got)
	}

	if _, err := client.HelpDocArticles.Create(context.Background(), &models.HelpDocArticleResponse{
		HelpDocArticle: models.HelpDocArticle{Title: ptr("Siteless")},
	}); err == nil {
		t.Error("expected an error for an article naming no site")
	}
}

// TestHelpDocArticleBodyUsesTheAPIWrapperKey pins the key a write body is
// wrapped in, and that an unset category list is left out of it.
//
// The routes read and answer "helpdocarticle". Under any other key the endpoint
// binds nothing and rejects the write as having no title, while a response
// decodes into an empty struct — and both look like an empty article rather
// than a mismatched key.
//
// The categories half matters because the update route binds the body over the
// stored article and revalidates the result: a null clears them, and the
// revalidation then rejects the update for having none.
func TestHelpDocArticleBodyUsesTheAPIWrapperKey(t *testing.T) {
	mockTransport := NewMockRoundTripper()
	mockTransport.AddResponse(http.MethodPatch, "/helpdocssites/7/helpdocarticles/42.json",
		http.StatusOK, map[string]any{"helpdocarticle": map[string]any{"id": 42, "title": "Updated"}})

	client := NewClient("https://example.com", WithHTTPClient(&http.Client{Transport: mockTransport}))

	updated, err := client.HelpDocArticles.UpdateForSite(context.Background(), 7, 42,
		&models.HelpDocArticleResponse{
			HelpDocArticle: models.HelpDocArticle{Title: ptr("Updated")},
		})
	if err != nil {
		t.Fatalf("UpdateForSite() returned error: %v", err)
	}
	if updated.HelpDocArticle.Title == nil || *updated.HelpDocArticle.Title != "Updated" {
		t.Errorf("expected the response to decode under %q, got %+v", "helpdocarticle", updated.HelpDocArticle)
	}

	requests := mockTransport.GetRequests()
	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
	body, err := io.ReadAll(requests[0].Body)
	if err != nil {
		t.Fatalf("failed to read the request body: %v", err)
	}

	var sent map[string]any
	if err := json.Unmarshal(body, &sent); err != nil {
		t.Fatalf("failed to decode the request body %q: %v", body, err)
	}
	article, ok := sent["helpdocarticle"].(map[string]any)
	if !ok {
		t.Fatalf("expected the article under %q, got %q", "helpdocarticle", body)
	}
	if got, want := article["title"], "Updated"; got != want {
		t.Errorf("expected title %v, got %v", want, got)
	}
	for _, key := range []string{"categories", "helpdocsite", "contents", "status"} {
		if _, ok := article[key]; ok {
			t.Errorf("expected no %s in the body, got %v", key, article[key])
		}
	}
}
