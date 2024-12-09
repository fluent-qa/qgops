package webhooks

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pocketbase/pocketbase"
)

// WebhookRouter handles the routing of incoming webhooks to their appropriate handlers
type WebhookRouter struct {
	app      *pocketbase.PocketBase
	handlers map[string][]WebhookHandler
	mu       sync.RWMutex
}

// WebhookHandler defines the interface for webhook handlers
type WebhookHandler interface {
	// Handle processes the webhook payload
	Handle(payload *WebhookPayload) error
	// ShouldHandle determines if this handler should process the webhook
	ShouldHandle(payload *WebhookPayload) bool
}

// NewWebhookRouter creates a new WebhookRouter instance
func NewWebhookRouter(app *pocketbase.PocketBase) *WebhookRouter {
	return &WebhookRouter{
		app:      app,
		handlers: make(map[string][]WebhookHandler),
	}
}

// RegisterHandler registers a new webhook handler for a specific collection
func (r *WebhookRouter) RegisterHandler(collection string, handler WebhookHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.handlers[collection] == nil {
		r.handlers[collection] = make([]WebhookHandler, 0)
	}
	r.handlers[collection] = append(r.handlers[collection], handler)
}

// Route processes an incoming webhook and routes it to appropriate handlers
func (r *WebhookRouter) Route(payload *WebhookPayload) error {
	r.mu.RLock()
	handlers := r.handlers[payload.Collection]
	r.mu.RUnlock()

	var errors []error
	for _, handler := range handlers {
		if handler.ShouldHandle(payload) {
			if err := handler.Handle(payload); err != nil {
				errors = append(errors, fmt.Errorf("handler error: %w", err))
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("webhook routing errors: %v", errors)
	}
	return nil
}

// Example handler for logging webhooks
type LoggingWebhookHandler struct{}

func (h *LoggingWebhookHandler) Handle(payload *WebhookPayload) error {
	// Log the webhook payload
	jsonData, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Printf("Received webhook: %s\n", string(jsonData))
	return nil
}

func (h *LoggingWebhookHandler) ShouldHandle(payload *WebhookPayload) bool {
	return true // Handle all webhooks
}

// Example handler for specific collection changes
type CollectionChangeHandler struct {
	collection string
	action     string
}

func NewCollectionChangeHandler(collection, action string) *CollectionChangeHandler {
	return &CollectionChangeHandler{
		collection: collection,
		action:     action,
	}
}

func (h *CollectionChangeHandler) Handle(payload *WebhookPayload) error {
	// Handle the specific collection change
	fmt.Printf("Handling %s action for collection %s\n", payload.Action, payload.Collection)
	return nil
}

func (h *CollectionChangeHandler) ShouldHandle(payload *WebhookPayload) bool {
	return payload.Collection == h.collection && payload.Action == h.action
}
