package admin

import beego "github.com/beego/beego/v2/server/web"

// FocusController ...
type FocusController struct {
	beego.Controller
}

// Get ...
func (c *FocusController) Get() {
	c.Ctx.WriteString("后台管理系统的轮播图管理")
}
