package routers

import (
	_default "beego_xiaomi/controllers/default"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &_default.IndexController{})
    beego.Router("/login", &_default.LoginController{})

}
