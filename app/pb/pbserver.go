package main

//https://github.com/pocketbase/pocketbase.git
//https://github.com/pedrozadotdev/pocketblocks

import (
	"github.com/pocketbase/pocketbase/core"
	"log"
	"os"

	"github.com/fluent-qa/qgops/app/pb/webhooks"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
)

func main() {
	app := pocketbase.New()
	webhooks.AttachWebhooks(app)
	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
