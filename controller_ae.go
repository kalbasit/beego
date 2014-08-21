// +build appengine

package beego

import (
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
