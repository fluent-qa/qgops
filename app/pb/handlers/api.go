package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// RegisterAPIRoutes registers custom API routes
func RegisterAPIRoutes(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Custom API endpoint for health check
		e.Router.GET("/api/health", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{
				"status": "healthy",
			})
		})

		// Example of protected endpoint
		e.Router.GET("/api/protected", func(c echo.Context) error {
			// Require authenticated user
			user, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
			if user == nil {
				return apis.NewForbiddenError("Only authenticated users can access this endpoint", nil)
			}

			return c.JSON(http.StatusOK, map[string]string{
				"message": "You have access to protected data",
			})
		}, apis.RequireAdminOrRecordAuth())

		return nil
	})
}
