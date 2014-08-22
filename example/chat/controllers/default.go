// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/eMxyzptlk/beego for the canonical source repository
// @license     http://github.com/eMxyzptlk/beego/blob/master/LICENSE
// @authors     Unknwon

package controllers

import (
	"github.com/eMxyzptlk/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["host"] = this.Ctx.Request.Host
	this.TplNames = "index.tpl"
}
