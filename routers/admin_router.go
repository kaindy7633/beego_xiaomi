package routers

import (
	"beego_xiaomi/controllers/admin"
	"beego_xiaomi/middleware"
	"beego_xiaomi/models"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 通过命名空间 namespace 来配置路由
	ns := beego.NewNamespace("/"+models.AdminPath,
		// 路由拦截
		beego.NSBefore(middleware.AdminAuth),

		beego.NSRouter("/", &admin.MainController{}),
		beego.NSRouter("/welcome", &admin.MainController{}, "get:Welcome"),

		// 管理员管理
		beego.NSRouter("/manager", &admin.ManagerController{}),
		beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
		beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),
		beego.NSRouter("/manager/doAdd", &admin.ManagerController{}, "post:DoAdd"),
		beego.NSRouter("/manager/doEdit", &admin.ManagerController{}, "post:DoEdit"),
		beego.NSRouter("/manager/delete", &admin.ManagerController{}, "get:Delete"),

		// 用户登录与退出
		beego.NSRouter("/login", &admin.LoginController{}),
		beego.NSRouter("/login/doLogin", &admin.LoginController{}, "post:DoLogin"),
		beego.NSRouter("/login/loginOut", &admin.LoginController{}, "get:LoginOut"),

		beego.NSRouter("/focus", &admin.FocusController{}),

		// 角色管理
		beego.NSRouter("/role", &admin.RoleController{}),
		beego.NSRouter("/role/add", &admin.RoleController{}, "get:Add"),
		beego.NSRouter("/role/edit", &admin.RoleController{}, "get:Edit"),
		beego.NSRouter("/role/doAdd", &admin.RoleController{}, "post:DoAdd"),
		beego.NSRouter("/role/doEdit", &admin.RoleController{}, "post:DoEdit"),
		beego.NSRouter("/role/delete", &admin.RoleController{}, "get:Delete"),
		beego.NSRouter("/role/auth", &admin.RoleController{}, "get:Auth"),
		beego.NSRouter("/role/doAuth", &admin.RoleController{}, "post:DoAuth"),

		// 权限管理
		beego.NSRouter("/access", &admin.AccessController{}),
		beego.NSRouter("/access/add", &admin.AccessController{}, "get:Add"),
		beego.NSRouter("/access/edit", &admin.AccessController{}, "get:Edit"),
		beego.NSRouter("/access/doAdd", &admin.AccessController{}, "post:DoAdd"),
		beego.NSRouter("/access/doEdit", &admin.AccessController{}, "post:DoEdit"),
		beego.NSRouter("/access/delete", &admin.AccessController{}, "get:Delete"),
	)
	beego.AddNamespace(ns)
}
