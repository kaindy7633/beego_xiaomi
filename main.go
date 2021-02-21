package main

import (
	"beego_xiaomi/models"
	_ "beego_xiaomi/routers"
	"encoding/gob"

	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func init() {
	// 发生错误 gob: name not registered for interface: "beego_xiaomi/models.Manager"
	// 因为 session 内部采用了 gob 来注册存储的对象，例如 struct，所以如果你采用了非 memory 的引擎，
	// 请自己在 main.go 的 init 里面注册需要保存的这些结构体，不然会引起应用重启之后出现无法解析的错误
	// 参考官网：https://beego.me/docs/mvc/controller/session.md
	// 注册这个结构体
	gob.Register(models.Manager{})
}

func main() {
	// 注册模板函数
	beego.AddFuncMap("unixToDate", models.UnixToDate)

	// 配置session保存在redis中
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"

	beego.Run()
}
