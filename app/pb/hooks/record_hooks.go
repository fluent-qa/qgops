package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterRecordHooks registers all record-related hooks
func RegisterRecordHooks(app *pocketbase.PocketBase) {
	// Before create hook example
	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		// Add custom validation or data modification before creation
		if e.Collection.Name == "posts" {
			// Example: Set default status for new posts
			e.Record.Set("status", "draft")
		}
		return nil
	})

	// After create hook example
	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		// Perform actions after record creation
		if e.Collection.Name == "users" {
			// Example: Send welcome email or create related records
		}
		return nil
	})

	// Before update hook
	app.OnRecordBeforeUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
		// Add validation or modification before update
		return nil
	})

	// After delete hook
	app.OnRecordAfterDeleteRequest().Add(func(e *core.RecordDeleteEvent) error {
		// Clean up related data after deletion
		return nil
	})
}
