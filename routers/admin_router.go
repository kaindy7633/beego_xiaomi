package routers

import (
	"beego_xiaomi/controllers/admin"
	"beego_xiaomi/models"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 通过命名空间 namespace 来配置路由
	ns := beego.NewNamespace("/"+models.AdminPath,
		// 路由拦截
		beego.NSBefore(func(ctx *context.Context) {
			pathname := ctx.Request.URL.String()
			userinfo, ok := ctx.Input.Session("userinfo").(models.Manager)
			if !(ok && userinfo.Username != "") {
				if pathname != "/"+models.AdminPath+"/login" && pathname != "/"+models.AdminPath+"/login/doLogin" {
					ctx.Redirect(302, "/"+models.AdminPath+"/login")
				}
			}
		}),

		beego.NSRouter("/", &admin.MainController{}),
		beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),

		// 管理员管理
		beego.NSRouter("/manager", &admin.ManagerController{}),
		beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
		beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),
		beego.NSRouter("/manager/doAdd", &admin.ManagerController{}, "post:DoAdd"),
		beego.NSRouter("/manager/doEdit", &admin.ManagerController{}, "post:DoEdit"),
		beego.NSRouter("/manager/delete", &admin.ManagerController{}, "post:Delete"),

		beego.NSRouter("/login", &admin.LoginController{}),
		beego.NSRouter("/login/doLogin", &admin.LoginController{}, "post:DoLogin"),

		beego.NSRouter("/focus", &admin.FocusController{}),

		// 角色管理
		beego.NSRouter("/role", &admin.RoleController{}),
		beego.NSRouter("/role/add", &admin.RoleController{}, "get:Add"),
		beego.NSRouter("/role/edit", &admin.RoleController{}, "get:Edit"),
		beego.NSRouter("/role/doAdd", &admin.RoleController{}, "post:DoAdd"),
		beego.NSRouter("/role/doEdit", &admin.RoleController{}, "post:DoEdit"),
		beego.NSRouter("/role/delete", &admin.RoleController{}, "get:Delete"),
	)
	beego.AddNamespace(ns)
}
