package webhooks

import "github.com/pocketbase/pocketbase/models"

type Webhook struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Collection  string `db:"collection" json:"collection"`
	Destination string `db:"destination" json:"destination"`
}

type WebhookPayload struct {
	Action     string         `json:"action"`
	Collection string         `json:"collection"`
	Record     *models.Record `json:"record"`
	Auth       *models.Record `json:"auth,omitempty"`
	Admin      *models.Admin  `json:"admin,omitempty"`
}
