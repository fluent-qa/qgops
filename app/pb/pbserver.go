package main

import (
	"log"
	"os"

	"github.com/fluent-qa/qgops/app/pb/handlers"
	"github.com/fluent-qa/qgops/app/pb/hooks"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	// Register migrations
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:          "./pb_migrations",
		Automigrate:  true,
		TemplateLang: "",
	})

	// Register custom API routes
	handlers.RegisterAPIRoutes(app)

	// Register record hooks
	hooks.RegisterRecordHooks(app)

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
