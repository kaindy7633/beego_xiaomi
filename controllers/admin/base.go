package admin

import (
	"beego_xiaomi/models"

	beego "github.com/beego/beego/v2/server/web"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

// Success ...
func (c *BaseController) Success(message string, redirect string) {
	c.Data["message"] = message
	c.Data["redirect"] = "/" + models.AdminPath + redirect
	c.TplName = "admin/public/success.html"
}

// Error ...
func (c *BaseController) Error(message string, redirect string) {
	c.Data["message"] = message
	c.Data["redirect"] = "/" + models.AdminPath + redirect
	c.TplName = "admin/public/error.html"
}
