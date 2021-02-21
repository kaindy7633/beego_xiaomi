package admin

import (
	beego "github.com/beego/beego/v2/server/web"
)

// MainController ...
type MainController struct {
	beego.Controller
}

// Get ...
func (c *MainController) Get() {
	c.TplName = "admin/main/index.html"
}

// Welcome ...
func (c *MainController) Welcome() {
	c.TplName = "admin/main/welcome.html"
}
