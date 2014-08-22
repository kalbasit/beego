// +build appengine

package beego

import (
	"net/http"
	"path/filepath"

	"github.com/astaxie/beego/context"

	"appengine"
)

type Controller struct {
	baseController
	AppEngineCtx appengine.Context
}

func (c *Controller) initForAE(ctx *context.Context) {
	c.AppEngineCtx = appengine.NewContext(ctx.Request)
}

// Run beego application.
func (app *App) Run() {
	http.Handle("/", app.Handlers)
}

func registerAppPathAndAppConfigPath() {
	AppPath = ""
	AppConfigPath = filepath.Join(AppPath, "conf", "app.conf")
}
