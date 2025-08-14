package api

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"strings"
)

// Service defines the interface that all service types must implement
type Service[T any, R any, L any] interface {
	Get(ctx context.Context, id int) (*R, error)
	List(ctx context.Context, params url.Values) (*L, error)
	Create(ctx context.Context, item *T) (*R, error)
	Update(ctx context.Context, id int, item *T) (*R, error)
}

// Call is a generic function to handle any resource type
func Call[T any, R any, L any](ctx context.Context, service Service[T, R, L], action string, id int, createItem func() *T) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	switch strings.ToLower(action) {
	case "get":
		if id == 0 {
			log.Fatal("ID is required for get action")
		}
		item, err := service.Get(ctx, id)
		if err != nil {
			log.Fatal(err)
		}

		enc.Encode(item)

	case "list":
		items, err := service.List(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		enc.Encode(items)

	case "create":
		item := createItem()
		created, err := service.Create(ctx, item)
		if err != nil {
			log.Print(err)
			return
		}

		enc.Encode(created)

	case "update":
		if id == 0 {
			log.Fatal("ID is required for update action")
		}
		item := createItem()
		updated, err := service.Update(ctx, id, item)
		if err != nil {
			log.Print(err)
			return
		}
		enc.Encode(updated)

	default:
		log.Fatalf("Unsupported action: %s", action)
	}
}
