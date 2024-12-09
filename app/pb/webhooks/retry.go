package webhooks

import (
	"context"
	"fmt"
	"time"
)

// RetryConfig defines the configuration for webhook retries
type RetryConfig struct {
	MaxRetries  int
	InitialWait time.Duration
	MaxWait     time.Duration
}

// DefaultRetryConfig provides default retry configuration
var DefaultRetryConfig = RetryConfig{
	MaxRetries:  3,
	InitialWait: 5 * time.Second,
	MaxWait:     1 * time.Minute,
}

// RetryableWebhook wraps a webhook with retry capability
type RetryableWebhook struct {
	webhook Webhook
	config  RetryConfig
}

// NewRetryableWebhook creates a new RetryableWebhook instance
func NewRetryableWebhook(webhook Webhook, config RetryConfig) *RetryableWebhook {
	return &RetryableWebhook{
		webhook: webhook,
		config:  config,
	}
}

// Send attempts to send the webhook with retries
func (r *RetryableWebhook) Send(ctx context.Context, payload []byte) error {
	var lastErr error
	wait := r.config.InitialWait

	for attempt := 0; attempt <= r.config.MaxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return fmt.Errorf("context cancelled during retry: %w", ctx.Err())
			case <-time.After(wait):
				// Exponential backoff
				wait *= 2
				if wait > r.config.MaxWait {
					wait = r.config.MaxWait
				}
			}
		}

		err := sendWebhook(ctx, r.webhook, payload)
		if err == nil {
			return nil
		}

		lastErr = err
		// Log retry attempt
		fmt.Printf("Webhook delivery failed (attempt %d/%d): %v\n", 
			attempt+1, r.config.MaxRetries+1, err)
	}

	return fmt.Errorf("webhook delivery failed after %d attempts: %w", 
		r.config.MaxRetries+1, lastErr)
}
