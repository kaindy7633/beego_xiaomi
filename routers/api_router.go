package routers

import (
	"beego_xiaomi/controllers/api"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSRouter("/login", &api.LoginController{}),
		)
	beego.AddNamespace(ns)

}