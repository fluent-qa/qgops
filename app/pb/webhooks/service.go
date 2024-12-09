package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

const webhooksCollection = "webhooks"

type WebhookService struct {
	app    *pocketbase.PocketBase
	router *WebhookRouter
}

func NewWebhookService(app *pocketbase.PocketBase) *WebhookService {
	return &WebhookService{
		app:    app,
		router: NewWebhookRouter(app),
	}
}

func AttachWebhooks(app *pocketbase.PocketBase) {
	service := NewWebhookService(app)
	
	// Register default handlers
	service.router.RegisterHandler("*", &LoggingWebhookHandler{})

	migrations.Register(func(db dbx.Builder) error {
		return daos.New(db).SaveCollection(&models.Collection{
			Name:   webhooksCollection,
			Type:   models.CollectionTypeBase,
			System: true,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:     "name",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "collection",
					Type:     schema.FieldTypeText,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "destination",
					Type:     schema.FieldTypeUrl,
					Required: true,
				},
				&schema.SchemaField{
					Name:     "secret",
					Type:     schema.FieldTypeText,
					Required: false,
				},
				&schema.SchemaField{
					Name:     "enabled",
					Type:     schema.FieldTypeBool,
					Required: true,
					Default:  true,
				},
			),
		})
	}, func(db dbx.Builder) error {
		dao := daos.New(db)
		collection, err := dao.FindCollectionByNameOrId(webhooksCollection)
		if err != nil {
			return err
		}
		return dao.DeleteCollection(collection)
	})

	app.OnModelAfterCreate().Add(func(e *core.ModelEvent) error {
		return service.event(app, "create", e.Model.Collection().Name, e.Model, e.HttpContext)
	})

	app.OnModelAfterUpdate().Add(func(e *core.ModelEvent) error {
		return service.event(app, "update", e.Model.Collection().Name, e.Model, e.HttpContext)
	})

	app.OnModelAfterDelete().Add(func(e *core.ModelEvent) error {
		return service.event(app, "delete", e.Model.Collection().Name, e.Model, e.HttpContext)
	})
}

func (s *WebhookService) event(app *pocketbase.PocketBase, action, collection string, record *models.Record, ctx echo.Context) error {
	dao := app.Dao()
	webhooks := []Webhook{}

	err := dao.DB().Select("*").From(webhooksCollection).
		Where(dbx.HashExp{"collection": collection}).
		Where(dbx.HashExp{"enabled": true}).
		All(&webhooks)

	if err != nil {
		return fmt.Errorf("failed to fetch webhooks: %w", err)
	}

	if len(webhooks) == 0 {
		return nil
	}

	var admin *models.Admin
	var auth *models.Record

	if ctx != nil {
		admin = apis.RequestAdmin(ctx)
		auth = apis.RequestAuthRecord(ctx)
	}

	payload := &WebhookPayload{
		Action:     action,
		Collection: collection,
		Record:     record,
		Admin:      admin,
		Auth:       auth,
	}

	// Route the webhook through registered handlers
	if err := s.router.Route(payload); err != nil {
		fmt.Printf("Error routing webhook: %v\n", err)
	}

	// Send to registered webhook endpoints
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	var wg sync.WaitGroup
	for _, webhook := range webhooks {
		wg.Add(1)
		go func(w Webhook) {
			defer wg.Done()
			retryable := NewRetryableWebhook(w, DefaultRetryConfig)
			if err := retryable.Send(context.Background(), jsonData); err != nil {
				fmt.Printf("Failed to send webhook to %s: %v\n", w.Destination, err)
			}
		}(webhook)
	}

	wg.Wait()
	return nil
}

func sendWebhook(ctx context.Context, webhook Webhook, payload []byte) error {
	req, err := http.NewRequestWithContext(ctx, "POST", webhook.Destination, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "PocketBase-Webhook")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("webhook returned error %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
