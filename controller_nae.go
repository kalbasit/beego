// +build !appengine

package beego

import "github.com/eMxyzptlk/beego/context"

type Controller struct {
	baseController
}

func (c *Controller) initForAE(ctx *context.Context) {}
